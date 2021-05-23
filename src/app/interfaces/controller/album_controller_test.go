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

type AlbumControllerTestSuite struct {
	suite.Suite
	controller AlbumController
}

type AlbumInteractorMock struct {
	mock.Mock
}

func (_m *AlbumInteractorMock) GetAlbum(albumId int) (domain.Album, error) {
	ret := _m.Called(albumId)
	return ret.Get(0).(domain.Album), ret.Error(1)
}

func (_m *AlbumInteractorMock) GetAlbumsByArtistId(albumId int) ([]domain.Album, error) {
	ret := _m.Called(albumId)
	return ret.Get(0).([]domain.Album), ret.Error(1)
}

func TestAlbumControllerTestSuite(t *testing.T) {
	suite.Run(t, new(AlbumControllerTestSuite))
}

func (suite *AlbumControllerTestSuite) TestGetAlbumWhenFound() {
	// mockを設定, setup
	album := domain.Album{ID: 1, Name: "Test Album"}
	expectedJson, _ := json.Marshal(album)
	expected := string(expectedJson) + "\n"
	albumInteractorMock := new(AlbumInteractorMock)
	albumInteractorMock.On("GetAlbum", 1).Return(album, nil).Once()
	albumController := AlbumController{
		Interactor: albumInteractorMock,
	}

	// GETリクエストを投げる準備
	rec, c := test.CreateContextInstance("/albums/:id", "id", "1")
	handler := albumController.GetAlbum()

	// 検証
	if assert.NoError(suite.T(), handler(c)) {
		assert.Equal(suite.T(), http.StatusOK, rec.Code)
		assert.Equal(suite.T(), expected, rec.Body.String())
	}
}

func (suite *AlbumControllerTestSuite) TestGetAlbumWhenNotFound() {
	// mockを設定, setup
	albumInteractorMock := new(AlbumInteractorMock)
	albumInteractorMock.On("GetAlbum", 1).Return(domain.Album{}, errors.New("album not found")).Once()
	albumController := AlbumController{
		Interactor: albumInteractorMock,
	}
	expectedJson, _ := json.Marshal(map[string]string{"Message": "Album Not Found"})
	expected := string(expectedJson) + "\n"

	// GETリクエストを投げる準備
	rec, c := test.CreateContextInstance("/albums/:id", "id", "1")
	handler := albumController.GetAlbum()

	// 検証
	if assert.NoError(suite.T(), handler(c)) {
		assert.Equal(suite.T(), http.StatusBadRequest, rec.Code)
		assert.Equal(suite.T(), expected, rec.Body.String())
	}
}

func (suite *AlbumControllerTestSuite) TestGetAlbumWhenInvalidId() {
	// mockを設定, setup
	albumInteractorMock := new(AlbumInteractorMock)
	albumController := AlbumController{
		Interactor: albumInteractorMock,
	}
	expectedJson, _ := json.Marshal(map[string]string{"Message": "Invalid Parameter"})
	expected := string(expectedJson) + "\n"

	// GETリクエストを投げる準備
	rec, c := test.CreateContextInstance("/albums/:id", "id", "invalid")
	handler := albumController.GetAlbum()

	// 検証
	if assert.NoError(suite.T(), handler(c)) {
		assert.Equal(suite.T(), http.StatusBadRequest, rec.Code)
		assert.Equal(suite.T(), expected, rec.Body.String())
	}
}

func (suite *AlbumControllerTestSuite) TestGetAlbumByArtistIdWhenFound() {
	// mockを設定, setup
	albums := []domain.Album{domain.Album{ID: 1, Name: "Test Album"}}
	expectedJson, _ := json.Marshal(albums)
	expected := string(expectedJson) + "\n"
	albumInteractorMock := new(AlbumInteractorMock)
	albumInteractorMock.On("GetAlbumsByArtistId", 1).Return(albums, nil).Once()
	albumController := AlbumController{
		Interactor: albumInteractorMock,
	}

	// GETリクエストを投げる準備
	rec, c := test.CreateContextInstance("/artists/:id/albums", "id", "1")
	handler := albumController.GetAlbumsByArtistId()

	// 検証
	if assert.NoError(suite.T(), handler(c)) {
		assert.Equal(suite.T(), http.StatusOK, rec.Code)
		assert.Equal(suite.T(), expected, rec.Body.String())
	}
}

func (suite *AlbumControllerTestSuite) TestGetAlbumByArtistIdWhenNotFound() {
	// mockを設定, setup
	albumInteractorMock := new(AlbumInteractorMock)
	albumInteractorMock.On("GetAlbumsByArtistId", 1).Return([]domain.Album{}, errors.New("album not found")).Once()
	albumController := AlbumController{
		Interactor: albumInteractorMock,
	}
	expectedJson, _ := json.Marshal(map[string]string{"Message": "Albums Not Found"})
	expected := string(expectedJson) + "\n"

	// GETリクエストを投げる準備
	rec, c := test.CreateContextInstance("/artists/:id/albums", "id", "1")
	handler := albumController.GetAlbumsByArtistId()

	// 検証
	if assert.NoError(suite.T(), handler(c)) {
		assert.Equal(suite.T(), http.StatusBadRequest, rec.Code)
		assert.Equal(suite.T(), expected, rec.Body.String())
	}
}

func (suite *AlbumControllerTestSuite) TestGetAlbumByArtistIdWhenInvalidId() {
	// mockを設定, setup
	albumInteractorMock := new(AlbumInteractorMock)
	albumController := AlbumController{
		Interactor: albumInteractorMock,
	}
	expectedJson, _ := json.Marshal(map[string]string{"Message": "Invalid Parameter"})
	expected := string(expectedJson) + "\n"

	// GETリクエストを投げる準備
	rec, c := test.CreateContextInstance("/artists/:id/albums", "id", "invalid")
	handler := albumController.GetAlbumsByArtistId()

	// 検証
	if assert.NoError(suite.T(), handler(c)) {
		assert.Equal(suite.T(), http.StatusBadRequest, rec.Code)
		assert.Equal(suite.T(), expected, rec.Body.String())
	}
}
