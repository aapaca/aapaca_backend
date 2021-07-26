package domain

import "errors"

type SongLinks struct {
	Links map[string]string
}

func NewSongLinks() *SongLinks {
	return &SongLinks{Links: map[string]string{}}
}

func (l *SongLinks) Length() int {
	return len(l.Links)
}

func (l *SongLinks) AddLink(id, serviceName string) error {
	switch serviceName {
	case "amazon_music":
		l.Links["amazonMusic"] = "https://www.amazon.com/dp/" + id
	case "apple_music":
		l.Links["appleMusic"] = "https://music.apple.com/album/" + id
	case "spotify":
		l.Links["spotify"] = "https://open.spotify.com/track/" + id
	default:
		return errors.New("invalid service name")
	}
	return nil
}
