package domain

type IEntity interface {
	GetID() string
	SetID(string)
}

type List struct {
	Limit  int64     `json:"limit"`
	Offset int64     `json:"offset"`
	Total  int64     `json:"total"`
	Data   []IEntity `json:"data"`
}
