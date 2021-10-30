package interactor

import (
	"domain"
	"errors"
	"usecases/repository"
)

type ArtistInteractor struct {
	ArtistRepository repository.ArtistRepository
}

func (interactor *ArtistInteractor) GetArtist(artistId int) (domain.Artist, error) {
	artist, err := interactor.ArtistRepository.GetArtist(artistId)
	if err != nil {
		return artist, err
	}
	if artist.Name == "" {
		return artist, errors.New("artist not found")
	}

	members, err := interactor.ArtistRepository.FindMembers(artistId)
	if err != nil {
		return artist, err
	}
	if len(members) > 0 {
		artist.Members = members
	}

	aliases, err := interactor.ArtistRepository.FindAliases(artistId)
	if err != nil {
		return artist, err
	}
	if len(aliases) > 0 {
		artist.Aliases = aliases
	}

	return artist, nil
}
