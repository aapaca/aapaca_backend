package repository

import "domain"

type SongRepository interface {
	/*
	 * IDに対応する曲を取得する
	 */
	GetSong(int) (domain.Song, error)

	/*
	 * アーティストIDをもとに、
	 * アーティストが参加している曲のリストを取得する
	 * ただし参加曲のうちアーティストがPrimary Artistとして参加している曲は除く
	 */
	GetAttendedSongs(int) ([]domain.Song, error)

	/*
	 * アルバムIDをもとに、
	 * アルバムに収録されている曲のリストを取得する
	 */
	GetSongsInAlbum(int) ([]domain.Song, error)
}
