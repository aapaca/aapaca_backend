package domain

import "errors"

type AlbumLinks struct {
	Links map[string]string
}

func NewAlbumLinks() *AlbumLinks {
	return &AlbumLinks{Links: map[string]string{}}
}

func (l *AlbumLinks) Length() int {
	return len(l.Links)
}

func (l *AlbumLinks) AddLink(id, serviceName string) error {
	switch serviceName {
	case "amazon_music":
		l.Links["amazonMusic"] = "https://www.amazon.com/dp/" + id
	case "apple_music":
		l.Links["appleMusic"] = "https://music.apple.com/album/" + id
	case "spotify":
		l.Links["spotify"] = "https://open.spotify.com/album/" + id
	default:
		return errors.New("invalid service name")
	}
	return nil
}
