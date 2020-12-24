package usecase

import "domain"

type ArtistRepository interface {
	GetArtist(int) (domain.Artist, error)
}