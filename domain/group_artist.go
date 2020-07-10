package domain

import "time"

type GroupArtist struct {
	Id       int
	Name     string
	Country  string
	Birthday time.Time
	Members  []interface{}
}
