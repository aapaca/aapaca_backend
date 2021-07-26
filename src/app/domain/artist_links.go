package domain

import "errors"

type ArtistLinks struct {
	Links map[string]string
}

func NewArtistLinks() *ArtistLinks {
	return &ArtistLinks{Links: map[string]string{}}
}

func (l *ArtistLinks) Length() int {
	return len(l.Links)
}

func (l *ArtistLinks) AddLink(id, serviceName string) error {
	switch serviceName {
	case "amazon_music":
		l.Links["amazonMusic"] = "https://www.amazon.com/" + id
	case "apple_music":
		l.Links["appleMusic"] = "https://music.apple.com/artist/" + id
	case "spotify":
		l.Links["spotify"] = "https://open.spotify.com/artist/" + id
	default:
		return errors.New("invalid service name")
	}
	return nil
}
