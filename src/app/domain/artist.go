package domain

import "time"

type Artist struct {
	ID       int
	Name     string
	Country  string
	Birthday *time.Time
	Members  interface{}
	ImageURL string
}
