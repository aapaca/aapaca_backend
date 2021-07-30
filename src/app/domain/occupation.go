package domain

type Occupation struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

type Occupations struct {
	Occupations []Occupation
}

func (o *Occupations) Append(oc Occupation) {
	o.Occupations = append(o.Occupations, oc)
}

func (o *Occupations) Contains(id int) bool {
	for _, p := range o.Occupations {
		if p.ID == id {
			return true
		}
	}
	return false
}

func (o *Occupations) IsEmpty() bool {
	return len(o.Occupations) == 0
}
