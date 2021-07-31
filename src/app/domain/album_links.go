package domain

import (
	"encoding/json"
	"errors"
)

type AlbumLinks struct {
	links map[string]string
}

func NewAlbumLinks() *AlbumLinks {
	return &AlbumLinks{links: map[string]string{}}
}

func (l *AlbumLinks) Length() int {
	return len(l.links)
}

func (l *AlbumLinks) AddLink(id, serviceName string) error {
	switch serviceName {
	case "amazon_music":
		l.links["amazonMusic"] = "https://www.amazon.com/dp/" + id
	case "apple_music":
		l.links["appleMusic"] = "https://music.apple.com/album/" + id
	case "spotify":
		l.links["spotify"] = "https://open.spotify.com/album/" + id
	default:
		return errors.New("invalid service name")
	}
	return nil
}

func (l *AlbumLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(l.links)
}
