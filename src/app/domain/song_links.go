package domain

import (
	"encoding/json"
	"errors"
)

type SongLinks struct {
	links map[string]string
}

func NewSongLinks() *SongLinks {
	return &SongLinks{links: map[string]string{}}
}

func (l *SongLinks) Length() int {
	return len(l.links)
}

func (l *SongLinks) AddLink(id, serviceName string) error {
	switch serviceName {
	case "amazon_music":
		l.links["amazonMusic"] = "https://www.amazon.com/dp/" + id
	case "apple_music":
		l.links["appleMusic"] = "https://music.apple.com/album/" + id
	case "spotify":
		l.links["spotify"] = "https://open.spotify.com/track/" + id
	default:
		return errors.New("invalid service name")
	}
	return nil
}

func (l *SongLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(l.links)
}
