package database

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShortenSongLen(t *testing.T) {
	t.Run("less than 10 minutes", func(t *testing.T) {
		in := "00:02:40"
		out := "2:40"
		got := shortenSongLen(in)
		assert.Equal(t, out, got, "Error")
	})
	t.Run("less than 1 minute", func(t *testing.T) {
		in := "00:00:15"
		out := "0:15"
		got := shortenSongLen(in)
		assert.Equal(t, out, got, "Error")
	})
	t.Run("more than 10 minutes", func(t *testing.T) {
		in := "00:12:34"
		out := "12:34"
		got := shortenSongLen(in)
		assert.Equal(t, out, got, "Error")
	})
	t.Run("more than 1 hour", func(t *testing.T) {
		in := "02:09:00"
		out := "2:09:00"
		got := shortenSongLen(in)
		assert.Equal(t, out, got, "Error")
	})
}
