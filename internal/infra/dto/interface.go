package dto

import "github.com/program-world-labs/DDDGo/internal/domain"

type IRepoEntity interface {
	domain.IEntity
	TableName() string
	Transform(domain.IEntity) (IRepoEntity, error)
	BackToDomain() (domain.IEntity, error)
	ToJSON() (string, error)
	DecodeJSON(string) error
}
