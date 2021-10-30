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

const (
	TestURL = "http://www.example.com"
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

	err := repository.InitDb("testdata/get_artist_init.sql", suite.sqlHandler)
	if err != nil {
		suite.T().Error(err)
	}
	suite.artistRepository = ArtistRepository{
		SqlHandler: suite.sqlHandler,
	}
}

func (suite *GetArtistTestSuite) TearDownSuite() {
	err := infrastructure.DeleteAllRecords(suite.sqlHandler)
	if err != nil {
		suite.T().Error(err)
	}
}

func (suite *GetArtistTestSuite) TestGetArtist() {
	testPart1 := domain.Occupation{ID: 1, Title: "Part 1"}
	testPart2 := domain.Occupation{ID: 2, Title: "Part 2"}
	testParts1 := domain.NewOccupations()
	testParts1.Append(testPart1)
	testParts1.Append(testPart2)
	links := domain.NewArtistLinks()
	links.AddLink("TEST1111", "amazon_music")
	links.AddLink("1111", "apple_music")
	links.AddLink("Test1111", "spotify")
	expectedArtist := domain.Artist{
		ID:          1,
		Name:        "Artist 1",
		ImageURL:    TestURL,
		Description: "This is test artist 1",
		Links:       links,
		Parts:       testParts1,
	}

	artist, err := suite.artistRepository.GetArtist(1)

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expectedArtist.ID, artist.ID)
	assert.Equal(suite.T(), expectedArtist.Name, artist.Name)
	assert.Equal(suite.T(), expectedArtist.ImageURL, artist.ImageURL)
	assert.Equal(suite.T(), expectedArtist.Description, artist.Description)
	assert.Equal(suite.T(), expectedArtist.Links, artist.Links)
	repository.AssertParts(suite.T(), expectedArtist.Parts, artist.Parts)
}

func (suite *GetArtistTestSuite) TestGetArtistAlias() {
	testPart2 := domain.Occupation{ID: 2, Title: "Part 2"}
	testParts2 := domain.NewOccupations()
	testParts2.Append(testPart2)
	links := domain.NewArtistLinks()
	links.AddLink("Test2222", "spotify")
	expectedArtist := domain.Artist{
		ID:       2,
		Name:     "Artist 2",
		ImageURL: TestURL,
		Links:    links,
		Parts:    testParts2,
	}

	artist, err := suite.artistRepository.GetArtist(2)

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expectedArtist, artist)
}

func (suite *GetArtistTestSuite) TestGetArtistGroup() {
	testPart1 := domain.Occupation{ID: 1, Title: "Part 1"}
	testPart2 := domain.Occupation{ID: 2, Title: "Part 2"}
	testParts4 := domain.NewOccupations()
	testParts4.Append(testPart1)
	testParts4.Append(testPart2)
	expectedArtist := domain.Artist{
		ID:          4,
		Name:        "Group Artist 1",
		Description: "This is test group artist 1",
		ImageURL:    TestURL,
		Parts:       testParts4,
	}

	artist, err := suite.artistRepository.GetArtist(4)

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expectedArtist.ID, artist.ID)
	assert.Equal(suite.T(), expectedArtist.Name, artist.Name)
	assert.Equal(suite.T(), expectedArtist.Description, artist.Description)
	assert.Equal(suite.T(), expectedArtist.ImageURL, artist.ImageURL)
	repository.AssertParts(suite.T(), expectedArtist.Parts, artist.Parts)
}

func (suite *GetArtistTestSuite) TestGetArtistGroupAlias() {
	expectedArtist := domain.Artist{
		ID:       5,
		Name:     "Group Artist 2",
		ImageURL: TestURL,
	}

	artist, err := suite.artistRepository.GetArtist(5)

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expectedArtist.ID, artist.ID)
	assert.Equal(suite.T(), expectedArtist.Name, artist.Name)
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

type FindMembersTestSuite struct {
	suite.Suite
	sqlHandler       rdb.SqlHandler
	artistRepository ArtistRepository
}

func TestFindMembersTestSuite(t *testing.T) {
	suite.Run(t, new(FindMembersTestSuite))
}

