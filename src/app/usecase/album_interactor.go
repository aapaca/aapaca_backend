package usecase

import "domain"

type AlbumInteractor struct {
	AlbumRepository AlbumRepository
}

func (interactor *AlbumInteractor) GetAlbum(albumId int) (album domain.Album, err error) {
	album, err = interactor.AlbumRepository.GetAlbum(albumId)
	return
}

func (interactor *AlbumInteractor) GetAlbumsByArtistId(artistId int) (albums []domain.Album, err error) {
	albums, err = interactor.AlbumRepository.GetAlbumsByArtistId(artistId)
	return
}
