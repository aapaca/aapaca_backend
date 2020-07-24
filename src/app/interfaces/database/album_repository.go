package database

import (
	"fmt"
	"../../domain"
)

type AlbumRepository struct {
	SqlHandler
}

func (repo *AlbumRepository) FindById(id int) (album domain.Album, err error) {
	row, err := repo.Query("SELECT id, name FROM albums WHERE id = ?", id)
	defer row.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	var album_id int
	var name string
	row.Next()
	if err = row.Scan(&album_id, &name); err != nil {
		return
	}
	album.ID = album_id
	album.Name = name
	return
}