func (suite *FindMembersTestSuite) SetupSuite() {
	suite.sqlHandler = infrastructure.NewSqlHandler()

	err := repository.InitDb("testdata/find_members_init.sql", suite.sqlHandler)
	if err != nil {
		suite.T().Error(err)
	}
	suite.artistRepository = ArtistRepository{
		SqlHandler: suite.sqlHandler,
	}
}

func (suite *FindMembersTestSuite) TearDownSuite() {
	err := infrastructure.DeleteAllRecords(suite.sqlHandler)
	if err != nil {
		suite.T().Error(err)
	}
}

func (suite *FindMembersTestSuite) TestFindMembers() {
	testPart1 := domain.Occupation{ID: 1, Title: "Part 1"}
	testPart2 := domain.Occupation{ID: 2, Title: "Part 2"}
	testParts1 := domain.NewOccupations()
	testParts2 := domain.NewOccupations()
	testParts1.Append(testPart1)
	testParts1.Append(testPart2)
	testParts2.Append(testPart2)
	expect := []domain.Credit{
		{
			Artist: &domain.Artist{ID: 1, Name: "Artist 1", ImageURL: TestURL},
			Parts:  testParts1,
		},
		{
			Artist: &domain.Artist{ID: 2, Name: "Artist 2", ImageURL: TestURL},
			Parts:  testParts2,
		},
	}

	actual, err := suite.artistRepository.FindMembers(4)

	assert.NoError(suite.T(), err)
	repository.AssertCredits(suite.T(), expect, actual)
}

func (suite *FindMembersTestSuite) TestNoParts() {
	expect := []domain.Credit{
		{
			Artist: &domain.Artist{ID: 1, Name: "Artist 1", ImageURL: TestURL},
			Parts:  domain.NewOccupations(),
		},
		{
			Artist: &domain.Artist{ID: 2, Name: "Artist 2", ImageURL: TestURL},
			Parts:  domain.NewOccupations(),
		},
		{
			Artist: &domain.Artist{ID: 4, Name: "Group Artist 1", ImageURL: TestURL},
			Parts:  domain.NewOccupations(),
		},
	}

	actual, err := suite.artistRepository.FindMembers(5)

	assert.NoError(suite.T(), err)
	repository.AssertCredits(suite.T(), expect, actual)
}

func (suite *FindMembersTestSuite) TestNoMembers() {
	actual, err := suite.artistRepository.FindMembers(100)

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), 0, len(actual))
}

type FindAliasesTestSuite struct {
	suite.Suite
	sqlHandler       rdb.SqlHandler
	artistRepository ArtistRepository
}

func TestFindAliasesTestSuite(t *testing.T) {
	suite.Run(t, new(FindAliasesTestSuite))
}

func (suite *FindAliasesTestSuite) SetupSuite() {
	suite.sqlHandler = infrastructure.NewSqlHandler()

	err := repository.InitDb("testdata/find_aliases_init.sql", suite.sqlHandler)
	if err != nil {
		suite.T().Error(err)
	}
	suite.artistRepository = ArtistRepository{
		SqlHandler: suite.sqlHandler,
	}
}

func (suite *FindAliasesTestSuite) TearDownSuite() {
	err := infrastructure.DeleteAllRecords(suite.sqlHandler)
	if err != nil {
		suite.T().Error(err)
	}
}

func (suite *FindAliasesTestSuite) TestFindAliases() {
	testPart3 := domain.Occupation{ID: 3, Title: "Part 3"}
	testParts3 := domain.NewOccupations()
	testParts3.Append(testPart3)
	expect := []domain.Credit{
		{
			Artist: &domain.Artist{ID: 3, Name: "Alias Artist 2", ImageURL: TestURL},
			Parts:  testParts3,
		},
	}

	actual, err := suite.artistRepository.FindAliases(2)

	assert.NoError(suite.T(), err)
	repository.AssertCredits(suite.T(), expect, actual)
}

func (suite *FindAliasesTestSuite) TestNoAliases() {
	actual, err := suite.artistRepository.FindAliases(100)

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), 0, len(actual))
}
