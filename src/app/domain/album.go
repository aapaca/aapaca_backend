package domain

import "time"

type Album struct {
	ID            int         `json:"id"`
	Name          string      `json:"name"`
	PrimaryArtist interface{} `json:"primaryArtist,omitempty"`
	Credits       []Credit    `json:"credits,omitempty"`
	Label         string      `json:"label,omitempty"` // レーベルも構造体にすべきかもしれない
	ReleasedDate  *time.Time  `json:"releasedDate,omitempty"`
	ImageURL      string      `json:"imageUrl"`
	Description   string      `json:"description,omitempty"`
	Links         map[string]string `json:"links,omitempty"`
}
