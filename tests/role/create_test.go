package role_test

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/cucumber/godog"
	"github.com/cucumber/godog/colors"
	gomock "github.com/golang/mock/gomock"
	"github.com/program-world-labs/pwlogger"
	"github.com/stretchr/testify/assert"

	application_role "github.com/program-world-labs/DDDGo/internal/application/role"
	"github.com/program-world-labs/DDDGo/internal/domain/domainerrors"
	"github.com/program-world-labs/DDDGo/internal/domain/entity"
	infra_sql "github.com/program-world-labs/DDDGo/internal/infra/datasource/sql"
	infra_repo "github.com/program-world-labs/DDDGo/internal/infra/repository"
	"github.com/program-world-labs/DDDGo/tests"
	mock_repo "github.com/program-world-labs/DDDGo/tests/mocks"
	mocks "github.com/program-world-labs/DDDGo/tests/mocks/role"
	mocks_user "github.com/program-world-labs/DDDGo/tests/mocks/user"
)

type roleMatcher struct {
	expected *entity.Role
}

func (m *roleMatcher) Matches(x interface{}) bool {
	role, ok := x.(*entity.Role)
	if !ok {
		return false
	}

	// 將 CreatedAt 和 UpdatedAt 設置為固定的值或者 nil
	roleCopy := *role
	roleCopy.CreatedAt = time.Time{}
	roleCopy.UpdatedAt = time.Time{}

	expectedCopy := *m.expected
	expectedCopy.CreatedAt = time.Time{}
	expectedCopy.UpdatedAt = time.Time{}

	return reflect.DeepEqual(roleCopy, expectedCopy)
}

func (m *roleMatcher) String() string {
	return fmt.Sprintf("is equal to %v", m.expected)
}

func RoleEquals(expected *entity.Role) gomock.Matcher {
	return &roleMatcher{expected: expected}
}

type ServiceTest struct {
	t        *testing.T
	mockCtrl *gomock.Controller

	input         *application_role.CreatedInput
	expect        *application_role.Output
	repoMock      *mocks.MockRoleRepository
	userRepoMock  *mocks_user.MockUserRepository
	transRepoMock *mock_repo.MockITransactionRepo
	service       *application_role.ServiceImpl
	eventStoreDB  *mock_repo.MockEventStore
	producer      *mock_repo.MockProducer
}

func (st *ServiceTest) InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
		st.mockCtrl = gomock.NewController(st.t)
		st.reset()

		return ctx, nil
	})

	ctx.After(func(ctx context.Context, sc *godog.Scenario, err error) (context.Context, error) {
		st.mockCtrl.Finish()

		return ctx, nil
	})

	ctx.Step(`^提供 (.*?), (.*?), (.*?)$`, st.givenData)
	ctx.Step(`^創建一個新角色$`, st.whenCreateNewRole)
	ctx.Step(`^角色(.*?), (.*?), (.*?)，成功被創建$`, st.thenSuccess)
	ctx.Step(`^嘗試創建一個已存在的角色名稱$`, st.whenCreateExistingRole)
	ctx.Step(`^帶入有問題的輸入$`, st.whenInvalidFormat)
	ctx.Step(`^應該返回一個錯誤，說明角色名稱已存在$`, st.thenErrorRoleAlreadyExists)
	ctx.Step(`^應該返回一個錯誤，說明權限格式不正確$`, st.thenErrorInvalidPermissionFormat)
	ctx.Step(`^應該返回一個錯誤，說明名稱格式不正確$`, st.thenErrorInvalidNameFormat)
	ctx.Step(`^應該返回一個錯誤，說明角色名稱長度超過最大值$`, st.thenErrorRoleNameExceedsMaxLength)
	ctx.Step(`^應該返回一個錯誤，說明角色描述長度超過最大值$`, st.thenErrorRoleDescriptionExceedsMaxLength)
}

var (
	ErrRoleIsExist                     = errors.New("role is exist")
	ErrRolePermission                  = errors.New("validation failed: Role Permissions invalid permission format")
	ErrRoleDescriptionExceedsMaxLength = errors.New("validation failed: Role Description exceeds max length")
	ErrRoleNameExceedsMaxLength        = errors.New("validation failed: Role Name exceeds max length")
	ErrRoleNameFormat                  = errors.New("validation failed: Role Name invalid format")
)

