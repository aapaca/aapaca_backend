package domain

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type AlbumLinksTestSuite struct {
	suite.Suite
}

func TestAlbumLinksTestSuite(t *testing.T) {
	suite.Run(t, new(AlbumLinksTestSuite))
}

func (suite *AlbumLinksTestSuite) TestLength() {
	links := NewAlbumLinks()
	links.AddLink("1", "amazon_music")
	links.AddLink("2", "apple_music")
	links.AddLink("3", "spotify")

	assert.Equal(suite.T(), 3, links.Length())
}

func (suite *AlbumLinksTestSuite) TestLengthWhenEmpty() {
	links := NewAlbumLinks()

	assert.Equal(suite.T(), 0, links.Length())
}

func (suite *AlbumLinksTestSuite) TestAddLinkAmazonMusic() {
	links := NewAlbumLinks()
	err := links.AddLink("1", "amazon_music")

	expect, _ := json.Marshal(map[string]string{"amazonMusic": "https://www.amazon.com/dp/1"})
	actual, _ := links.MarshalJSON()
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expect, actual)
}

func (suite *AlbumLinksTestSuite) TestAddLinkAppleMusic() {
	links := NewAlbumLinks()
	err := links.AddLink("1", "apple_music")

	expect, _ := json.Marshal(map[string]string{"appleMusic": "https://music.apple.com/album/1"})
	actual, _ := links.MarshalJSON()
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expect, actual)
}

func (suite *AlbumLinksTestSuite) TestAddLinkSpotify() {
	links := NewAlbumLinks()
	err := links.AddLink("1", "spotify")

	expect, _ := json.Marshal(map[string]string{"spotify": "https://open.spotify.com/album/1"})
	actual, _ := links.MarshalJSON()
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expect, actual)
}

func (suite *AlbumLinksTestSuite) TestAddLinkInvalidServiceName() {
	links := NewAlbumLinks()
	err := links.AddLink("1", "hoge")

	expectedError := errors.New("invalid service name")
	assert.Equal(suite.T(), expectedError, err)
}
