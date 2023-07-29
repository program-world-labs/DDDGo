package pwsql

import (
	"errors"
	"log"
	"time"

	"github.com/program-world-labs/DDDGo/pkg/pwsql/mysql"
	"github.com/program-world-labs/DDDGo/pkg/pwsql/postgresql"
)

var err = errors.New("unknown sql type")

func InitSQL(sqlType string, dsn string, poolMax int, connAttempts int, connTimeout time.Duration) (ISQLGorm, error) {
	switch sqlType {
	case "mysql":
		return mysql.New(dsn, mysql.MaxPoolSize(poolMax), mysql.ConnAttempts(connAttempts), mysql.ConnTimeout(connTimeout))
	case "postgresql":
		return postgresql.New(dsn, postgresql.MaxPoolSize(poolMax), postgresql.ConnAttempts(connAttempts), postgresql.ConnTimeout(connTimeout))
	default:
		log.Fatalf(err.Error()+": %s", sqlType)
	}

	return nil, err
}