func newRolseExistError() *domainerrors.ErrorInfo {
	return domainerrors.Wrap(infra_repo.ErrorCodeDatasource, domainerrors.Wrap(infra_sql.ErrorCodeSQLCreate, ErrRoleIsExist))
}

func newRoleExistFullError() *domainerrors.ErrorInfo {
	return domainerrors.Wrap(application_role.ErrorCodeRepository, domainerrors.Wrap(infra_repo.ErrorCodeDatasource, domainerrors.Wrap(infra_sql.ErrorCodeSQLCreate, ErrRoleIsExist)))
}

func newRolePermissionError() *domainerrors.ErrorInfo {
	return domainerrors.Wrap(application_role.ErrorCodeValidateInput, ErrRolePermission)
}

func newRoleNameTooLongError() *domainerrors.ErrorInfo {
	return domainerrors.Wrap(application_role.ErrorCodeValidateInput, ErrRoleNameExceedsMaxLength)
}

func newRoleDescriptionTooLongError() *domainerrors.ErrorInfo {
	return domainerrors.Wrap(application_role.ErrorCodeValidateInput, ErrRoleDescriptionExceedsMaxLength)
}

func newRoleNameFormatError() *domainerrors.ErrorInfo {
	return domainerrors.Wrap(application_role.ErrorCodeValidateInput, ErrRoleNameFormat)
}

func (st *ServiceTest) reset() {
	logger := pwlogger.NewDevelopmentLogger("")

	st.input = nil
	st.expect = nil
	st.repoMock = mocks.NewMockRoleRepository(st.mockCtrl)
	st.userRepoMock = mocks_user.NewMockUserRepository(st.mockCtrl)
	st.transRepoMock = mock_repo.NewMockITransactionRepo(st.mockCtrl)
	st.producer = mock_repo.NewMockProducer(st.mockCtrl)
	st.eventStoreDB = mock_repo.NewMockEventStore(st.mockCtrl)
	st.service = application_role.NewServiceImpl(st.repoMock, st.userRepoMock, st.transRepoMock, st.producer, st.eventStoreDB, logger)
	st.producer = mock_repo.NewMockProducer(st.mockCtrl)
	st.eventStoreDB = mock_repo.NewMockEventStore(st.mockCtrl)
	st.service = application_role.NewServiceImpl(st.repoMock, st.userRepoMock, st.transRepoMock, st.producer, st.eventStoreDB, logger)
}

func (st *ServiceTest) givenData(name, description, permission string) error {
	st.input = &application_role.CreatedInput{
		Name:        name,
		Description: description,
		Permissions: permission,
	}
	st.expect = &application_role.Output{}

	return nil
}

