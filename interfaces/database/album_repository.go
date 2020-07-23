package database

import (
	"time"
	"../../domain"
)

type AlbumRepository struct {
	SqlHandler
}

func (repo *AlbumRepository) FindById(id int) (album domain.Album, err error) {
	row, err := repo.Query("SELECT id, name, primary_artist_id, label, released_date FROM albums WHERE id = ?", id)
	defer row.Close()
	if err != nil {
		return
	}
	var album_id int
	var name string
	var label string
	var released_date time.Time
	row.Next()
	if err = row.Scan(&album_id, &name, &label, &released_date); err != nil {
		return
	}
	album.ID = album_id
	album.Name = name
	album.Label = label
	album.ReleasedDate = released_date
	return
}