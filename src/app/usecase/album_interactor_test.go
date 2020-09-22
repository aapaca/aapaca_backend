package usecase

import (
	"domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

type AlbumRepositoryMock struct {
	mock.Mock
}

func (_m *AlbumRepositoryMock) FindById(id int) (domain.Album, error) {
	ret := _m.Called(id)
	return ret.Get(0).(domain.Album), ret.Error(1)
}

func TestAlbumById(t *testing.T) {
	rd, _ := time.Parse("2006-01-02", "2020-09-22")
	testAlbum := domain.Album{ID: 1, Name: "Test T. Test", PrimaryArtist: domain.Artist{ID: 100}, Label: "Test Label", ReleasedDate: &rd, ImageURL: "https://www.example.com"}
	albumRepositoryMock := new(AlbumRepositoryMock)
	albumRepositoryMock.On("FindById", testAlbum.ID).Return(testAlbum, nil)
	albumInteractor := &AlbumInteractor{
		AlbumRepository: albumRepositoryMock,
	}
	got, err := albumInteractor.AlbumById(1)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, testAlbum, got, "Error")
}
