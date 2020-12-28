package database

import (
	"database/sql"
	"domain"
	"errors"
)

type SongRepository struct {
	SqlHandler
}

func generateSongLink(id string, serviceName string) (string, string, error) {
	if serviceName == "amazon_music" {
		return "amazonMusic", "https://www.amazon.com/dp/" + id, nil
	}
	if serviceName == "apple_music" {
		return "appleMusic", "https://music.apple.com/album/" + id, nil
	}
	if serviceName == "spotify" {
		return "spotify", "https://open.spotify.com/track/" + id, nil
	}
	return "", "", errors.New("invalid service name")
}

func shortenSongLen(songLen string) string {
	var i int
	for i = 0; i < 4; i++ {
		if string(songLen[i]) != ":" && string(songLen[i]) != "0" {
			break
		}
	}
	return songLen[i:]
}

func (repo *SongRepository) GetSong(id int) (song domain.Song, err error) {
	rows, err := repo.Query(`SELECT songs.id, songs.name, songs.genre, songs.song_len,
								p_artist.id, p_artist.name, p_artist.image_url,
								albums.id, albums.name, albums.image_url, albums.released_date,
								artists.id, artists.name, artists.image_url, oc.id, oc.title,
								external_ids.external_id, external_services.name
							FROM songs
							INNER JOIN artists as p_artist
								ON songs.primary_artist_id = p_artist.id
								AND songs.id = ?
							INNER JOIN contents
								ON songs.id = contents.song_id
							INNER JOIN albums
								ON contents.album_id = albums.id
							LEFT OUTER JOIN performances
								ON performances.song_id = songs.id
							LEFT OUTER JOIN occupations as oc
								ON oc.id = performances.occupation_id
							LEFT OUTER JOIN artists
								ON artists.id = performances.artist_id
							LEFT OUTER JOIN external_ids
								ON external_ids.record_id = songs.id
								AND external_ids.record_type = 3
							LEFT OUTER JOIN external_services
								ON external_ids.service_id = external_services.id
							`, id)
	defer rows.Close()
	var genre, songLen sql.NullString // mysqlのTIME型はgoのtime.Timeで受け取れない
	var releasedDate sql.NullTime
	pArtist := domain.Artist{}
	album := domain.Album{}
	creditMap := map[int]*domain.Credit{}
	links := map[string]string{}
	for rows.Next() {
		var aID, pID sql.NullInt64
		var aName, aImgURL, pTitle, extID, extSName sql.NullString
		if err = rows.Scan(&song.ID, &song.Name, &genre, &songLen, &pArtist.ID, &pArtist.Name, &pArtist.ImageURL, &album.ID, &album.Name, &album.ImageURL, &releasedDate, &aID, &aName, &aImgURL, &pID, &pTitle, &extID, &extSName); err != nil {
			return
		}
		if extID.Valid {
			c, l, e := generateSongLink(extID.String, extSName.String)
			if err = e; err != nil {
				return
			}
			links[c] = l
		}
		if !aID.Valid { // no credit information
			continue
		}
		artistID := int(aID.Int64)
		if _, ok := creditMap[artistID]; !ok {
			artist := domain.Artist{ID: artistID, Name: aName.String, ImageURL: aImgURL.String}
			creditMap[artistID] = &domain.Credit{Artist: artist, Parts: []domain.Occupation{}}
		}
		partID := int(pID.Int64)
		if partExists(creditMap[artistID].Parts, partID) {
			continue
		}
		part := domain.Occupation{ID: partID, Title: pTitle.String}
		creditMap[artistID].Parts = append(creditMap[artistID].Parts, part)
	}
	if genre.Valid {
		song.Genre = genre.String
	}
	if songLen.Valid {
		song.SongLen = shortenSongLen(songLen.String)
	}
	for _, v := range creditMap {
		song.Credits = append(song.Credits, *v)
	}
	if releasedDate.Valid {
		album.ReleasedDate = &releasedDate.Time
	}
	song.Album = album
	if len(links) > 0 {
		song.Links = links
	}
	song.PrimaryArtist = pArtist
	return
}

func (repo *SongRepository) GetAttendedSongs(artistId int) (songs []domain.Song, err error) {
	rows, err := repo.Query(`SELECT DISTINCT songs.id, songs.name, albums.id, albums.name, albums.image_url, albums.released_date
							FROM songs
							INNER JOIN performances
								ON performances.song_id = songs.id
								AND performances.artist_id = ?
							INNER JOIN contents
								ON contents.song_id = songs.id
							INNER JOIN albums
								ON albums.id = contents.album_id
							WHERE songs.primary_artist_id <> ?
							`, artistId, artistId)
	defer rows.Close()
	for rows.Next() {
		song := domain.Song{}
		album := domain.Album{}
		var releasedDate sql.NullTime
		if err = rows.Scan(&song.ID, &song.Name, &album.ID, &album.Name, &album.ImageURL, &releasedDate); err != nil {
			return
		}
		if releasedDate.Valid {
			album.ReleasedDate = &releasedDate.Time
		}
		song.Album = album
		songs = append(songs, song)
	}
	return
}

func (repo *SongRepository) GetSongsInAlbum(albumId int) (songs []domain.Song, err error) {
	rows, err := repo.Query(`SELECT songs.id, songs.name, songs.song_len,
								external_ids.external_id, external_services.name, contents.song_order
							FROM songs
							INNER JOIN contents
								ON contents.album_id = ?
								AND songs.id = contents.song_id
							LEFT OUTER JOIN external_ids
								ON external_ids.record_id = songs.id
								AND external_ids.record_type = 3
							LEFT OUTER JOIN external_services
								ON external_ids.service_id = external_services.id
							`, albumId)
	defer rows.Close()
	songMap := map[int]*domain.Song{}
	for rows.Next() {
		song := domain.Song{}
		var songLen, extID, extSName sql.NullString
		if err = rows.Scan(&song.ID, &song.Name, &songLen, &extID, &extSName, &song.Order); err != nil {
			return
		}
		if _, ok := songMap[song.ID]; !ok {
			if songLen.Valid {
				song.SongLen = shortenSongLen(songLen.String)
			}
			songMap[song.ID] = &song
		}
		if !extID.Valid {
			continue
		}
		if songMap[song.ID].Links == nil {
			links := map[string]string{}
			songMap[song.ID].Links = links
		}
		c, l, e := generateSongLink(extID.String, extSName.String)
		if err = e; err != nil {
			return
		}
		songMap[song.ID].Links[c] = l
	}
	for _, v := range songMap {
		songs = append(songs, *v)
	}
	return
}
