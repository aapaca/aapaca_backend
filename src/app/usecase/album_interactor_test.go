package usecase

import (
	"domain"
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
	t.Run("1", func(t *testing.T) {
		rd, _ := time.Parse("2006-01-02", "2020-09-22")
		testAlbum := domain.Album{ID: 1, Name: "Test T. Test", PrimaryArtist: domain.Artist{ID: 100}, Label: "Test Label", ReleasedDate: &rd, ImageURL: "https://www.example.com"}
		albumRepositoryMock := new(AlbumRepositoryMock)
		albumRepositoryMock.On("GetAlbum", testAlbum.ID).Return(testAlbum, nil)
		albumInteractor := &AlbumInteractor{
			AlbumRepository: albumRepositoryMock,
		}
		got, err := albumInteractor.GetAlbum(1)
		if err != nil {
			t.Error(err)
		}
		assert.Equal(t, testAlbum, got, "Error")
	})
}

func TestGetAlbumsByAlbumId(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		rd, _ := time.Parse("2006-01-02", "2020-09-22")
		testArtist := domain.Artist{ID: 100}
		testAlbum1 := domain.Album{ID: 1, Name: "Test T. Test", PrimaryArtist: testArtist, Label: "Test Label", ReleasedDate: &rd, ImageURL: "https://www.example.com"}
		testAlbum2 := domain.Album{ID: 2, Name: "Test T. Test", PrimaryArtist: testArtist, Label: "Test Label", ReleasedDate: &rd, ImageURL: "https://www.example.com"}
		testAlbums := []domain.Album{testAlbum1, testAlbum2}
		albumRepositoryMock := new(AlbumRepositoryMock)
		albumRepositoryMock.On("GetAlbumsByArtistId", testArtist.ID).Return(testAlbums, nil)
		albumInteractor := &AlbumInteractor{
			AlbumRepository: albumRepositoryMock,
		}
		got, err := albumInteractor.GetAlbumsByArtistId(100)
		if err != nil {
			t.Error(err)
		}
		assert.ElementsMatch(t, testAlbums, got, "Error")
	})
}
