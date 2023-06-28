package entity

import "github.com/program-world-labs/DDDGo/internal/domain"

type IEntity interface {
	domain.IEntity
	TableName() string
	Transform(domain.IEntity) (IEntity, error)
}
