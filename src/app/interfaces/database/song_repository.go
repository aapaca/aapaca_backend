package database

import (
	"database/sql"
	"domain"
	"errors"
)

type SongRepository struct {
	SqlHandler
}

func values(albumMap map[int]*domain.Album) []domain.Album {
	vs := []domain.Album{}
	for _, v := range albumMap {
		vs = append(vs, *v)
	}
	return vs
}

func (repo *SongRepository) GetSong(id int) (song domain.Song, err error) {
	rows, err := repo.Query(`SELECT songs.id, songs.name, p_artist.id, p_artist.name, p_artist.image_url,
								albums.id, albums.name, albums.image_url, albums.released_date
							FROM songs
							INNER JOIN artists as p_artist
								ON songs.primary_artist_id = p_artist.id
								AND songs.id = ?
							LEFT OUTER JOIN contains
								ON songs.id = contains.song_id
							LEFT OUTER JOIN albums
								ON contains.album_id = albums.id
							`, id)
	defer rows.Close()
	pArtist := domain.Artist{}
	albumMap := map[int]*domain.Album{}
	for rows.Next() {
		album := domain.Album{}
		var releasedDate sql.NullTime
		if err = rows.Scan(&song.ID, &song.Name, &pArtist.ID, &pArtist.Name, &pArtist.ImageURL, &album.ID, &album.Name, &album.ImageURL, &releasedDate); err != nil {
			return
		}
		if _, ok := albumMap[album.ID]; !ok {
			if releasedDate.Valid {
				album.ReleasedDate = &releasedDate.Time
			}
			albumMap[album.ID] = album
		}
	}
	song.Albums = values(albumMap)
	song.PrimaryArtist = pArtist
	return
}

func (repo *SongRepository) GetAttendedSongs(artistId int) (songs []domain.Song, err error) {
	rows, err := repo.Query(`SELECT DISTINCT songs.id, songs.name, albums.id, albums.name, albums.image_url, albums.released_date
							FROM songs
							INNER JOIN performs
								ON performs.song_id = songs.id
								AND performs.artist_id = ?
							INNER JOIN contains
								ON contains.song_id = songs.id
							INNER JOIN albums
								ON albums.id = contains.album_id
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
	if len(songs) == 0 {
		err = errors.New("songs not found")
		return
	}
	return
}

func (repo *SongRepository) GetSongsInAlbum(albumId int) (songs []domain.Song, err error) {
	// TODO: add song length (time)
	rows, err := repo.Query(`SELECT songs.id, songs.name, songs.amazon_music_url, songs.apple_music_url, songs.spotify_url, contains.song_order
							FROM songs
							INNER JOIN contains
								ON contains.album_id = ?
								AND songs.id = contains.song_id
							`, albumId)
	defer rows.Close()
	for rows.Next() {
		song := domain.Song{}
		var amazon, apple, spotify string
		if err = rows.Scan(&song.ID, &song.Name, &amazon, &apple, &spotify, &song.Order); err != nil {
			return
		}
		links := map[string]string{}
		if len(amazon) > 0 {
			links["amazonMusic"] = amazon
		}
		if len(apple) > 0 {
			links["appleMusic"] = apple
		}
		if len(spotify) > 0 {
			links["spotify"] = spotify
		}
		if len(links) > 0 {
			song.Links = links
		}
		songs = append(songs, song)
	}
	if len(songs) == 0 {
		err = errors.New("songs not found")
		return
	}
	return
}
