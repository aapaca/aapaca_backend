package domain

import "time"

type SoloArtist struct {
	Id       int
	Name     string
	Parts    []Occupation
	Country  string
	Birthday time.Time
}
