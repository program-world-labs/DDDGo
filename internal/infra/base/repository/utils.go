package repository

import (
	"fmt"
	"reflect"

	"github.com/program-world-labs/DDDGo/internal/domain"
	"github.com/program-world-labs/DDDGo/internal/infra/base/entity"
)

var ErrEntityNil = fmt.Errorf("entity is nil")

func Transform(e domain.IEntity, dtoEntity entity.IEntity) (entity.IEntity, error) {
	if e == nil {
		return nil, ErrEntityNil
	}

	entityValue := reflect.ValueOf(e).Elem()
	dtoValue := reflect.ValueOf(dtoEntity).Elem()

	for i := 0; i < entityValue.NumField(); i++ {
		fieldName := entityValue.Type().Field(i).Name
		dtoFieldValue := dtoValue.FieldByName(fieldName)

		if dtoFieldValue.IsValid() && dtoFieldValue.CanSet() {
			dtoFieldValue.Set(entityValue.Field(i))
		}
	}

	return dtoEntity, nil
}
