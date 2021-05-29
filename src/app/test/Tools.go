package test

import (
	"domain"
	"net/http"
	"net/http/httptest"
	"sort"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func AssertCredits(t *testing.T, expected []domain.Credit, got []domain.Credit) {
	assert.Equal(t, len(expected), len(got))
	// sort expected and got by ArtistID
	sort.Slice(expected, func(i, j int) bool {
		p, q := expected[i], expected[j]
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

func CreateContextInstance(path string, paramName string, paramValue string) (*httptest.ResponseRecorder, echo.Context) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath(path)
	c.SetParamNames(paramName)
	c.SetParamValues(paramValue)
	return rec, c
}
