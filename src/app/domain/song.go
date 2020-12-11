package domain

import "time"

type Song struct {
	ID            int               `json:"id"`
	Name          string            `json:"name"`
	PrimaryArtist interface{}       `json:"primaryArtist,omitempty"`
	Credits       []Credit          `json:"credits,omitempty"`
	SongLen       string            `json:"length,omitempty"`
	Order         string            `json:"order,omitempty"`
	Label         string            `json:"label,omitempty"` // TODO レーベルも構造体にする
	Albums        []Album           `json:"albums,omitempty"`
	Genre         string            `json:"genre,omitempty"` // TODO ジャンルも構造体にする
	ReleasedDate  *time.Time        `json:"releasedDate,omitempty"`
	ImageURL      string            `json:"imageUrl,omitempty"`
	Links         map[string]string `json:"links,omitempty"`
}
