package usecase

import "domain"

type AlbumUsecase interface {
	GetAlbum(int) (domain.Album, error)
	GetAlbumsByArtistId(int) ([]domain.Album, error)
}
