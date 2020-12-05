package usecase

import (
	"domain"
	"errors"
)

type AlbumInteractor struct {
	AlbumRepository AlbumRepository
}

func (interactor *AlbumInteractor) GetAlbum(albumId int) (album domain.Album, err error) {
	album, err = interactor.AlbumRepository.GetAlbum(albumId)
	if album.Name == "" {
		err = errors.New("album not found")
	}
	return
}

func (interactor *AlbumInteractor) GetAlbumsByArtistId(artistId int) (albums []domain.Album, err error) {
	albums, err = interactor.AlbumRepository.GetAlbumsByArtistId(artistId)
	if len(albums) == 0 {
		err = errors.New("albums not found")
	}
	return
}
