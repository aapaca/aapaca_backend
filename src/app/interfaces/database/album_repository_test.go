package database

import (
	"domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValues(t *testing.T) {
	t.Run("Empty map", func(t *testing.T){
		credits := []domain.Credit{}
		creditMap := map[int]*domain.Credit{}
		for i := range credits {
			creditMap[i] = &credits[i]
		}

		ret := values(creditMap)
		assert.Equal(t, credits, ret, "Error")
	})
	t.Run("2 credits", func(t *testing.T){
		artist1 := domain.Artist{ID: 1}
		artist2 := domain.Artist{ID: 2}
		part := []domain.Occupation{{ID: 1}}
		credits := []domain.Credit{
			{Artist: artist1, Parts: part},
			{Artist: artist2, Parts: part},
		}

		creditMap := map[int]*domain.Credit{}
		for i := range credits {
			creditMap[i] = &credits[i]
		}

		ret := values(creditMap)
		assert.Equal(t, credits, ret, "Error")
	})
}
