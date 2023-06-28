package entity

import (
	"fmt"
	"reflect"

	"github.com/program-world-labs/DDDGo/internal/domain"
)

var _ IEntity = (*Base)(nil)

type Base struct {
}

func (d *Base) TableName() string {
	return ""
}

func (d *Base) GetID() string {
	return ""
}

func (d *Base) SetID(_ string) {
}

var ErrEntityNil = fmt.Errorf("entity is nil")

func (d *Base) Transform(entity domain.IEntity) (IEntity, error) {
	if entity == nil {
		return nil, ErrEntityNil
	}

	entityValue := reflect.ValueOf(entity).Elem()
	amountValue := reflect.ValueOf(d).Elem()

	for i := 0; i < entityValue.NumField(); i++ {
		fieldName := entityValue.Type().Field(i).Name
		amountFieldValue := amountValue.FieldByName(fieldName)

		if amountFieldValue.IsValid() && amountFieldValue.CanSet() {
			amountFieldValue.Set(entityValue.Field(i))
		}
	}

	return d, nil
}
