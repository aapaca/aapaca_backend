package repository

import (
	"domain"
	"interfaces/repository/rdb"
	"test/infrastructure"
	"test/interfaces/repository"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type GetAlbumTestSuite struct {
	suite.Suite
	sqlHandler      rdb.SqlHandler
	albumRepository AlbumRepository
}

func TestGetAlbumTestSuite(t *testing.T) {
	suite.Run(t, new(GetAlbumTestSuite))
}

func (suite *GetAlbumTestSuite) SetupSuite() {
	suite.sqlHandler = infrastructure.NewSqlHandler()

	err := repository.InitDb("testdata/get_album_init.sql", suite.sqlHandler)
	if err != nil {
		suite.T().Error(err)
	}
	suite.albumRepository = AlbumRepository{
		SqlHandler: suite.sqlHandler,
	}
}

func (suite *GetAlbumTestSuite) TearDownSuite() {
	err := infrastructure.DeleteAllRecords(suite.sqlHandler)
	if err != nil {
		suite.T().Error(err)
	}
}

func (suite *GetAlbumTestSuite) TestGetAlbum() {
	testURL := "http://www.example.com"
	testDate, _ := time.Parse("2006-01-02", "2021-01-13")
	testArtist1 := domain.Artist{ID: 1, Name: "Artist 1", ImageURL: testURL}
	testArtist2 := domain.Artist{ID: 2, Name: "Artist 2", ImageURL: testURL}
	testPart1 := domain.Occupation{ID: 1, Title: "Part 1"}
	testPart2 := domain.Occupation{ID: 2, Title: "Part 2"}
	testParts1 := domain.NewOccupations()
	testParts1.Append(testPart1)
	testParts2 := domain.NewOccupations()
	testParts2.Append(testPart1)
	testParts2.Append(testPart2)
	links := domain.NewAlbumLinks()
	links.AddLink("TEST1111", "amazon_music")
	links.AddLink("1111", "apple_music")
	links.AddLink("Test1111", "spotify")
	expectedAlbum := domain.Album{
		ID:            1,
		Name:          "Album 1",
		PrimaryArtist: testArtist1,
		Credits: []domain.Credit{
			{Artist: &testArtist1, Parts: testParts1},
			{Artist: &testArtist2, Parts: testParts2},
		},
		Label:        "Label 1",
		ReleasedDate: &testDate,
		ImageURL:     testURL,
		Description:  "This is test album 1",
		Links:        links,
	}
	album, err := suite.albumRepository.GetAlbum(1)
	if err != nil {
		suite.T().Error(err)
	}
	assert.Equal(suite.T(), expectedAlbum.ID, album.ID)
	assert.Equal(suite.T(), expectedAlbum.Name, album.Name)
	assert.Equal(suite.T(), expectedAlbum.PrimaryArtist, album.PrimaryArtist)
	repository.AssertCredits(suite.T(), expectedAlbum.Credits, album.Credits)
	assert.Equal(suite.T(), expectedAlbum.Label, album.Label)
	assert.Equal(suite.T(), expectedAlbum.ReleasedDate, album.ReleasedDate)
	assert.Equal(suite.T(), expectedAlbum.ImageURL, album.ImageURL)
	assert.Equal(suite.T(), expectedAlbum.Description, album.Description)
	assert.Equal(suite.T(), expectedAlbum.Links, album.Links)
}

func (suite *GetAlbumTestSuite) TestGetAlbumNoCredit() {
	testURL := "http://www.example.com"
	testDate, _ := time.Parse("2006-01-02", "2021-01-13")
	testArtist2 := domain.Artist{ID: 2, Name: "Artist 2", ImageURL: testURL}
	links := domain.NewAlbumLinks()
	links.AddLink("Test2222", "spotify")
	expectedAlbum := domain.Album{
		ID:            2,
		Name:          "Album 2",
		PrimaryArtist: testArtist2,
		Label:         "Label 2",
		ReleasedDate:  &testDate,
		ImageURL:      testURL,
		Description:   "This is test album 2",
		Links:         links,
	}
	album, err := suite.albumRepository.GetAlbum(2)
	if err != nil {
		suite.T().Error(err)
	}
	assert.Equal(suite.T(), expectedAlbum, album)
}

func (suite *GetAlbumTestSuite) TestGetAlbumInvalidID() {
	emptyAlbum := domain.Album{}
	album, err := suite.albumRepository.GetAlbum(100)
	if err != nil {
		suite.T().Error(err)
	}
	assert.Equal(suite.T(), emptyAlbum, album)
}

type GetAlbumByArtistIdTestSuite struct {
	suite.Suite
	sqlHandler      rdb.SqlHandler
	albumRepository AlbumRepository
}

func TestGetAlbumByArtistIdTestSuite(t *testing.T) {
	suite.Run(t, new(GetAlbumByArtistIdTestSuite))
}

func (suite *GetAlbumByArtistIdTestSuite) SetupSuite() {
	suite.sqlHandler = infrastructure.NewSqlHandler()

	err := repository.InitDb("testdata/get_album_by_artist_id_init.sql", suite.sqlHandler)
	if err != nil {
		suite.T().Error(err)
	}
	suite.albumRepository = AlbumRepository{
		SqlHandler: suite.sqlHandler,
	}
}

func (suite *GetAlbumByArtistIdTestSuite) TearDownSuite() {
	err := infrastructure.DeleteAllRecords(suite.sqlHandler)
	if err != nil {
		suite.T().Error(err)
	}
}

func (suite *GetAlbumByArtistIdTestSuite) TestGetAlbumByArtistId() {
	testDate1, _ := time.Parse("2006-01-02", "1999-07-13")
	testDate2, _ := time.Parse("2006-01-02", "2021-01-13")
	testURL := "http://www.example.com"
	testAlbum1 := domain.Album{ID: 1, Name: "Album 1", ReleasedDate: &testDate1, ImageURL: testURL}
	testAlbum2 := domain.Album{ID: 2, Name: "Album 2", ReleasedDate: &testDate2, ImageURL: testURL}
	expectedAlbums := []domain.Album{testAlbum1, testAlbum2}
	albums, err := suite.albumRepository.GetAlbumsByArtistId(1)
	if err != nil {
		suite.T().Error(err)
	}
	assert.ElementsMatch(suite.T(), expectedAlbums, albums)
}

func (suite *GetAlbumByArtistIdTestSuite) TestGetAlbumByArtistIdNoAlbum() {
	albums, err := suite.albumRepository.GetAlbumsByArtistId(2)
	if err != nil {
		suite.T().Error(err)
	}
	assert.Equal(suite.T(), 0, len(albums))
}

func (suite *GetAlbumByArtistIdTestSuite) TestGetAlbumByArtistIdInvalidId() {
	albums, err := suite.albumRepository.GetAlbumsByArtistId(100)
	if err != nil {
		suite.T().Error(err)
	}
	assert.Equal(suite.T(), 0, len(albums))
}
