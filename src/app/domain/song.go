package domain

type Song struct {
	ID            int
	Name          string
	PrimaryArtist interface{}
	Credits       []Credit
	Label         string // TODO レーベルも構造体にする
	Albums        []Album
	Genre         string // TODO ジャンルも構造体にする
}
