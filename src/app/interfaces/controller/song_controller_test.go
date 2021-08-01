package controller

import (
	"domain"
	"encoding/json"
	"errors"
	"net/http"
	"test"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type SongControllerTestSuite struct {
	suite.Suite
}

type SongInteractorMock struct {
	mock.Mock
}

func (_m *SongInteractorMock) GetSong(songId int) (domain.Song, error) {
	ret := _m.Called(songId)
	return ret.Get(0).(domain.Song), ret.Error(1)
}

func (_m *SongInteractorMock) GetAttendedSongs(artistId int) ([]domain.Song, error) {
	ret := _m.Called(artistId)
	return ret.Get(0).([]domain.Song), ret.Error(1)
}

func (_m *SongInteractorMock) GetSongsInAlbum(albumId int) ([]domain.Song, error) {
	ret := _m.Called(albumId)
	return ret.Get(0).([]domain.Song), ret.Error(1)
}

func TestSongControllerTestSuite(t *testing.T) {
	suite.Run(t, new(SongControllerTestSuite))
}

func (suite *SongControllerTestSuite) TestGetSongWhenFound() {
	// mockを設定, setup
	song := domain.Song{ID: 1, Name: "Test Song"}
	expectedJson, _ := json.Marshal(song)
	expected := string(expectedJson) + "\n"
	songInteractorMock := new(SongInteractorMock)
	songInteractorMock.On("GetSong", 1).Return(song, nil).Once()
	songController := SongController{
		Usecase: songInteractorMock,
	}

	// GETリクエストを投げる準備
	rec, c := test.CreateContextInstance("/songs/:id", "id", "1")
	handler := songController.GetSong()

	// 検証
	if assert.NoError(suite.T(), handler(c)) {
		assert.Equal(suite.T(), http.StatusOK, rec.Code)
		assert.Equal(suite.T(), expected, rec.Body.String())
	}
}

func (suite *SongControllerTestSuite) TestGetSongWhenNotFound() {
	// mockを設定, setup
	songInteractorMock := new(SongInteractorMock)
	songInteractorMock.On("GetSong", 1).Return(domain.Song{}, errors.New("song not found")).Once()
	songController := SongController{
		Usecase: songInteractorMock,
	}
	expectedJson, _ := json.Marshal(map[string]string{"Message": "Song Not Found"})
	expected := string(expectedJson) + "\n"

	// GETリクエストを投げる準備
	rec, c := test.CreateContextInstance("/songs/:id", "id", "1")
	handler := songController.GetSong()

	// 検証
	if assert.NoError(suite.T(), handler(c)) {
		assert.Equal(suite.T(), http.StatusNotFound, rec.Code)
		assert.Equal(suite.T(), expected, rec.Body.String())
	}
}

func (suite *SongControllerTestSuite) TestGetSongWhenInvalidId() {
	// mockを設定, setup
	songInteractorMock := new(SongInteractorMock)
	songController := SongController{
		Usecase: songInteractorMock,
	}
	expectedJson, _ := json.Marshal(map[string]string{"Message": "Invalid Parameter"})
	expected := string(expectedJson) + "\n"

	// GETリクエストを投げる準備
	rec, c := test.CreateContextInstance("/songs/:id", "id", "invalid")
	handler := songController.GetSong()

	// 検証
	if assert.NoError(suite.T(), handler(c)) {
		assert.Equal(suite.T(), http.StatusBadRequest, rec.Code)
		assert.Equal(suite.T(), expected, rec.Body.String())
	}
}

func (suite *SongControllerTestSuite) TestGetAttendedSongsWhenFound() {
	// mockを設定, setup
	songs := []domain.Song{domain.Song{ID: 1, Name: "Test Song"}}
	expectedJson, _ := json.Marshal(songs)
	expected := string(expectedJson) + "\n"
	songInteractorMock := new(SongInteractorMock)
	songInteractorMock.On("GetAttendedSongs", 1).Return(songs, nil).Once()
	songController := SongController{
		Usecase: songInteractorMock,
	}

	// GETリクエストを投げる準備
	rec, c := test.CreateContextInstance("/artists/:id/songs", "id", "1")
	handler := songController.GetAttendedSongs()

	// 検証
	if assert.NoError(suite.T(), handler(c)) {
		assert.Equal(suite.T(), http.StatusOK, rec.Code)
		assert.Equal(suite.T(), expected, rec.Body.String())
	}
}

func (suite *SongControllerTestSuite) TestGetAttendedSongsWhenNotFound() {
	// mockを設定, setup
	songInteractorMock := new(SongInteractorMock)
	songInteractorMock.On("GetAttendedSongs", 1).Return([]domain.Song{}, errors.New("songs not found")).Once()
	songController := SongController{
		Usecase: songInteractorMock,
	}
	expectedJson, _ := json.Marshal(map[string]string{"Message": "Songs Not Found"})
	expected := string(expectedJson) + "\n"

	// GETリクエストを投げる準備
	rec, c := test.CreateContextInstance("/artists/:id/songs", "id", "1")
	handler := songController.GetAttendedSongs()

	// 検証
	if assert.NoError(suite.T(), handler(c)) {
		assert.Equal(suite.T(), http.StatusNotFound, rec.Code)
		assert.Equal(suite.T(), expected, rec.Body.String())
	}
}

func (suite *SongControllerTestSuite) TestGetAttendedSongsWhenInvalidId() {
	// mockを設定, setup
	songInteractorMock := new(SongInteractorMock)
	songController := SongController{
		Usecase: songInteractorMock,
	}
	expectedJson, _ := json.Marshal(map[string]string{"Message": "Invalid Parameter"})
	expected := string(expectedJson) + "\n"

	// GETリクエストを投げる準備
	rec, c := test.CreateContextInstance("/artists/:id/songs", "id", "invalid")
	handler := songController.GetAttendedSongs()

	// 検証
	if assert.NoError(suite.T(), handler(c)) {
		assert.Equal(suite.T(), http.StatusBadRequest, rec.Code)
		assert.Equal(suite.T(), expected, rec.Body.String())
	}
}

func (suite *SongControllerTestSuite) TestGetSongsInAlbumWhenFound() {
	// mockを設定, setup
	songs := []domain.Song{domain.Song{ID: 1, Name: "Test Song"}}
	expectedJson, _ := json.Marshal(songs)
	expected := string(expectedJson) + "\n"
	songInteractorMock := new(SongInteractorMock)
	songInteractorMock.On("GetSongsInAlbum", 1).Return(songs, nil).Once()
	songController := SongController{
		Usecase: songInteractorMock,
	}

	// GETリクエストを投げる準備
	rec, c := test.CreateContextInstance("/albums/:id/songs", "id", "1")
	handler := songController.GetSongsInAlbum()

	// 検証
	if assert.NoError(suite.T(), handler(c)) {
		assert.Equal(suite.T(), http.StatusOK, rec.Code)
		assert.Equal(suite.T(), expected, rec.Body.String())
	}
}

func (suite *SongControllerTestSuite) TestGetSongsInAlbumWhenNotFound() {
	// mockを設定, setup
	songInteractorMock := new(SongInteractorMock)
	songInteractorMock.On("GetSongsInAlbum", 1).Return([]domain.Song{}, errors.New("songs not found")).Once()
	songController := SongController{
		Usecase: songInteractorMock,
	}
	expectedJson, _ := json.Marshal(map[string]string{"Message": "Songs Not Found"})
	expected := string(expectedJson) + "\n"

	// GETリクエストを投げる準備
	rec, c := test.CreateContextInstance("/albums/:id/songs", "id", "1")
	handler := songController.GetSongsInAlbum()

	// 検証
	if assert.NoError(suite.T(), handler(c)) {
		assert.Equal(suite.T(), http.StatusNotFound, rec.Code)
		assert.Equal(suite.T(), expected, rec.Body.String())
	}
}

func (suite *SongControllerTestSuite) TestGetSongsInAlbumWhenInvalidId() {
	// mockを設定, setup
	songInteractorMock := new(SongInteractorMock)
	songController := SongController{
		Usecase: songInteractorMock,
	}
	expectedJson, _ := json.Marshal(map[string]string{"Message": "Invalid Parameter"})
	expected := string(expectedJson) + "\n"

	// GETリクエストを投げる準備
	rec, c := test.CreateContextInstance("/albums/:id/songs", "id", "invalid")
	handler := songController.GetSongsInAlbum()

	// 検証
	if assert.NoError(suite.T(), handler(c)) {
		assert.Equal(suite.T(), http.StatusBadRequest, rec.Code)
		assert.Equal(suite.T(), expected, rec.Body.String())
	}
}
