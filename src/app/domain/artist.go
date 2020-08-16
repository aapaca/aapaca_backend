package domain

import "time"

type Artist struct {
	Id       int
	Name     string
	Country  string
	Birthday time.Time
	Members  []interface{}
}
