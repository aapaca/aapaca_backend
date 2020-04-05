package domain

import "time"

type GroupArtist struct {
	id       int
	name     string
	Country  string
	Birthday time.Time
	Members []interface{}
}

func (g *GroupArtist) ID() int {
	return g.id
}

func (g *GroupArtist) Name() string {
	return g.name
}
