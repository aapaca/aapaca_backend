package repository

import (
	"domain"
	"interfaces/repository/rdb"
	"test/infrastructure"
	"test/interfaces/repository"
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
	err := repository.InitDb("testdata/get_artist_init.sql", suite.sqlHandler)
	if err != nil {
		suite.T().Error(err)
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
	testParts1 := domain.Occupations{[]domain.Occupation{testPart1, testPart2}}
	links := domain.NewArtistLinks()
	links.AddLink("TEST1111", "amazon_music")
	links.AddLink("1111", "apple_music")
	links.AddLink("Test1111", "spotify")
	expectedArtist := domain.Artist{
		ID:          1,
		Name:        "Artist 1",
		ImageURL:    testURL,
		Description: "This is test artist 1",
		Links:       links,
		Parts:       &testParts1,
	}
	artist, err := suite.artistRepository.GetArtist(1)
	if err != nil {
		panic(err)
	}
	assert.Equal(suite.T(), expectedArtist.ID, artist.ID)
	assert.Equal(suite.T(), expectedArtist.Name, artist.Name)
	assert.Equal(suite.T(), expectedArtist.ImageURL, artist.ImageURL)
	assert.Equal(suite.T(), expectedArtist.Description, artist.Description)
	assert.Equal(suite.T(), expectedArtist.Links, artist.Links)
	assert.ElementsMatch(suite.T(), expectedArtist.Parts.Occupations, artist.Parts.Occupations)
}

func (suite *GetArtistTestSuite) TestGetArtistAlias() {
	testURL := "http://www.example.com"
	testPart2 := domain.Occupation{ID: 2, Title: "Part 2"}
	testPart3 := domain.Occupation{ID: 3, Title: "Part 3"}
	testParts2 := domain.Occupations{[]domain.Occupation{testPart2}}
	testParts3 := domain.Occupations{[]domain.Occupation{testPart3}}
	links := domain.NewArtistLinks()
	links.AddLink("Test2222", "spotify")
	expectedArtist := domain.Artist{
		ID:   2,
		Name: "Artist 2",
		Aliases: []domain.Credit{
			{
				Artist: &domain.Artist{ID: 3, Name: "Alias Artist 2", ImageURL: testURL},
				Parts:  &testParts3,
			},
		},
		ImageURL: testURL,
		Links:    links,
		Parts:    &testParts2,
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
	testParts1 := domain.Occupations{[]domain.Occupation{testPart1, testPart2}}
	testParts2 := domain.Occupations{[]domain.Occupation{testPart2}}
	testParts4 := domain.Occupations{[]domain.Occupation{testPart1}}
	expectedArtist := domain.Artist{
		ID:   4,
		Name: "Group Artist 1",
		Members: []domain.Credit{
			{
				Artist: &domain.Artist{ID: 1, Name: "Artist 1", ImageURL: testURL},
				Parts:  &testParts1,
			},
			{
				Artist: &domain.Artist{ID: 2, Name: "Artist 2", ImageURL: testURL},
				Parts:  &testParts2,
			},
		},
		Description: "This is test group artist 1",
		ImageURL:    testURL,
		Parts:       &testParts4,
	}
	artist, err := suite.artistRepository.GetArtist(4)
	if err != nil {
		suite.T().Error(err)
	}
	assert.Equal(suite.T(), expectedArtist.ID, artist.ID)
	assert.Equal(suite.T(), expectedArtist.Name, artist.Name)
	repository.AssertCredits(suite.T(), expectedArtist.Members.([]domain.Credit), artist.Members.([]domain.Credit))
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
				Artist: &domain.Artist{ID: 1, Name: "Artist 1", ImageURL: testURL},
				Parts:  &domain.Occupations{},
			},
			{
				Artist: &domain.Artist{ID: 2, Name: "Artist 2", ImageURL: testURL},
				Parts:  &domain.Occupations{},
			},
			{
				Artist: &domain.Artist{ID: 4, Name: "Group Artist 1", ImageURL: testURL},
				Parts:  &domain.Occupations{},
			},
		},
		Aliases: []domain.Credit{
			{
				Artist: &domain.Artist{ID: 6, Name: "Alias Group Artist 2", ImageURL: testURL},
				Parts:  &domain.Occupations{[]domain.Occupation{testPart3}},
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
