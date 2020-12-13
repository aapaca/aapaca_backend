package usecase

import (
	"domain"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type ArtistRepositoryMock struct {
	mock.Mock
}

func (_m *ArtistRepositoryMock) GetArtist(artistId int) (domain.Artist, error) {
	ret := _m.Called(artistId)
	return ret.Get(0).(domain.Artist), ret.Error(1)
}

func TestGetArtist(t *testing.T) {
	testArtist := domain.Artist{ID: 1, Name: "Test Artist"}
	emptyArtist := domain.Artist{}
	dbError := errors.New("DB error")
	artistRepositoryMock := new(ArtistRepositoryMock)
	artistRepositoryMock.On("GetArtist", testArtist.ID).Return(testArtist, nil).Once()
	artistRepositoryMock.On("GetArtist", 0).Return(emptyArtist, nil)
	artistRepositoryMock.On("GetArtist", -1).Return(emptyArtist, dbError)
	artistInteractor := &ArtistInteractor{
		ArtistRepository: artistRepositoryMock,
	}
	t.Run("Normal Case", func(t *testing.T) {
		got, err := artistInteractor.GetArtist(1)
		if err != nil {
			t.Error(err)
		}
		assert.Equal(t, testArtist, got, "Error")
	})
	t.Run("Error Case (Artist does not exist)", func(t *testing.T) {
		_, err := artistInteractor.GetArtist(0)
		expectedError := errors.New("artist not found")
		assert.Equal(t, expectedError, err, "Error")
	})
	t.Run("Error Case (Repository returns error)", func(t *testing.T) {
		_, err := artistInteractor.GetArtist(-1)
		assert.Equal(t, dbError, err, "Error")
	})
}
