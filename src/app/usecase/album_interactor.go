package usecase

import "domain"

type AlbumInteractor struct {
	AlbumRepository AlbumRepository
}

func (interactor *AlbumInteractor) AlbumById(id int) (album domain.Album, err error) {
	album, err = interactor.AlbumRepository.FindById(id)
	return
}
