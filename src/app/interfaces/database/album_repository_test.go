package database

import (
	"domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValues(t *testing.T) {
	t.Run("Convert Empty map to array", func(t *testing.T) {
		credits := []domain.Credit{}
		creditMap := map[int]*domain.Credit{}
		for i := range credits {
			creditMap[i] = &credits[i]
		}

		got := values(creditMap)
		assert.Equal(t, credits, got, "Error")
	})
	t.Run("Convert a map containing 3 credits to array", func(t *testing.T) {
		artist1 := domain.Artist{ID: 1}
		artist2 := domain.Artist{ID: 2}
		artist3 := domain.Artist{ID: 3}
		part := []domain.Occupation{{ID: 1}}
		credits := []domain.Credit{
			{Artist: artist1, Parts: part},
			{Artist: artist2, Parts: part},
			{Artist: artist3, Parts: part},
		}
		creditMap := map[int]*domain.Credit{}
		for i := range credits {
			creditMap[i] = &credits[i]
		}

		got := values(creditMap)
		assert.ElementsMatch(t, credits, got)
	})
}
