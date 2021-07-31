package domain

import "encoding/json"

type Occupation struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

type Occupations struct {
	occupations []Occupation
}

func NewOccupations() *Occupations {
	return &Occupations{occupations: []Occupation{}}
}

func (o *Occupations) Append(oc Occupation) {
	o.occupations = append(o.occupations, oc)
}

func (o *Occupations) Contains(id int) bool {
	for _, p := range o.occupations {
		if p.ID == id {
			return true
		}
	}
	return false
}

func (o *Occupations) IsEmpty() bool {
	return len(o.occupations) == 0
}

func (o *Occupations) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.occupations)
}
