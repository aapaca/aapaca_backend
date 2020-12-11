package usecase

import (
	"domain"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type SongRepositoryMock struct {
	mock.Mock
}

func (_m *SongRepositoryMock) GetSong(songId int) (domain.Song, error) {
	ret := _m.Called(songId)
	return ret.Get(0).(domain.Song), ret.Error(1)
}

func (_m *SongRepositoryMock) GetAttendedSongs(artistId int) ([]domain.Song, error) {
	ret := _m.Called(artistId)
	return ret.Get(0).([]domain.Song), ret.Error(1)
}

func (_m *SongRepositoryMock) GetSongsInAlbum(albumId int) ([]domain.Song, error) {
	ret := _m.Called(albumId)
	return ret.Get(0).([]domain.Song), ret.Error(1)
}

func TestGetSong(t *testing.T) {
	testSong := domain.Song{ID: 1, Name: "Song1"}
	emptySong := domain.Song{}
	dbError := errors.New("DB error")
	SongRepositoryMock := new(SongRepositoryMock)
	SongRepositoryMock.On("GetSong", testSong.ID).Return(testSong, nil).Once()
	SongRepositoryMock.On("GetSong", 0).Return(emptySong, nil).Once()
	SongRepositoryMock.On("GetSong", -1).Return(emptySong, dbError).Once()
	songInteractor := &SongInteractor{
		SongRepository: SongRepositoryMock,
	}
	t.Run("Normal Case", func(t *testing.T) {
		got, err := songInteractor.GetSong(1)
		if err != nil {
			t.Error(err)
		}
		assert.Equal(t, testSong, got, "Error")
	})
	t.Run("Error Case (Song does not exist)", func(t *testing.T) {
		_, err := songInteractor.GetSong(0)
		expectedError := errors.New("song not found")
		assert.Equal(t, expectedError, err, "Error")
	})
	t.Run("Error Case (Repository returns error)", func(t *testing.T) {
		_, err := songInteractor.GetSong(-1)
		assert.Equal(t, dbError, err, "Error")
	})
}

func TestGetAttendedSongs(t *testing.T) {
	testArtist := domain.Artist{ID: 100}
	credit := domain.Credit{Artist: testArtist}
	testSong1 := domain.Song{ID: 1, Name: "Song1", Credits: []domain.Credit{credit}}
	testSong2 := domain.Song{ID: 2, Name: "Song2", Credits: []domain.Credit{credit}}
	testSongs := []domain.Song{testSong1, testSong2}
	emptySongs := []domain.Song{}
	dbError := errors.New("DB error")
	SongRepositoryMock := new(SongRepositoryMock)
	SongRepositoryMock.On("GetAttendedSongs", testArtist.ID).Return(testSongs, nil).Once()
	SongRepositoryMock.On("GetAttendedSongs", 0).Return(emptySongs, nil).Once()
	SongRepositoryMock.On("GetAttendedSongs", -1).Return(emptySongs, dbError).Once()
	songInteractor := &SongInteractor{
		SongRepository: SongRepositoryMock,
	}
	t.Run("Normal Case", func(t *testing.T) {
		got, err := songInteractor.GetAttendedSongs(testArtist.ID)
		if err != nil {
			t.Error(err)
		}
		assert.ElementsMatch(t, testSongs, got, "Error")
	})
	t.Run("Error Case (Song does not exist)", func(t *testing.T) {
		_, err := songInteractor.GetAttendedSongs(0)
		expectedError := errors.New("songs not found")
		assert.Equal(t, expectedError, err, "Error")
	})
	t.Run("Error Case (Repository returns error)", func(t *testing.T) {
		_, err := songInteractor.GetAttendedSongs(-1)
		assert.Equal(t, dbError, err, "Error")
	})
}

func TestGetSongsInAlbum(t *testing.T) {
	testAlbum := domain.Album{ID: 100}
	testSong1 := domain.Song{ID: 1, Name: "Song1"}
	testSong2 := domain.Song{ID: 2, Name: "Song2"}
	testSongs := []domain.Song{testSong1, testSong2}
	emptySong := []domain.Song{}
	dbError := errors.New("DB error")
	SongRepositoryMock := new(SongRepositoryMock)
	SongRepositoryMock.On("GetSongsInAlbum", testAlbum.ID).Return(testSongs, nil).Once()
	SongRepositoryMock.On("GetSongsInAlbum", 0).Return(emptySong, nil).Once()
	SongRepositoryMock.On("GetSongsInAlbum", -1).Return(emptySong, dbError).Once()
	songInteractor := &SongInteractor{
		SongRepository: SongRepositoryMock,
	}
	t.Run("Normal Case", func(t *testing.T) {
		got, err := songInteractor.GetSongsInAlbum(testAlbum.ID)
		if err != nil {
			t.Error(err)
		}
		assert.ElementsMatch(t, testSongs, got, "Error")
	})
	t.Run("Error Case (Song does not exist)", func(t *testing.T) {
		_, err := songInteractor.GetSongsInAlbum(0)
		expectedError := errors.New("songs not found")
		assert.Equal(t, expectedError, err, "Error")
	})
	t.Run("Error Case (Repository returns error)", func(t *testing.T) {
		_, err := songInteractor.GetSongsInAlbum(-1)
		assert.Equal(t, dbError, err, "Error")
	})
}
