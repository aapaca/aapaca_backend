package usecase

import "domain"

type SongInteractor struct {
	SongRepository SongRepository
}

func (interactor *SongInteractor) GetSong(songId int) (song domain.Song, err error) {
	song, err = interactor.SongRepository.GetSong(songId)
	return
}

func (interactor *SongInteractor) GetAttendedSongs(artistId int) (songs []domain.Song, err error) {
	songs, err = interactor.SongRepository.GetAttendedSongs(artistId)
	return
}

func (interactor *SongInteractor) GetSongsInAlbum(albumId int) (songs []domain.Song, err error) {
	songs, err = interactor.SongRepository.GetSongsInAlbum(albumId)
	return
}
