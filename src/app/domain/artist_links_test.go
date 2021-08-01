package domain

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ArtistLinksTestSuite struct {
	suite.Suite
}

func TestArtistLinksTestSuite(t *testing.T) {
	suite.Run(t, new(ArtistLinksTestSuite))
}

func (suite *ArtistLinksTestSuite) TestLength() {
	links := NewArtistLinks()
	links.AddLink("1", "amazon_music")
	links.AddLink("2", "apple_music")
	links.AddLink("3", "spotify")

	assert.Equal(suite.T(), 3, links.Length())
}

func (suite *ArtistLinksTestSuite) TestLengthWhenEmpty() {
	links := NewArtistLinks()

	assert.Equal(suite.T(), 0, links.Length())
}

func (suite *ArtistLinksTestSuite) TestAddLinkAmazonMusic() {
	links := NewArtistLinks()
	err := links.AddLink("1", "amazon_music")

	expect, _ := json.Marshal(map[string]string{"amazonMusic": "https://www.amazon.com/1"})
	actual, _ := links.MarshalJSON()
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expect, actual)
}

func (suite *ArtistLinksTestSuite) TestAddLinkAppleMusic() {
	links := NewArtistLinks()
	err := links.AddLink("1", "apple_music")

	expect, _ := json.Marshal(map[string]string{"appleMusic": "https://music.apple.com/artist/1"})
	actual, _ := links.MarshalJSON()
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expect, actual)
}

func (suite *ArtistLinksTestSuite) TestAddLinkSpotify() {
	links := NewArtistLinks()
	err := links.AddLink("1", "spotify")

	expect, _ := json.Marshal(map[string]string{"spotify": "https://open.spotify.com/artist/1"})
	actual, _ := links.MarshalJSON()
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expect, actual)
}

func (suite *ArtistLinksTestSuite) TestAddLinkInvalidServiceName() {
	links := NewArtistLinks()
	err := links.AddLink("1", "hoge")

	expectedError := errors.New("invalid service name")
	assert.Equal(suite.T(), expectedError, err)
}
