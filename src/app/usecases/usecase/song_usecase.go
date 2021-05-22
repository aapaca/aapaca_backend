package usecase

import "domain"

type SongUsecase interface {
	GetSong(int) (domain.Song, error)
	GetAttendedSongs(int) ([]domain.Song, error)
	GetSongsInAlbum(int) ([]domain.Song, error)
}
