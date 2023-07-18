// Code generated by MockGen. DO NOT EDIT.
// Source: internal/application/role/interface.go

// Package role is a generated GoMock package.
package role

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	role "github.com/program-world-labs/DDDGo/internal/application/role"
)

// MockIService is a mock of IService interface.
type MockIService struct {
	ctrl     *gomock.Controller
	recorder *MockIServiceMockRecorder
}

// MockIServiceMockRecorder is the mock recorder for MockIService.
type MockIServiceMockRecorder struct {
	mock *MockIService
}

// NewMockIService creates a new mock instance.
func NewMockIService(ctrl *gomock.Controller) *MockIService {
	mock := &MockIService{ctrl: ctrl}
	mock.recorder = &MockIServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIService) EXPECT() *MockIServiceMockRecorder {
	return m.recorder
}

// CreateRole mocks base method.
func (m *MockIService) CreateRole(ctx context.Context, roleInfo *role.CreatedInput) (*role.Output, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateRole", ctx, roleInfo)
	ret0, _ := ret[0].(*role.Output)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateRole indicates an expected call of CreateRole.
func (mr *MockIServiceMockRecorder) CreateRole(ctx, roleInfo interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRole", reflect.TypeOf((*MockIService)(nil).CreateRole), ctx, roleInfo)
}

// DeleteRole mocks base method.
func (m *MockIService) DeleteRole(ctx context.Context, roleInfo *role.DeletedInput) (*role.Output, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteRole", ctx, roleInfo)
	ret0, _ := ret[0].(*role.Output)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteRole indicates an expected call of DeleteRole.
func (mr *MockIServiceMockRecorder) DeleteRole(ctx, roleInfo interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteRole", reflect.TypeOf((*MockIService)(nil).DeleteRole), ctx, roleInfo)
}


// GetRoleDetail mocks base method.
func (m *MockIService) GetRoleDetail(ctx context.Context, roleInfo *role.DetailGotInput) (*role.Output, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRoleDetail", ctx, roleInfo)
	ret0, _ := ret[0].(*role.Output)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRoleDetail indicates an expected call of GetRoleDetail.
func (mr *MockIServiceMockRecorder) GetRoleDetail(ctx, roleInfo interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRoleDetail", reflect.TypeOf((*MockIService)(nil).GetRoleDetail), ctx, roleInfo)
}

// GetRoleList mocks base method.
func (m *MockIService) GetRoleList(ctx context.Context, roleInfo *role.ListGotInput) (*role.OutputList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRoleList", ctx, roleInfo)
	ret0, _ := ret[0].(*role.OutputList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRoleList indicates an expected call of GetRoleList.
func (mr *MockIServiceMockRecorder) GetRoleList(ctx, roleInfo interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRoleList", reflect.TypeOf((*MockIService)(nil).GetRoleList), ctx, roleInfo)
}

// UpdateRole mocks base method.
func (m *MockIService) UpdateRole(ctx context.Context, roleInfo *role.UpdatedInput) (*role.Output, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateRole", ctx, roleInfo)
	ret0, _ := ret[0].(*role.Output)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateRole indicates an expected call of UpdateRole.
func (mr *MockIServiceMockRecorder) UpdateRole(ctx, roleInfo interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateRole", reflect.TypeOf((*MockIService)(nil).UpdateRole), ctx, roleInfo)
}
