package usecase

import "domain"

type ArtistUsecase interface {
	GetArtist(int) (domain.Artist, error)
}
