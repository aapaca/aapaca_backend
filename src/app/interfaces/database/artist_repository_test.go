package database

import (
	"domain"
	"interfaces/database/rdb"
	"test/infrastructure"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type GetArtistTestSuite struct {
	suite.Suite
	sqlHandler       rdb.SqlHandler
	artistRepository ArtistRepository
}

func TestGetArtistTestSuite(t *testing.T) {
	suite.Run(t, new(GetArtistTestSuite))
}

func (suite *GetArtistTestSuite) SetupSuite() {
	suite.sqlHandler = infrastructure.NewSqlHandler()
}

func (suite *GetArtistTestSuite) SetupTest() {
	queries := []string{
		"INSERT INTO artists (name, status, image_url, description) VALUES ('Artist 1', 0, 'http://www.example.com', 'This is test artist 1')",
		"INSERT INTO artists (name, status, image_url) VALUES ('Artist 2', 0, 'http://www.example.com')",
		"INSERT INTO artists (name, status, image_url) VALUES ('Alias Artist 2', 0, 'http://www.example.com')",
		"INSERT INTO artists (name, status, image_url, description) VALUES ('Group Artist 1', 1, 'http://www.example.com', 'This is test group artist 1')",
		"INSERT INTO artists (name, status, image_url) VALUES ('Group Artist 2', 1, 'http://www.example.com')",
		"INSERT INTO artists (name, status, image_url, description) VALUES ('Alias Group Artist 2', 1, 'http://www.example.com', 'This is test alias group artist 2')",
		"INSERT INTO aliases (artist_id, alias_artist_id) VALUES (2, 3);",
		"INSERT INTO aliases (artist_id, alias_artist_id) VALUES (3, 2);",
		"INSERT INTO aliases (artist_id, alias_artist_id) VALUES (5, 6);",
		"INSERT INTO aliases (artist_id, alias_artist_id) VALUES (6, 5);",
		"INSERT INTO memberships (member_id, group_id) VALUES (1, 4);",
		"INSERT INTO memberships (member_id, group_id) VALUES (2, 4);",
		"INSERT INTO memberships (member_id, group_id) VALUES (1, 5);",
		"INSERT INTO memberships (member_id, group_id) VALUES (2, 5);",
		"INSERT INTO memberships (member_id, group_id) VALUES (4, 5);",
		"INSERT INTO external_services (name) VALUE ('amazon_music')",
		"INSERT INTO external_services (name) VALUE ('apple_music')",
		"INSERT INTO external_services (name) VALUE ('spotify')",
		"INSERT INTO external_ids (record_id, record_type, external_id, service_id) VALUE (1, 'artist', 'TEST1111', 1)",
		"INSERT INTO external_ids (record_id, record_type, external_id, service_id) VALUE (1, 'artist', '1111', 2)",
		"INSERT INTO external_ids (record_id, record_type, external_id, service_id) VALUE (1, 'artist', 'Test1111', 3)",
		"INSERT INTO external_ids (record_id, record_type, external_id, service_id) VALUE (2, 'artist', 'Test2222', 3)",
		"INSERT INTO songs (name, primary_artist_id) VALUES('Song 1', 1);",
		"INSERT INTO songs (name, primary_artist_id) VALUES('Song 2', 4);",
		"INSERT INTO albums (name, primary_artist_id) VALUES('Album 1', 1);",
		"INSERT INTO albums (name, primary_artist_id) VALUES('Album 2', 4);",
		"INSERT INTO occupations (title) VALUES ('Part 1')",
		"INSERT INTO occupations (title) VALUES ('Part 2')",
		"INSERT INTO occupations (title) VALUES ('Part 3')",
		"INSERT INTO participations (artist_id, album_id, occupation_id) VALUES (1, 1, 1)",
		"INSERT INTO participations (artist_id, album_id, occupation_id) VALUES (1, 2, 1)",
		"INSERT INTO participations (artist_id, album_id, occupation_id) VALUES (4, 2, 1)",
		"INSERT INTO participations (artist_id, album_id, occupation_id) VALUES (4, 2, 2)",
		"INSERT INTO performances (artist_id, song_id, occupation_id) VALUES (1, 1, 2)",
		"INSERT INTO performances (artist_id, song_id, occupation_id) VALUES (1, 2, 2)",
		"INSERT INTO performances (artist_id, song_id, occupation_id) VALUES (2, 1, 2)",
		"INSERT INTO performances (artist_id, song_id, occupation_id) VALUES (2, 2, 2)",
		"INSERT INTO performances (artist_id, song_id, occupation_id) VALUES (3, 1, 3)",
		"INSERT INTO performances (artist_id, song_id, occupation_id) VALUES (6, 2, 3)",
	}
	for _, query := range queries {
		_, err := suite.sqlHandler.Execute(query)
		if err != nil {
			suite.T().Error(err)
		}
	}
	suite.artistRepository = ArtistRepository{
		SqlHandler: suite.sqlHandler,
	}
}

func (suite *GetArtistTestSuite) TearDownTest() {
	err := infrastructure.DeleteAllRecords(suite.sqlHandler)
	if err != nil {
		suite.T().Error(err)
	}
}

func (suite *GetArtistTestSuite) TestGetArtist() {
	testURL := "http://www.example.com"
	testPart1 := domain.Occupation{ID: 1, Title: "Part 1"}
	testPart2 := domain.Occupation{ID: 2, Title: "Part 2"}
	expectedArtist := domain.Artist{
		ID:          1,
		Name:        "Artist 1",
		ImageURL:    testURL,
		Description: "This is test artist 1",
		Links: map[string]string{
			"amazonMusic": "https://www.amazon.com/TEST1111",
			"appleMusic":  "https://music.apple.com/artist/1111",
			"spotify":     "https://open.spotify.com/artist/Test1111",
		},
		Parts: []domain.Occupation{testPart1, testPart2},
	}
	artist, err := suite.artistRepository.GetArtist(1)
	if err != nil {
		panic(err)
	}
	assert.Equal(suite.T(), expectedArtist, artist)
}

func (suite *GetArtistTestSuite) TestGetArtistAlias() {
	testURL := "http://www.example.com"
	testPart2 := domain.Occupation{ID: 2, Title: "Part 2"}
	testPart3 := domain.Occupation{ID: 3, Title: "Part 3"}
	expectedArtist := domain.Artist{
		ID:   2,
		Name: "Artist 2",
		Aliases: []domain.Credit{
			{
				Artist: domain.Artist{ID: 3, Name: "Alias Artist 2", ImageURL: testURL},
				Parts:  []domain.Occupation{testPart3},
			},
		},
		ImageURL: testURL,
		Links: map[string]string{
			"spotify": "https://open.spotify.com/artist/Test2222",
		},
		Parts: []domain.Occupation{testPart2},
	}
	artist, err := suite.artistRepository.GetArtist(2)
	if err != nil {
		suite.T().Error(err)
	}
	assert.Equal(suite.T(), expectedArtist, artist)
}

func (suite *GetArtistTestSuite) TestGetArtistGroup() {
	testURL := "http://www.example.com"
	testPart1 := domain.Occupation{ID: 1, Title: "Part 1"}
	testPart2 := domain.Occupation{ID: 2, Title: "Part 2"}
	expectedArtist := domain.Artist{
		ID:   4,
		Name: "Group Artist 1",
		Members: []domain.Credit{
			{
				Artist: domain.Artist{ID: 1, Name: "Artist 1", ImageURL: testURL},
				Parts:  []domain.Occupation{testPart1, testPart2},
			},
			{
				Artist: domain.Artist{ID: 2, Name: "Artist 2", ImageURL: testURL},
				Parts:  []domain.Occupation{testPart2},
			},
		},
		Description: "This is test group artist 1",
		ImageURL:    testURL,
		Parts:       []domain.Occupation{testPart1},
	}
	artist, err := suite.artistRepository.GetArtist(4)
	if err != nil {
		suite.T().Error(err)
	}
	assert.Equal(suite.T(), expectedArtist.ID, artist.ID)
	assert.Equal(suite.T(), expectedArtist.Name, artist.Name)
	assert.ElementsMatch(suite.T(), expectedArtist.Members, artist.Members)
	assert.Equal(suite.T(), expectedArtist.Description, artist.Description)
	assert.Equal(suite.T(), expectedArtist.ImageURL, artist.ImageURL)
}

func (suite *GetArtistTestSuite) TestGetArtistGroupAlias() {
	testURL := "http://www.example.com"
	testPart3 := domain.Occupation{ID: 3, Title: "Part 3"}
	expectedArtist := domain.Artist{
		ID:   5,
		Name: "Group Artist 2",
		Members: []domain.Credit{
			{
				Artist: domain.Artist{ID: 1, Name: "Artist 1", ImageURL: testURL},
				Parts:  []domain.Occupation{},
			},
			{
				Artist: domain.Artist{ID: 2, Name: "Artist 2", ImageURL: testURL},
				Parts:  []domain.Occupation{},
			},
			{
				Artist: domain.Artist{ID: 4, Name: "Group Artist 1", ImageURL: testURL},
				Parts:  []domain.Occupation{},
			},
		},
		Aliases: []domain.Credit{
			{
				Artist: domain.Artist{ID: 6, Name: "Alias Group Artist 2", ImageURL: testURL},
				Parts:  []domain.Occupation{testPart3},
			},
		},
		ImageURL: testURL,
	}
	artist, err := suite.artistRepository.GetArtist(5)
	if err != nil {
		suite.T().Error(err)
	}
	assert.Equal(suite.T(), expectedArtist.ID, artist.ID)
	assert.Equal(suite.T(), expectedArtist.Name, artist.Name)
	assert.ElementsMatch(suite.T(), expectedArtist.Members, artist.Members)
	assert.Equal(suite.T(), expectedArtist.Aliases, artist.Aliases)
	assert.Equal(suite.T(), expectedArtist.ImageURL, artist.ImageURL)
}

func (suite *GetArtistTestSuite) TestGetArtistInvalidID() {
	emptyArtist := domain.Artist{}
	artist, err := suite.artistRepository.GetArtist(100)
	if err != nil {
		suite.T().Error(err)
	}
	assert.Equal(suite.T(), emptyArtist, artist)
}
