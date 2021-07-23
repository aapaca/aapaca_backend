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

type ArtistControllerTestSuite struct {
	suite.Suite
}

type ArtistInteractorMock struct {
	mock.Mock
}

func (_m *ArtistInteractorMock) GetArtist(artistId int) (domain.Artist, error) {
	ret := _m.Called(artistId)
	return ret.Get(0).(domain.Artist), ret.Error(1)
}

func TestArtistControllerTestSuite(t *testing.T) {
	suite.Run(t, new(ArtistControllerTestSuite))
}

func (suite *ArtistControllerTestSuite) TestGetArtistWhenFound() {
	// mockを設定, setup
	artist := domain.Artist{ID: 1, Name: "Test Artist"}
	expectedJson, _ := json.Marshal(artist)
	expected := string(expectedJson) + "\n"
	artistInteractorMock := new(ArtistInteractorMock)
	artistInteractorMock.On("GetArtist", 1).Return(artist, nil).Once()
	artistController := ArtistController{
		Interactor: artistInteractorMock,
	}

	// GETリクエストを投げる準備
	rec, c := test.CreateContextInstance("/artists/:id", "id", "1")
	handler := artistController.GetArtist()

	// 検証
	if assert.NoError(suite.T(), handler(c)) {
		assert.Equal(suite.T(), http.StatusOK, rec.Code)
		assert.Equal(suite.T(), expected, rec.Body.String())
	}
}

func (suite *ArtistControllerTestSuite) TestGetArtistWhenNotFound() {
	// mockを設定, setup
	artistInteractorMock := new(ArtistInteractorMock)
	artistInteractorMock.On("GetArtist", 1).Return(domain.Artist{}, errors.New("artist not found")).Once()
	artistController := ArtistController{
		Interactor: artistInteractorMock,
	}
	expectedJson, _ := json.Marshal(map[string]string{"Message": "Artist Not Found"})
	expected := string(expectedJson) + "\n"

	// GETリクエストを投げる準備
	rec, c := test.CreateContextInstance("/artists/:id", "id", "1")
	handler := artistController.GetArtist()

	// 検証
	if assert.NoError(suite.T(), handler(c)) {
		assert.Equal(suite.T(), http.StatusNotFound, rec.Code)
		assert.Equal(suite.T(), expected, rec.Body.String())
	}
}

func (suite *ArtistControllerTestSuite) TestGetArtistWhenInvalidId() {
	// mockを設定, setup
	artistInteractorMock := new(ArtistInteractorMock)
	artistController := ArtistController{
		Interactor: artistInteractorMock,
	}
	expectedJson, _ := json.Marshal(map[string]string{"Message": "Invalid Parameter"})
	expected := string(expectedJson) + "\n"

	// GETリクエストを投げる準備
	rec, c := test.CreateContextInstance("/artists/:id", "id", "invalid")
	handler := artistController.GetArtist()

	// 検証
	if assert.NoError(suite.T(), handler(c)) {
		assert.Equal(suite.T(), http.StatusBadRequest, rec.Code)
		assert.Equal(suite.T(), expected, rec.Body.String())
	}
}
