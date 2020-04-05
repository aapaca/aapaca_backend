package domain

import "time"

type SoloArtist struct {
	Id       int
	Name     string
	Part     []Occupation
	Country  string
	Birthday time.Time
}
