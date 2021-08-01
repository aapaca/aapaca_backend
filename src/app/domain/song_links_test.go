package domain

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type SongLinksTestSuite struct {
	suite.Suite
}

func TestSongLinksTestSuite(t *testing.T) {
	suite.Run(t, new(SongLinksTestSuite))
}

func (suite *SongLinksTestSuite) TestLength() {
	links := NewSongLinks()
	links.AddLink("1", "amazon_music")
	links.AddLink("2", "apple_music")
	links.AddLink("3", "spotify")

	assert.Equal(suite.T(), 3, links.Length())
}

func (suite *SongLinksTestSuite) TestLengthWhenEmpty() {
	links := NewSongLinks()

	assert.Equal(suite.T(), 0, links.Length())
}

func (suite *SongLinksTestSuite) TestAddLinkAmazonMusic() {
	links := NewSongLinks()
	err := links.AddLink("1", "amazon_music")

	expect, _ := json.Marshal(map[string]string{"amazonMusic": "https://www.amazon.com/dp/1"})
	actual, _ := links.MarshalJSON()
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expect, actual)
}

func (suite *SongLinksTestSuite) TestAddLinkAppleMusic() {
	links := NewSongLinks()
	err := links.AddLink("1", "apple_music")

	expect, _ := json.Marshal(map[string]string{"appleMusic": "https://music.apple.com/album/1"})
	actual, _ := links.MarshalJSON()
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expect, actual)
}

func (suite *SongLinksTestSuite) TestAddLinkSpotify() {
	links := NewSongLinks()
	err := links.AddLink("1", "spotify")

	expect, _ := json.Marshal(map[string]string{"spotify": "https://open.spotify.com/track/1"})
	actual, _ := links.MarshalJSON()
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expect, actual)
}

func (suite *SongLinksTestSuite) TestAddLinkInvalidServiceName() {
	links := NewSongLinks()
	err := links.AddLink("1", "hoge")

	expectedError := errors.New("invalid service name")
	assert.Equal(suite.T(), expectedError, err)
}
