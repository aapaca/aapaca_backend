package repository

import "domain"

type AlbumRepository interface {
	GetAlbum(int) (domain.Album, error)
	GetAlbumsByArtistId(int) ([]domain.Album, error)
}
