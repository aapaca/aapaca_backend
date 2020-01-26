package domain

import "time"

type Album struct {
	ID int
	Name string
	PrimaryArtist Artist
	AttendedArtist []Artist
	Label string // レーベルも構造体にすべきかもしれない
	ReleasedDate time.Time
}