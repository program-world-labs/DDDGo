package user_test

import (
	context "context"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/cucumber/godog"
	gomock "github.com/golang/mock/gomock"
	"github.com/program-world-labs/pwlogger"
	"github.com/stretchr/testify/assert"

	application_user "github.com/program-world-labs/DDDGo/internal/application/user"
	"github.com/program-world-labs/DDDGo/internal/domain/user/entity"
	"github.com/program-world-labs/DDDGo/tests"
	mocks_user "github.com/program-world-labs/DDDGo/tests/mocks/user"
)

// godogsCtxKey is the key used to store the available godogs in the context.Context.
type userServiceTest struct {
	t                *testing.T // 新增這一行
	userRepoMock     *mocks_user.MockUserRepository
	userService      *application_user.ServiceImpl
	e                *entity.User
	o                *application_user.Output
	shouldMockCreate bool
}

// 註冊的使用者已經存在.
func (t *userServiceTest) givenUserExists(ctx context.Context) error {
	t.t.Log("givenUserExists")
	t.userRepoMock.EXPECT().GetByID(ctx, t.e).Return(t.e, nil)
	t.shouldMockCreate = false

	return nil
}

// 註冊的使用者不存在.
func (t *userServiceTest) givenUserDoesNotExist(ctx context.Context) error {
	t.t.Log("givenUserDoesNotExist")
	t.userRepoMock.EXPECT().GetByID(ctx, t.e).Return(nil, nil)
	t.shouldMockCreate = true

	return nil
}

// 使用者已經登入.
func (t *userServiceTest) givenUserIsLoggedIn(_ context.Context) error {
	t.t.Log("givenUserIsLoggedIn")

	return nil
}

// 註冊一個新用戶.
func (t *userServiceTest) whenRegisterANewUser(ctx context.Context) error {
	t.t.Log("theUserIsRegistered")

	if t.shouldMockCreate {
		t.userRepoMock.EXPECT().Create(ctx, t.e).Return(t.e, nil)
	}

	return nil
}

// 取得使用者個人資料.
func (t *userServiceTest) whenGetProfileOfUser(ctx context.Context) error {
	t.t.Log("theGetProfileOfUser")
	t.userRepoMock.EXPECT().GetByID(ctx, t.e).Return(t.e, nil)

	return nil
}

var ErrCreateUser = errors.New("failed to create user")

// 用戶註冊成功.
func (t *userServiceTest) thenUserIsRegisteredSuccessfully(ctx context.Context) error {
	t.t.Log("thenUserIsRegisteredSuccessfully")
	createdUser, err := t.userService.RegisterUseCase(ctx, t.e)
	// assert.NoError(t.userRepoMock.AssertExpectations(t.T()))
	if err != nil {
		return err
	}

	return tests.AssertExpectedAndActual(assert.Equal, t.o, createdUser, "Expected %w , but there is %w", t.o, createdUser)
}

// 用戶註冊失敗.
func (t *userServiceTest) thenUserIsNotRegisteredSuccessfully(ctx context.Context) error {
	t.t.Log("thenUserIsNotRegisteredSuccessfully")
	_, err := t.userService.RegisterUseCase(ctx, t.e)

	if err == nil {
		return fmt.Errorf("expected an error, but got nil: %w", ErrCreateUser)
	}

	return nil
}

// 取得使用者個人資料成功.
func (t *userServiceTest) thenGetProfileOfUserSuccessfully(ctx context.Context) error {
	t.t.Log("thenGetProfileOfUserSuccessfully")
	createdUser, err := t.userService.GetByIDUseCase(ctx, t.e.ID)
	// assert.NoError(t.userRepoMock.AssertExpectations(t.T()))
	if err != nil {
		return err
	}

	return tests.AssertExpectedAndActual(assert.Equal, t.o, createdUser, "Expected %w , but there is %w", t.o, createdUser)
}

func TestUserUsecase(t *testing.T) {
	t.Parallel()

	mockCtl := gomock.NewController(t)
	defer mockCtl.Finish()

	repo := mocks_user.NewMockUserRepository(gomock.NewController(t))
	logger := pwlogger.NewDevelopmentLogger("")
	service := application_user.NewServiceImpl(repo, logger)
	u, err := entity.NewUser("test")

	if err != nil {
		t.Fatal(err)
	}

	o := application_user.NewOutput(u)

	userServiceTest := &userServiceTest{
		t:                t,
		userRepoMock:     repo,
		userService:      service,
		e:                u,
		o:                o,
		shouldMockCreate: false,
	}

	reportPath := filepath.Join("..", "report", "TestUserUsecase.json")
	// Create the directory if it does not exist
	err = os.MkdirAll(filepath.Dir(reportPath), os.ModePerm)
	if err != nil {
		t.Fatal(err)
	}

	suite := godog.TestSuite{
		Name:                "Register",
		ScenarioInitializer: userServiceTest.InitializeScenario,
		Options: &godog.Options{
			Format:   "cucumber:" + reportPath,
			Paths:    []string{"features/user.feature"},
			TestingT: t,
		},
	}

	if suite.Run() != 0 {
		t.Log("non-zero status returned, failed to run feature tests")
	}
}

func (t *userServiceTest) InitializeScenario(ctx *godog.ScenarioContext) {
	// Given
	ctx.Step(`^註冊的使用者不存在$`, t.givenUserDoesNotExist)
	ctx.Step(`^註冊的使用者已經存在$`, t.givenUserExists)
	ctx.Step(`^使用者已經登入$`, t.givenUserIsLoggedIn)

	// When
	ctx.Step(`^註冊一個新用戶$`, t.whenRegisterANewUser)
	ctx.Step(`^取得使用者個人資料$`, t.whenGetProfileOfUser)

	// Then
	ctx.Step(`^用戶註冊成功$`, t.thenUserIsRegisteredSuccessfully)
	ctx.Step(`^用戶註冊失敗$`, t.thenUserIsNotRegisteredSuccessfully)
	ctx.Step(`^取得使用者個人資料成功$`, t.thenGetProfileOfUserSuccessfully)
}
