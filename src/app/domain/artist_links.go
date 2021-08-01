package domain

import (
	"encoding/json"
	"errors"
)

type ArtistLinks struct {
	links map[string]string
}

func NewArtistLinks() *ArtistLinks {
	return &ArtistLinks{links: map[string]string{}}
}

func (l *ArtistLinks) Length() int {
	return len(l.links)
}

func (l *ArtistLinks) AddLink(id, serviceName string) error {
	switch serviceName {
	case "amazon_music":
		l.links["amazonMusic"] = "https://www.amazon.com/" + id
	case "apple_music":
		l.links["appleMusic"] = "https://music.apple.com/artist/" + id
	case "spotify":
		l.links["spotify"] = "https://open.spotify.com/artist/" + id
	default:
		return errors.New("invalid service name")
	}
	return nil
}

func (l *ArtistLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(l.links)
}
