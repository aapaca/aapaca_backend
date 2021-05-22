package repository

import "domain"

type SongRepository interface {
	GetSong(int) (domain.Song, error)
	GetAttendedSongs(int) ([]domain.Song, error)
	GetSongsInAlbum(int) ([]domain.Song, error)
}
