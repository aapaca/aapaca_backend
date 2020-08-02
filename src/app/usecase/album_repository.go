package usecase

import "domain"

type AlbumRepository interface {
	FindById(int) (domain.Album, error)
}
