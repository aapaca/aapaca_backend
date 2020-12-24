package usecase

import (
	"domain"
	"errors"
)

type ArtistInteractor struct {
	ArtistRepository ArtistRepository
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
