package interactor

import (
	"domain"
	"errors"
	"usecases/repository"
)

type SongInteractor struct {
	SongRepository repository.SongRepository
}

func (interactor *SongInteractor) GetSong(songId int) (song domain.Song, err error) {
	song, err = interactor.SongRepository.GetSong(songId)
	if err != nil {
		return
	}
	if song.Name == "" {
		err = errors.New("song not found")
	}
	return
}

func (interactor *SongInteractor) GetAttendedSongs(artistId int) (songs []domain.Song, err error) {
	songs, err = interactor.SongRepository.GetAttendedSongs(artistId)
	if err != nil {
		return
	}
	if len(songs) == 0 {
		err = errors.New("songs not found")
	}
	return
}

func (interactor *SongInteractor) GetSongsInAlbum(albumId int) (songs []domain.Song, err error) {
	songs, err = interactor.SongRepository.GetSongsInAlbum(albumId)
	if err != nil {
		return
	}
	if len(songs) == 0 {
		err = errors.New("songs not found")
	}
	return
}
