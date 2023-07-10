package pwsql

import "gorm.io/gorm"

type ISQLGorm interface {
	GetDB() *gorm.DB
	Close() error
}
