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
	ReleasedDate  *time.Time
	Length        *time.Time // 曲の長さ
	ImageURL      string
	Links         map[string]string
}
