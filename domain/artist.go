package domain

type Artist interface {
	ID() int
	Name() string
}