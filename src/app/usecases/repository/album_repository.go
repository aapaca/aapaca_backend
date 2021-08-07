package repository

import "domain"

type AlbumRepository interface {
	/*
	 * IDに対応するアルバムを取得する
	 */
	GetAlbum(int) (domain.Album, error)

	/*
	 * アーティストIDをもとに、
	 * primary artistとしてリリースしたアルバムのリストを取得する
	 */
	GetAlbumsByArtistId(int) ([]domain.Album, error)
}
