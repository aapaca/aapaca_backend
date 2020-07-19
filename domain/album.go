package domain

import "time"

type Album struct {
	ID              int
	Name            string
	PrimaryArtist   interface{}
	AttendedArtists []interface{}
	Label           string // レーベルも構造体にすべきかもしれない
	ReleasedDate    time.Time
}