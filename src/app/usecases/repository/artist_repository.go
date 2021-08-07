package repository

import "domain"

type ArtistRepository interface {
	/*
	 * IDに対応するアーティストを取得する
	 */
	GetArtist(int) (domain.Artist, error)

	/*
	 * グループメンバーと各メンバーのグループにおける担当パートを取得する
	 */
	FindMembers(int) ([]domain.Credit, error)

	/*
	 * アーティストのエイリアス(別名)を取得する
	 */
	FindAliases(int) ([]domain.Credit, error)
}
