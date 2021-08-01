package domain

type Credit struct {
	Artist *Artist      `json:"artist"`
	Parts  *Occupations `json:"parts"`
}
