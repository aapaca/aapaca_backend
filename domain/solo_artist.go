package domain

import "time"

type SoloArtist struct {
	id       int
	name     string
	Part     []Occupation
	Country  string
	Birthday time.Time
}

func (s *SoloArtist) ID() int {
	return s.id
}

func (s *SoloArtist) Name() string {
	return s.name
}
