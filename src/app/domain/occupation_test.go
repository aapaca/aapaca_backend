package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type OccupationsTestSuite struct {
	suite.Suite
}

func TestOccupationsTestSuite(t *testing.T) {
	suite.Run(t, new(OccupationsTestSuite))
}

func (suite *OccupationsTestSuite) TestAppend() {
	occupations := NewOccupations()
	occupation1 := Occupation{1, "oc1"}
	occupation2 := Occupation{2, "oc2"}
	occupations.Append(occupation1)
	occupations.Append(occupation2)

	expect := []Occupation{occupation1, occupation2}
	actual := occupations.GetOccupationList()
	assert.Equal(suite.T(), expect, actual)
}

func (suite *OccupationsTestSuite) TestContains() {
	occupations := NewOccupations()
	occupation1 := Occupation{1, "oc1"}
	occupation2 := Occupation{2, "oc2"}
	occupations.Append(occupation1)
	occupations.Append(occupation2)

	assert.True(suite.T(), occupations.Contains(1))
	assert.True(suite.T(), occupations.Contains(2))
	assert.False(suite.T(), occupations.Contains(3))
}

func (suite *OccupationsTestSuite) TestIsEmptyWhenTrue() {
	occupations := NewOccupations()

	assert.True(suite.T(), occupations.IsEmpty())
}

func (suite *OccupationsTestSuite) TestIsEmptyWhenFalse() {
	occupations := NewOccupations()
	occupation := Occupation{1, "oc1"}
	occupations.Append(occupation)

	assert.False(suite.T(), occupations.IsEmpty())
}
