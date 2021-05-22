package interactor

import (
	"domain"
	"errors"
	"usecase/repository"
)

type AlbumInteractor struct {
	AlbumRepository repository.AlbumRepository
}

func (interactor *AlbumInteractor) GetAlbum(albumId int) (album domain.Album, err error) {
	album, err = interactor.AlbumRepository.GetAlbum(albumId)
	if err != nil {
		return
	}
	if album.Name == "" {
		err = errors.New("album not found")
	}
	return
}

func (interactor *AlbumInteractor) GetAlbumsByArtistId(artistId int) (albums []domain.Album, err error) {
	albums, err = interactor.AlbumRepository.GetAlbumsByArtistId(artistId)
	if err != nil {
		return
	}
	if len(albums) == 0 {
		err = errors.New("albums not found")
	}
	return
}
