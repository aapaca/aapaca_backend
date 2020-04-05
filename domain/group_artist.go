package domain

import "time"

type GroupArtist struct {
	id       int
	name     string
	Country  string
	Birthday time.Time
	Members []interface{}
}
