package interactor

import (
	"domain"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type ArtistRepositoryMock struct {
	mock.Mock
}

func (_m *ArtistRepositoryMock) GetArtist(artistId int) (domain.Artist, error) {
	ret := _m.Called(artistId)
	return ret.Get(0).(domain.Artist), ret.Error(1)
}

func (_m *ArtistRepositoryMock) FindMembers(artistId int) ([]domain.Credit, error) {
	ret := _m.Called(artistId)
	return ret.Get(0).([]domain.Credit), ret.Error(1)
}

func (_m *ArtistRepositoryMock) FindAliases(artistId int) ([]domain.Credit, error) {
	ret := _m.Called(artistId)
	return ret.Get(0).([]domain.Credit), ret.Error(1)
}

type GetArtistTestSuite struct {
	suite.Suite
}

func TestGetArtistTestSuite(t *testing.T) {
	suite.Run(t, new(GetArtistTestSuite))
}

func (suite *GetArtistTestSuite) TestCallAllFunctions() {
	artist := domain.Artist{ID: 1, Name: "Test Artist"}
	members := []domain.Credit{domain.Credit{Artist: &domain.Artist{ID: 2}}}
	aliases := []domain.Credit{domain.Credit{Artist: &domain.Artist{ID: 3}}}
	artistRepositoryMock := new(ArtistRepositoryMock)
	artistRepositoryMock.On("GetArtist", artist.ID).Return(artist, nil).Once()
	artistRepositoryMock.On("FindMembers", artist.ID).Return(members, nil).Once()
	artistRepositoryMock.On("FindAliases", artist.ID).Return(aliases, nil).Once()
	artistInteractor := &ArtistInteractor{
		ArtistRepository: artistRepositoryMock,
	}
	expect := domain.Artist{ID: 1, Name: "Test Artist", Members: members, Aliases: aliases}

	actual, err := artistInteractor.GetArtist(1)

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expect, actual)
}

func (suite *GetArtistTestSuite) TestArtistNotExist() {
	artist := domain.Artist{ID: 0, Name: ""}
	artistRepositoryMock := new(ArtistRepositoryMock)
	artistRepositoryMock.On("GetArtist", 1).Return(artist, nil).Once()
	artistInteractor := &ArtistInteractor{
		ArtistRepository: artistRepositoryMock,
	}
	expectError := errors.New("artist not found")

	_, err := artistInteractor.GetArtist(1)

	assert.Equal(suite.T(), expectError, err)
}

func (suite *GetArtistTestSuite) TestWhenMembersAndAlisesAreEmpty() {
	artist := domain.Artist{ID: 1, Name: "Test Artist"}
	emptyMember := []domain.Credit{}
	emptyAlias := []domain.Credit{}
	artistRepositoryMock := new(ArtistRepositoryMock)
	artistRepositoryMock.On("GetArtist", artist.ID).Return(artist, nil).Once()
	artistRepositoryMock.On("FindMembers", artist.ID).Return(emptyMember, nil).Once()
	artistRepositoryMock.On("FindAliases", artist.ID).Return(emptyAlias, nil).Once()
	artistInteractor := &ArtistInteractor{
		ArtistRepository: artistRepositoryMock,
	}
	expect := domain.Artist{ID: 1, Name: "Test Artist"}

	actual, err := artistInteractor.GetArtist(1)

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expect, actual)
}

func (suite *GetArtistTestSuite) TestGetArtistReturnsDbError() {
	artist := domain.Artist{}
	expectError := errors.New("DB Error")
	artistRepositoryMock := new(ArtistRepositoryMock)
	artistRepositoryMock.On("GetArtist", 1).Return(artist, expectError).Once()
	artistInteractor := &ArtistInteractor{
		ArtistRepository: artistRepositoryMock,
	}

	_, err := artistInteractor.GetArtist(1)

	assert.Equal(suite.T(), expectError, err)
}

func (suite *GetArtistTestSuite) TestFindMembersReturnsDbError() {
	artist := domain.Artist{ID: 1, Name: "Test Artist"}
	members := []domain.Credit{}
	expectError := errors.New("DB Error")
	artistRepositoryMock := new(ArtistRepositoryMock)
	artistRepositoryMock.On("GetArtist", 1).Return(artist, nil).Once()
	artistRepositoryMock.On("FindMembers", 1).Return(members, expectError).Once()
	artistInteractor := &ArtistInteractor{
		ArtistRepository: artistRepositoryMock,
	}

	_, err := artistInteractor.GetArtist(1)

	assert.Equal(suite.T(), expectError, err)
}

func (suite *GetArtistTestSuite) TestFindAliasesReturnsDbError() {
	artist := domain.Artist{ID: 1, Name: "Test Artist"}
	members := []domain.Credit{}
	aliases := []domain.Credit{}
	expectError := errors.New("DB Error")
	artistRepositoryMock := new(ArtistRepositoryMock)
	artistRepositoryMock.On("GetArtist", 1).Return(artist, nil).Once()
	artistRepositoryMock.On("FindMembers", 1).Return(members, nil).Once()
	artistRepositoryMock.On("FindAliases", 1).Return(aliases, expectError).Once()
	artistInteractor := &ArtistInteractor{
		ArtistRepository: artistRepositoryMock,
	}

	_, err := artistInteractor.GetArtist(1)

	assert.Equal(suite.T(), expectError, err)
}
