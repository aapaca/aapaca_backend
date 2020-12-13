package domain

type Song struct {
	ID            int               `json:"id"`
	Name          string            `json:"name"`
	PrimaryArtist interface{}       `json:"primaryArtist,omitempty"`
	Credits       []Credit          `json:"credits,omitempty"`
	SongLen       string            `json:"length,omitempty"`
	Order         string            `json:"order,omitempty"`
	Label         string            `json:"label,omitempty"` // TODO レーベルも構造体にする
	Album         interface{}       `json:"album,omitempty"`
	Genre         string            `json:"genre,omitempty"` // TODO ジャンルも構造体にする
	Links         map[string]string `json:"links,omitempty"`
}
