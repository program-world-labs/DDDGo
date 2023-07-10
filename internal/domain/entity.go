package domain

type IEntity interface {
	GetID() string
	SetID(string)
}
