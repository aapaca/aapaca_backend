package usecase

import (
	"domain"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type AlbumRepositoryMock struct {
	mock.Mock
}

func (_m *AlbumRepositoryMock) GetAlbum(albumId int) (domain.Album, error) {
	ret := _m.Called(albumId)
	return ret.Get(0).(domain.Album), ret.Error(1)
}

func (_m *AlbumRepositoryMock) GetAlbumsByArtistId(artistId int) ([]domain.Album, error) {
	ret := _m.Called(artistId)
	return ret.Get(0).([]domain.Album), ret.Error(1)
}

func TestGetAlbum(t *testing.T) {
	rd, _ := time.Parse("2006-01-02", "2020-09-22")
	emptyAlbum := domain.Album{}                                                                                                                                               //for error case
	testAlbum := domain.Album{ID: 1, Name: "Test T. Test", PrimaryArtist: domain.Artist{ID: 100}, Label: "Test Label", ReleasedDate: &rd, ImageURL: "https://www.example.com"} //for normal case
	dbError := errors.New("DB error")
	albumRepositoryMock := new(AlbumRepositoryMock)
	albumRepositoryMock.On("GetAlbum", testAlbum.ID).Return(testAlbum, nil).Once()
	albumRepositoryMock.On("GetAlbum", 0).Return(emptyAlbum, nil).Once()
	albumRepositoryMock.On("GetAlbum", -1).Return(emptyAlbum, dbError).Once()
	albumInteractor := &AlbumInteractor{
		AlbumRepository: albumRepositoryMock,
	}
	t.Run("Normal Case", func(t *testing.T) {
		got, err := albumInteractor.GetAlbum(1)
		if err != nil {
			t.Error(err)
		}
		assert.Equal(t, testAlbum, got, "Error")
	})
	t.Run("Error Case (Album does not exist)", func(t *testing.T) {
		_, err := albumInteractor.GetAlbum(0)
		expectedError := errors.New("album not found")
		assert.Equal(t, expectedError, err, "Error")
	})
	t.Run("Error Case (Repository returns error)", func(t *testing.T) {
		_, err := albumInteractor.GetAlbum(-1)
		assert.Equal(t, dbError, err, "Error")
	})
}

func TestGetAlbumsByAlbumId(t *testing.T) {
	rd, _ := time.Parse("2006-01-02", "2020-09-22")
	testArtist := domain.Artist{ID: 100}
	testAlbum1 := domain.Album{ID: 1, Name: "Test T. Test", PrimaryArtist: testArtist, Label: "Test Label", ReleasedDate: &rd, ImageURL: "https://www.example.com"}
	testAlbum2 := domain.Album{ID: 2, Name: "Test T. Test", PrimaryArtist: testArtist, Label: "Test Label", ReleasedDate: &rd, ImageURL: "https://www.example.com"}
	dbError := errors.New("DB error")
	emptyAlbums := []domain.Album{}
	testAlbums := []domain.Album{testAlbum1, testAlbum2}
	albumRepositoryMock := new(AlbumRepositoryMock)
	albumRepositoryMock.On("GetAlbumsByArtistId", testArtist.ID).Return(testAlbums, nil).Once() // for normal case
	albumRepositoryMock.On("GetAlbumsByArtistId", 0).Return(emptyAlbums, nil).Once()
	albumRepositoryMock.On("GetAlbumsByArtistId", -1).Return(emptyAlbums, dbError).Once()
	albumInteractor := &AlbumInteractor{
		AlbumRepository: albumRepositoryMock,
	}
	t.Run("Normal Case", func(t *testing.T) {
		got, err := albumInteractor.GetAlbumsByArtistId(100)
		if err != nil {
			t.Error(err)
		}
		assert.ElementsMatch(t, testAlbums, got, "Error")
	})
	t.Run("Error Case (Artist has no albums)", func(t *testing.T) {
		_, err := albumInteractor.GetAlbumsByArtistId(0)
		expectedError := errors.New("albums not found")
		assert.Equal(t, expectedError, err, "Error")
	})
	t.Run("Error Case (Repository returns error)", func(t *testing.T) {
		_, err := albumInteractor.GetAlbumsByArtistId(-1)
		assert.Equal(t, dbError, err, "Error")
	})
}
