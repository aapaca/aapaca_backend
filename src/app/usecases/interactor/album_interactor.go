package interactor

import (
	"domain"
	"errors"
	"usecases/repository"
)

type AlbumInteractor struct {
	AlbumRepository repository.AlbumRepository
}

func (interactor *AlbumInteractor) GetAlbum(albumId int) (domain.Album, error) {
	album, err := interactor.AlbumRepository.GetAlbum(albumId)
	if err == nil && album.Name == "" {
		err = errors.New("album not found")
	}
	return album, err
}

func (interactor *AlbumInteractor) GetAlbumsByArtistId(artistId int) ([]domain.Album, error) {
	albums, err := interactor.AlbumRepository.GetAlbumsByArtistId(artistId)
	if err == nil && len(albums) == 0 {
		err = errors.New("albums not found")
	}
	return albums, err
}
