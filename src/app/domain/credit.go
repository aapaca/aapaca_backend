package domain

type Credit struct {
	Artist Artist       `json:"artist"`
	Parts  []Occupation `json:"parts"`
}
