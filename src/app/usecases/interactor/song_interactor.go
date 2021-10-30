package interactor

import (
	"domain"
	"errors"
	"usecases/repository"
)

type SongInteractor struct {
	SongRepository repository.SongRepository
}

func (interactor *SongInteractor) GetSong(songId int) (domain.Song, error) {
	song, err := interactor.SongRepository.GetSong(songId)
	if err == nil && song.Name == "" {
		err = errors.New("song not found")
	}
	return song, err
}

func (interactor *SongInteractor) GetAttendedSongs(artistId int) ([]domain.Song, error) {
	songs, err := interactor.SongRepository.GetAttendedSongs(artistId)
	if err == nil && len(songs) == 0 {
		err = errors.New("songs not found")
	}
	return songs, err
}

func (interactor *SongInteractor) GetSongsInAlbum(albumId int) ([]domain.Song, error) {
	songs, err := interactor.SongRepository.GetSongsInAlbum(albumId)
	if err == nil && len(songs) == 0 {
		err = errors.New("songs not found")
	}
	return songs, err
}
