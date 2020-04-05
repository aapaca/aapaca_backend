package domain

type Song struct {
	ID int
	Name string
	PrimaryArtist interface{}
	AttendedArtists []interface{}
	Label string // TODO レーベルも構造体にする
	Album Album
	Genre string // TODO ジャンルも構造体にする
}