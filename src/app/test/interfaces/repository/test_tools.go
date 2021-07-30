package repository

import (
	"domain"
	"interfaces/repository/rdb"
	"io/ioutil"
	"sort"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func InitDb(filepath string, sqlHandler rdb.SqlHandler) error {
	queries, err := readSqlFile(filepath)
	if err != nil {
		return err
	}
	for _, query := range queries {
		_, err := sqlHandler.Execute(query)
		if err != nil {
			return err
		}
	}
	return nil
}

func readSqlFile(filepath string) ([]string, error) {
	b, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	s := string(b)
	queries := strings.Split(s, "\n")
	return queries, nil
}

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
		assert.ElementsMatch(t, expected[i].Parts.Occupations, got[i].Parts.Occupations)
	}
}
