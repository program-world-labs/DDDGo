package relation

import (
	"database/sql"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"

	"github.com/program-world-labs/DDDGo/pkg/pwsql"
)

var _ pwsql.ISQLGorm = (*MockSQL)(nil)

type MockSQL struct {
	db *gorm.DB
}

func NewMock(mockdb *sql.DB) *MockSQL {
	db, err := gorm.Open(
		postgres.New(postgres.Config{
			PreferSimpleProtocol: true,
			Conn:                 mockdb,
		}),
		&gorm.Config{
			PrepareStmt: true,
			Logger: logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{
				Colorful: true,
				LogLevel: logger.Info,
			}),
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true,
			},
			SkipDefaultTransaction: true,
			AllowGlobalUpdate:      false,
		},
	)
	if err != nil {
		panic(err)
	}

	return &MockSQL{db: db}
}

func (m *MockSQL) GetDB() *gorm.DB {
	return m.db
}

func (m *MockSQL) Close() error {
	if m.db != nil {
		sqlDB, err := m.db.DB()
		if err != nil {
			log.Printf("failed to get sql.DB: %v", err)

			return err
		}

		sqlDB.Close()
	}

	return nil
}
