package role_test

import (
	"context"
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/lib/pq"
	"github.com/stretchr/testify/require"

	datasourceSQL "github.com/program-world-labs/DDDGo/internal/infra/datasource/sql"
	"github.com/program-world-labs/DDDGo/internal/infra/dto"
	"github.com/program-world-labs/DDDGo/pkg/pwsql/relation"
)

type test struct {
	name string
	mock func()
	res  interface{}
	err  error
}

func crudSQL(t *testing.T) (*datasourceSQL.CRUDDatasourceImpl, *sql.DB, sqlmock.Sqlmock) {
	t.Helper()

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	s := relation.NewMock(db)
	ds := datasourceSQL.NewCRUDDatasourceImpl(s)

	return ds, db, mock
}

func TestCreateRole(t *testing.T) {
	t.Parallel()
	datasource, _, mock := crudSQL(t)

	tests := []test{
		{
			name: "create role",
			mock: func() {
				mock.ExpectPrepare("^INSERT INTO \"Roles\".*")
				mock.ExpectExec("^INSERT INTO \"Roles\".*").
					WithArgs(sqlmock.AnyArg(), "test", "test", pq.Array([]string{"test", "test"}), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
			res: &dto.Role{
				Name:        "test",
				Description: "test",
				Permissions: []string{"test", "test"},
			},
			err: nil,
		},
	}

	for _, tc := range tests {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			tc.mock()

			// now we execute our method
			role := &dto.Role{
				Name:        "test",
				Description: "test",
				Permissions: []string{"test", "test"},
			}
			value, err := datasource.Create(context.Background(), role)
			v, ok := value.(*dto.Role)
			require.True(t, ok)
			r, ok := tc.res.(*dto.Role)
			require.True(t, ok)

			require.Equal(t, v.Name, r.Name)
			require.Equal(t, v.Description, r.Description)
			require.Equal(t, v.Permissions, r.Permissions)
			require.ErrorIs(t, err, tc.err)
			// we make sure that all expectations were met
			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}
