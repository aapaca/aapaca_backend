package domain

import "time"

type Artist struct {
	ID          int               `json:"id"`
	Name        string            `json:"name"`
	Country     string            `json:"country,omitempty"`
	Birthday    *time.Time        `json:"birthday,omitempty"`
	Members     interface{}       `json:"members,omitempty"`
	Aliases     interface{}       `json:"aliases,omitempty"`
	ImageURL    string            `json:"imageUrl,omitempty"`
	Description string            `json:"description,omitempty"`
	Links       map[string]string `json:"links,omitempty"`
}
