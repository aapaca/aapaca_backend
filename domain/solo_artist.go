package domain

import "time"

type SoloArtist struct {
	id       int
	name     string
	Part     []Occupation
	Country  string
	Birthday time.Time
}
