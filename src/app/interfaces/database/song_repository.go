package database

import (
	"database/sql"
	"domain"
)

type SongRepository struct {
	SqlHandler
}

func generateSongLinks(amazon string, apple string, spotify string) map[string]string {
	links := map[string]string{}
	if len(amazon) > 0 {
		links["amazonMusic"] = "https://www.amazon.com/dp/" + amazon
	}
	if len(apple) > 0 {
		links["appleMusic"] = "https://music.apple.com/album/" + apple
	}
	if len(spotify) > 0 {
		links["spotify"] = "https://open.spotify.com/track/" + spotify
	}
	return links
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
								contents.amazon_music_id, contents.apple_music_id, contents.spotify_id
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
							`, id)
	defer rows.Close()
	var genre, songLen sql.NullString // mysqlのTIME型はgoのtime.Timeで受け取れない
	pArtist := domain.Artist{}
	albumMap := map[int]domain.Album{}
	creditMap := map[int]*domain.Credit{}
	links := map[string]string{}
	for rows.Next() {
		album := domain.Album{}
		var aID, pID sql.NullInt64
		var aName, aImgURL, pTitle sql.NullString
		var amazon, apple, spotify string
		var releasedDate sql.NullTime
		if err = rows.Scan(&song.ID, &song.Name, &genre, &songLen, &pArtist.ID, &pArtist.Name, &pArtist.ImageURL, &album.ID, &album.Name, &album.ImageURL, &releasedDate, &aID, &aName, &aImgURL, &pID, &pTitle, &amazon, &apple, &spotify); err != nil {
			return
		}
		if _, ok := albumMap[album.ID]; !ok {
			if releasedDate.Valid {
				album.ReleasedDate = &releasedDate.Time
				// 最古のalbumのリリース日をsongのリリース日と定める
				if song.ReleasedDate == nil || album.ReleasedDate.Before(*song.ReleasedDate) {
					song.ReleasedDate = album.ReleasedDate
				}
			}
			albumMap[album.ID] = album
			// songs.linksに入れるリンクは、リンク数が多いアルバムを採用
			// 採用されたアルバムに対応するimageURLをsong.ImageURLに代入
			ls := generateSongLinks(amazon, apple, spotify)
			if len(links) < len(ls) {
				links = ls
				song.ImageURL = album.ImageURL
			}
			// song.ImageURLがnilになるのを防ぐ
			if len(song.ImageURL) == 0 && len(album.ImageURL) > 0 {
				song.ImageURL = album.ImageURL
			}
		}
		if !aID.Valid { // no credit information
			continue
		}
		artistID := int(aID.Int64)
		if _, ok := creditMap[artistID]; !ok {
			artist := domain.Artist{ID: artistID, Name: aName.String, ImageURL: aImgURL.String}
			part := domain.Occupation{ID: int(pID.Int64), Title: pTitle.String}
			creditMap[artistID] = &domain.Credit{Artist: artist, Parts: []domain.Occupation{part}}
		} else {
			partID := int(pID.Int64)
			exist := false
			for _, p := range creditMap[artistID].Parts {
				if p.ID == partID {
					exist = true
					break
				}
			}
			if !exist {
				part := domain.Occupation{ID: partID, Title: pTitle.String}
				creditMap[artistID].Parts = append(creditMap[artistID].Parts, part)
			}
		}
	}
	if genre.Valid {
		song.Genre = genre.String
	}
	if songLen.Valid {
		song.SongLen = shortenSongLen(songLen.String)
	}
	if len(links) > 0 {
		song.Links = links
	}
	for _, v := range albumMap {
		song.Albums = append(song.Albums, v)
	}
	for _, v := range creditMap {
		song.Credits = append(song.Credits, *v)
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
		song.Albums = append(song.Albums, album)
		songs = append(songs, song)
	}
	return
}

func (repo *SongRepository) GetSongsInAlbum(albumId int) (songs []domain.Song, err error) {
	// TODO: add song length (time)
	rows, err := repo.Query(`SELECT songs.id, songs.name, songs.song_len,
								contents.amazon_music_id, contents.apple_music_id, contents.spotify_id, contents.song_order
							FROM songs
							INNER JOIN contents
								ON contents.album_id = ?
								AND songs.id = contents.song_id
							`, albumId)
	defer rows.Close()
	for rows.Next() {
		song := domain.Song{}
		var amazon, apple, spotify string
		var songLen sql.NullString
		if err = rows.Scan(&song.ID, &song.Name, &songLen, &amazon, &apple, &spotify, &song.Order); err != nil {
			return
		}
		if songLen.Valid {
			song.SongLen = shortenSongLen(songLen.String)
		}
		links := generateSongLinks(amazon, apple, spotify)
		if len(links) > 0 {
			song.Links = links
		}
		songs = append(songs, song)
	}
	return
}
