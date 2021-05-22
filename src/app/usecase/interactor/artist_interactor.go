package interactor

import (
	"domain"
	"errors"
	"usecase/repository"
)

type ArtistInteractor struct {
	ArtistRepository repository.ArtistRepository
}

func (interactor *ArtistInteractor) GetArtist(artistId int) (artist domain.Artist, err error) {
	artist, err = interactor.ArtistRepository.GetArtist(artistId)
	if err != nil {
		return
	}
	if artist.Name == "" {
		err = errors.New("artist not found")
	}
	return
}