func (st *ServiceTest) whenCreateNewRole(_ context.Context) error {
	e := st.input.ToEntity()
	st.repoMock.EXPECT().Create(gomock.Any(), RoleEquals(e)).Return(e, nil)
	st.producer.EXPECT().PublishEvent(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
	st.eventStoreDB.EXPECT().Store(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
	st.expect.CreatedAt = e.CreatedAt
	st.expect.UpdatedAt = e.UpdatedAt

	return nil
}

func (st *ServiceTest) whenCreateExistingRole(_ context.Context) error {
	e := newRolseExistError()
	st.repoMock.EXPECT().Create(gomock.Any(), RoleEquals(st.input.ToEntity())).Return(nil, e)
	// st.producer.EXPECT().PublishEvent(gomock.Any(), gomock.Any()).Return(nil)

	return nil
}

func (st *ServiceTest) whenInvalidFormat(_ context.Context) error {
	return nil
}

func (st *ServiceTest) thenSuccess(name, description, permission string) error {
	st.expect = &application_role.Output{
		Name:        name,
		Description: description,
		Permissions: strings.Split(permission, ","),
		CreatedAt:   st.expect.CreatedAt,
		UpdatedAt:   st.expect.UpdatedAt,
	}
	actual, err := st.service.CreateRole(context.Background(), st.input)

	if err != nil {
		return err
	}

	err = tests.AssertExpectedAndActual(assert.Equal, st.expect, actual, "Expected %d role created success and equal, but there is %d", st.expect, actual)

	return err
}

func (st *ServiceTest) thenErrorRoleAlreadyExists() error {
	_, err := st.service.CreateRole(context.Background(), st.input)

	var e *domainerrors.ErrorInfo
	if errors.As(err, &e) {
		er := newRoleExistFullError()
		if err != nil {
			// return tests.AssertExpectedAndActual(
			// 	assert.Equal,
			// 	fullerr,
			// 	e,
			// 	"Expected %d role already exists and equal, but there is %d", fullerr, e)
			return tests.CompareFields(*er, *e, []string{"Code", "Message"})
		}
	}

	return fmt.Errorf("Expected error type is domainerrors.ErrorInfo, but there is %w", err)
}

func (st *ServiceTest) thenErrorInvalidPermissionFormat() error {
	_, err := st.service.CreateRole(context.Background(), st.input)

	var e *domainerrors.ErrorInfo
	if errors.As(err, &e) {
		er := newRolePermissionError()
		if err != nil {
			// return tests.AssertExpectedAndActual(assert.Equal, er, e, "Expected %d invalid permission format and equal, but there is %d", er, e)
			return tests.CompareFields(*er, *e, []string{"Code", "Message"})
		}
	}

	return fmt.Errorf("Expected error type is domainerrors.ErrorInfo, but there is %w", err)
}

func (st *ServiceTest) thenErrorInvalidNameFormat() error {
	_, err := st.service.CreateRole(context.Background(), st.input)

	var e *domainerrors.ErrorInfo
	if errors.As(err, &e) {
		er := newRoleNameFormatError()
		if err != nil {
			// return tests.AssertExpectedAndActual(assert.Equal, er, e, "Expected %d invalid name format and equal, but there is %d", er, e)
			return tests.CompareFields(*er, *e, []string{"Code", "Message"})
		}
	}

	return fmt.Errorf("Expected error type is domainerrors.ErrorInfo, but there is %w", err)
}

func (st *ServiceTest) thenErrorRoleNameExceedsMaxLength() error {
	_, err := st.service.CreateRole(context.Background(), st.input)

	var e *domainerrors.ErrorInfo
	if errors.As(err, &e) {
		er := newRoleNameTooLongError()
		if err != nil {
			// return tests.AssertExpectedAndActual(assert.Equal, er, e, "Expected %d role name or description exceeds max length and equal, but there is %d", er, e)
			return tests.CompareFields(*er, *e, []string{"Code", "Message"})
		}
	}

	return fmt.Errorf("Expected error type is domainerrors.ErrorInfo, but there is %w", err)
}

func (st *ServiceTest) thenErrorRoleDescriptionExceedsMaxLength() error {
	_, err := st.service.CreateRole(context.Background(), st.input)

	var e *domainerrors.ErrorInfo
	if errors.As(err, &e) {
		er := newRoleDescriptionTooLongError()
		if err != nil {
			// return tests.AssertExpectedAndActual(assert.Equal, er, e, "Expected %d role name or description exceeds max length and equal, but there is %d", er, e)
			return tests.CompareFields(*er, *e, []string{"Code", "Message"})
		}
	}

	return fmt.Errorf("Expected error type is domainerrors.ErrorInfo, but there is %w", err)
}

func TestCreate(t *testing.T) {
	t.Parallel()

	serviceTest := &ServiceTest{
		t: t,
	}

	// Create the report directory
	reportPath := filepath.Join("..", "report", "TestRoleCreateUsecase.json")
	// Create the directory if it does not exist
	err := os.MkdirAll(filepath.Dir(reportPath), os.ModePerm)
	if err != nil {
		t.Fatal(err)
	}

	// Run the test suite
	suite := godog.TestSuite{
		Name:                "Create",
		ScenarioInitializer: serviceTest.InitializeScenario,
		Options: &godog.Options{
			Format:   "cucumber:" + reportPath,
			Paths:    []string{"features/usecase/role_created.feature"},
			TestingT: t,
			Output:   colors.Colored(os.Stdout),
		},
	}

	if suite.Run() != 0 {
		t.Log("non-zero status returned, failed to run feature tests")
	}
}
