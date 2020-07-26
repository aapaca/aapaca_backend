package database

import (
	"domain"
)

type AlbumRepository struct {
	SqlHandler
}

func (repo *AlbumRepository) FindById(id int) (album domain.Album, err error) {
	row, err := repo.Query("SELECT albums.id, albums.name, artists.name FROM albums INNER JOIN artists ON albums.primary_artist_id = artists.id WHERE albums.id = ?", id)
	defer row.Close()
	if err != nil {
		return
	}
	var album_id int
	var name string
	var primary_artist_name string
	row.Next()
	if err = row.Scan(&album_id, &name, &primary_artist_name); err != nil {
		return
	}
	album.ID = album_id
	album.Name = name
	album.PrimaryArtist = primary_artist_name
	return
}