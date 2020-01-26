package domain

type Song struct {
	ID int
	Name string
	PrimaryArtist Artist
	AttendedArtists []Artist
	Label string // TODO レーベルも構造体にする
	Album Album
	Genre string // TODO ジャンルも構造体にする
}