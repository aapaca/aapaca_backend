package domain

import "time"

type Album struct {
	ID            int
	Name          string
	PrimaryArtist interface{}
	Credits       []Credit
	Label         string // レーベルも構造体にすべきかもしれない
	ReleasedDate  time.Time
}
