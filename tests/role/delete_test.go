package role_test

import (
	"context"
	"testing"

	"github.com/cucumber/godog"
	gomock "github.com/golang/mock/gomock"

	application_role "github.com/program-world-labs/DDDGo/internal/application/role"
	mock_repo "github.com/program-world-labs/DDDGo/tests/mocks"
	mocks "github.com/program-world-labs/DDDGo/tests/mocks/role"
	mocks_user "github.com/program-world-labs/DDDGo/tests/mocks/user"
)

type ServiceDeleteTest struct {
	t        *testing.T
	mockCtrl *gomock.Controller

	input         *application_role.DeletedInput
	expect        *application_role.Output
	repoMock      *mocks.MockRoleRepository
	userRepoMock  *mocks_user.MockUserRepository
	transRepoMock *mock_repo.MockITransactionRepo
	service       *application_role.ServiceImpl
}

func (st *ServiceDeleteTest) reset() {
	// logger := pwlogger.NewDevelopmentLogger("")

	st.input = nil
	st.expect = nil
	st.repoMock = mocks.NewMockRoleRepository(st.mockCtrl)
	st.userRepoMock = mocks_user.NewMockUserRepository(st.mockCtrl)
	st.transRepoMock = mock_repo.NewMockITransactionRepo(st.mockCtrl)
	// st.service = application_role.NewServiceImpl(st.repoMock, st.transRepoMock, logger)
}

func (st *ServiceDeleteTest) givenData(id string) error {
	st.input = &application_role.DeletedInput{
		ID: id,
	}
	st.expect = &application_role.Output{}

	return nil
}

func (st *ServiceDeleteTest) whenDeleteRole(_ context.Context) error {
	// e := st.input.ToEntity()
	// st.repoMock.EXPECT().Delete(gomock.Any(), RoleEquals(e)).Return(nil)

	return nil
}

func (st *ServiceDeleteTest) whenDeleteNotExistingRole(_ context.Context) error {
	// e := newRolseExistError()
	// st.repoMock.EXPECT().Create(gomock.Any(), RoleEquals(st.input.ToEntity())).Return(nil, e)

	return nil
}

func (st *ServiceDeleteTest) thenSuccess(name, description, permission string) error {
	// st.expect = &application_role.Output{
	// 	Name:        name,
	// 	Description: description,
	// 	Permissions: strings.Split(permission, ","),
	// 	CreatedAt:   st.expect.CreatedAt,
	// 	UpdatedAt:   st.expect.UpdatedAt,
	// }
	// actual, err := st.service.CreateRole(context.Background(), st.input)

	// if err != nil {
	// 	return err
	// }

	// err = tests.AssertExpectedAndActual(assert.Equal, st.expect, actual, "Expected %d role created success and equal, but there is %d", st.expect, actual)

	// return err
	return nil
}

func (st *ServiceDeleteTest) whenDeleteHasUserRole(_ context.Context) error {
	// e := newRolseExistError()
	// st.repoMock.EXPECT().Create(gomock.Any(), RoleEquals(st.input.ToEntity())).Return(nil, e)

	return nil
}

func (st *ServiceDeleteTest) InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
		st.mockCtrl = gomock.NewController(st.t)
		st.reset()

		return ctx, nil
	})

	ctx.After(func(ctx context.Context, sc *godog.Scenario, err error) (context.Context, error) {
		st.mockCtrl.Finish()

		return ctx, nil
	})

	ctx.Step(`^提供 (.*?)$`, st.givenData)
	ctx.Step(`^ID存在並嘗試刪除角色$`, st.whenDeleteRole)
	ctx.Step(`^ID不存在並嘗試刪除角色$`, st.whenDeleteNotExistingRole)
	ctx.Step(`^ID存在並且角色ID已經被分配給用戶$`, st.whenDeleteHasUserRole)
	ctx.Step(`^角色成功被刪除$`, st.thenSuccess)
	// ctx.Step(`^返回一個錯誤，說明角色ID不存在$`, st.whenCreateExistingRole)
	// ctx.Step(`^返回一個錯誤，說明角色已經被分配給用戶$`, st.whenInvalidFormat)
}

// func TestDelete(t *testing.T) {
// 	t.Parallel()

// 	serviceTest := &ServiceDeleteTest{
// 		t: t,
// 	}

// 	// Create the report directory
// 	reportPath := filepath.Join("..", "report", "TestRoleDeleteUsecase.json")
// 	// Create the directory if it does not exist
// 	err := os.MkdirAll(filepath.Dir(reportPath), os.ModePerm)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	// Run the test suite
// 	suite := godog.TestSuite{
// 		Name:                "Delete",
// 		ScenarioInitializer: serviceTest.InitializeScenario,
// 		Options: &godog.Options{
// 			Format:   "cucumber:" + reportPath,
// 			Paths:    []string{"features/usecase/role_deleted.feature"},
// 			TestingT: t,
// 			Output:   colors.Colored(os.Stdout),
// 		},
// 	}

// 	if suite.Run() != 0 {
// 		t.Log("non-zero status returned, failed to run feature tests")
// 	}
// }
