package domain

import "time"

type Song struct {
	ID            int
	Name          string
	PrimaryArtist interface{}
	Credits       []Credit
	Label         string // TODO レーベルも構造体にする
	Albums        []Album
	Genre         string // TODO ジャンルも構造体にする
	ReleaseDate   *time.Time
	Length        *time.Time // 曲の長さ
	Links         map[string]string
}
