package usecase

import "domain"

type ArtistUsecase interface {
	/*
	 * IDに対応するアーティストを取得する
	 */
	GetArtist(int) (domain.Artist, error)
}
