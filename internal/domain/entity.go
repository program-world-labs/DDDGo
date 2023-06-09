package domain

type IEntity interface {
	GetID() string
	SetID(string) error
	TableName() string
}
