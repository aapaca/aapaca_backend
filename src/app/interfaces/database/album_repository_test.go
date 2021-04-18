package database

import (
	"domain"
	"interfaces/database/rdb"
	"sort"
	"test/infrastructure"
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
}

func (suite *GetAlbumTestSuite) SetupTest() {
	queries := []string{
		"INSERT INTO artists (name, status, image_url) VALUES ('Artist 1', 0, 'http://www.example.com')",
		"INSERT INTO artists (name, status, image_url) VALUES ('Artist 2', 0, 'http://www.example.com')",
		"INSERT INTO albums (name, primary_artist_id, label, released_date, image_url, description) VALUES('Album 1', 1, 'Label 1', '2021-01-13', 'http://www.example.com', 'This is test album 1');",
		"INSERT INTO albums (name, primary_artist_id, label, released_date, image_url, description) VALUES('Album 2', 2, 'Label 2', '2021-01-13', 'http://www.example.com', 'This is test album 2');",
		"INSERT INTO occupations (title) VALUE ('Part 1')",
		"INSERT INTO occupations (title) VALUE ('Part 2')",
		"INSERT INTO external_services (name) VALUE ('amazon_music')",
		"INSERT INTO external_services (name) VALUE ('apple_music')",
		"INSERT INTO external_services (name) VALUE ('spotify')",
		"INSERT INTO external_ids (record_id, record_type, external_id, service_id) VALUE (1, 'album', 'TEST1111', 1)",
		"INSERT INTO external_ids (record_id, record_type, external_id, service_id) VALUE (1, 'album', '1111', 2)",
		"INSERT INTO external_ids (record_id, record_type, external_id, service_id) VALUE (1, 'album', 'Test1111', 3)",
		"INSERT INTO external_ids (record_id, record_type, external_id, service_id) VALUE (2, 'album', 'Test2222', 3)",
		"INSERT INTO participations (artist_id, album_id, occupation_id) VALUES (1, 1, 1)",
		"INSERT INTO participations (artist_id, album_id, occupation_id) VALUES (2, 1, 1)",
		"INSERT INTO participations (artist_id, album_id, occupation_id) VALUES (2, 1, 2)",
	}
	for _, query := range queries {
		_, err := suite.sqlHandler.Execute(query)
		if err != nil {
			suite.T().Error(err)
		}
	}
	suite.albumRepository = AlbumRepository{
		SqlHandler: suite.sqlHandler,
	}
}

func (suite *GetAlbumTestSuite) TearDownTest() {
	err := infrastructure.DeleteAllRecords(suite.sqlHandler)
	if err != nil {
		suite.T().Error(err)
	}
}

func assertCredits(t *testing.T, expected []domain.Credit, got []domain.Credit) {
	assert.Equal(t, len(expected), len(got))
	// sort expected and got by ArtistID
	sort.Slice(expected, func(i, j int) bool {
		p, q := got[i], got[j]
		return p.Artist.ID < q.Artist.ID
	})
	sort.Slice(got, func(i, j int) bool {
		p, q := got[i], got[j]
		return p.Artist.ID < q.Artist.ID
	})
	for i := range expected {
		assert.Equal(t, expected[i].Artist, got[i].Artist)
		assert.ElementsMatch(t, expected[i].Parts, got[i].Parts)
	}
}

func (suite *GetAlbumTestSuite) TestGetAlbum() {
	testURL := "http://www.example.com"
	testDate, _ := time.Parse("2006-01-02", "2021-01-13")
	testArtist1 := domain.Artist{ID: 1, Name: "Artist 1", ImageURL: testURL}
	testArtist2 := domain.Artist{ID: 2, Name: "Artist 2", ImageURL: testURL}
	testPart1 := domain.Occupation{ID: 1, Title: "Part 1"}
	testPart2 := domain.Occupation{ID: 2, Title: "Part 2"}
	expectedAlbum := domain.Album{
		ID:            1,
		Name:          "Album 1",
		PrimaryArtist: testArtist1,
		Credits: []domain.Credit{
			{Artist: testArtist1, Parts: []domain.Occupation{testPart1}},
			{Artist: testArtist2, Parts: []domain.Occupation{testPart1, testPart2}},
		},
		Label:        "Label 1",
		ReleasedDate: &testDate,
		ImageURL:     testURL,
		Description:  "This is test album 1",
		Links: map[string]string{
			"amazonMusic": "https://www.amazon.com/dp/TEST1111",
			"appleMusic":  "https://music.apple.com/album/1111",
			"spotify":     "https://open.spotify.com/album/Test1111",
		},
	}
	album, err := suite.albumRepository.GetAlbum(1)
	if err != nil {
		suite.T().Error(err)
	}
	assert.Equal(suite.T(), expectedAlbum.ID, album.ID)
	assert.Equal(suite.T(), expectedAlbum.Name, album.Name)
	assert.Equal(suite.T(), expectedAlbum.PrimaryArtist, album.PrimaryArtist)
	assertCredits(suite.T(), expectedAlbum.Credits, album.Credits)
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
	expectedAlbum := domain.Album{
		ID:            2,
		Name:          "Album 2",
		PrimaryArtist: testArtist2,
		Label:         "Label 2",
		ReleasedDate:  &testDate,
		ImageURL:      testURL,
		Description:   "This is test album 2",
		Links: map[string]string{
			"spotify": "https://open.spotify.com/album/Test2222",
		},
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
}

func (suite *GetAlbumByArtistIdTestSuite) SetupTest() {
	queries := []string{
		"INSERT INTO artists (name, status) VALUES ('Artist 1', 0);",
		"INSERT INTO artists (name, status) VALUES ('Artist 2', 0);",
		"INSERT INTO albums (name, primary_artist_id, label, released_date, image_url, description) VALUES('Album 1', 1, 'Label 1', '1999-07-13', 'http://www.example.com', 'This is test album 1');",
		"INSERT INTO albums (name, primary_artist_id, label, released_date, image_url) VALUES('Album 2', 1, 'Label 2', '2021-01-13', 'http://www.example.com');",
	}
	for _, query := range queries {
		_, err := suite.sqlHandler.Execute(query)
		if err != nil {
			suite.T().Error(err)
		}
	}
	suite.albumRepository = AlbumRepository{
		SqlHandler: suite.sqlHandler,
	}
}

func (suite *GetAlbumByArtistIdTestSuite) TearDownTest() {
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
