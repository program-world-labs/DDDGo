// Code generated by MockGen. DO NOT EDIT.
// Source: internal/domain/user/repository/role_repository.go

// Package role is a generated GoMock package.
package role

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	domain "github.com/program-world-labs/DDDGo/internal/domain"
)

// MockRoleRepository is a mock of RoleRepository interface.
type MockRoleRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRoleRepositoryMockRecorder
}

// MockRoleRepositoryMockRecorder is the mock recorder for MockRoleRepository.
type MockRoleRepositoryMockRecorder struct {
	mock *MockRoleRepository
}

// NewMockRoleRepository creates a new mock instance.
func NewMockRoleRepository(ctrl *gomock.Controller) *MockRoleRepository {
	mock := &MockRoleRepository{ctrl: ctrl}
	mock.recorder = &MockRoleRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRoleRepository) EXPECT() *MockRoleRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockRoleRepository) Create(ctx context.Context, e domain.IEntity) (domain.IEntity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, e)
	ret0, _ := ret[0].(domain.IEntity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockRoleRepositoryMockRecorder) Create(ctx, e interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockRoleRepository)(nil).Create), ctx, e)
}

// CreateTx mocks base method.
func (m *MockRoleRepository) CreateTx(arg0 context.Context, arg1 domain.IEntity, arg2 domain.ITransactionEvent) (domain.IEntity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTx", arg0, arg1, arg2)
	ret0, _ := ret[0].(domain.IEntity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTx indicates an expected call of CreateTx.
func (mr *MockRoleRepositoryMockRecorder) CreateTx(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTx", reflect.TypeOf((*MockRoleRepository)(nil).CreateTx), arg0, arg1, arg2)
}

// Delete mocks base method.
func (m *MockRoleRepository) Delete(ctx context.Context, e domain.IEntity) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, e)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockRoleRepositoryMockRecorder) Delete(ctx, e interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockRoleRepository)(nil).Delete), ctx, e)
}

// DeleteTx mocks base method.
func (m *MockRoleRepository) DeleteTx(arg0 context.Context, arg1 domain.IEntity, arg2 domain.ITransactionEvent) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTx", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTx indicates an expected call of DeleteTx.
func (mr *MockRoleRepositoryMockRecorder) DeleteTx(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTx", reflect.TypeOf((*MockRoleRepository)(nil).DeleteTx), arg0, arg1, arg2)
}

// GetAll mocks base method.
func (m *MockRoleRepository) GetAll(ctx context.Context, e domain.IEntity, sq *domain.SearchQuery) ([]domain.IEntity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", ctx, e, sq)
	ret0, _ := ret[0].([]domain.IEntity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockRoleRepositoryMockRecorder) GetAll(ctx, e, sq interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockRoleRepository)(nil).GetAll), ctx, e, sq)
}

// GetByID mocks base method.
func (m *MockRoleRepository) GetByID(ctx context.Context, e domain.IEntity) (domain.IEntity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", ctx, e)
	ret0, _ := ret[0].(domain.IEntity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockRoleRepositoryMockRecorder) GetByID(ctx, e interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockRoleRepository)(nil).GetByID), ctx, e)
}

// Update mocks base method.
func (m *MockRoleRepository) Update(ctx context.Context, e domain.IEntity) (domain.IEntity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, e)
	ret0, _ := ret[0].(domain.IEntity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockRoleRepositoryMockRecorder) Update(ctx, e interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockRoleRepository)(nil).Update), ctx, e)
}

// UpdateTx mocks base method.
func (m *MockRoleRepository) UpdateTx(arg0 context.Context, arg1 domain.IEntity, arg2 domain.ITransactionEvent) (domain.IEntity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTx", arg0, arg1, arg2)
	ret0, _ := ret[0].(domain.IEntity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateTx indicates an expected call of UpdateTx.
func (mr *MockRoleRepositoryMockRecorder) UpdateTx(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTx", reflect.TypeOf((*MockRoleRepository)(nil).UpdateTx), arg0, arg1, arg2)
}

// UpdateWithFields mocks base method.
func (m *MockRoleRepository) UpdateWithFields(ctx context.Context, e domain.IEntity, keys []string) (domain.IEntity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateWithFields", ctx, e, keys)
	ret0, _ := ret[0].(domain.IEntity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateWithFields indicates an expected call of UpdateWithFields.
func (mr *MockRoleRepositoryMockRecorder) UpdateWithFields(ctx, e, keys interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateWithFields", reflect.TypeOf((*MockRoleRepository)(nil).UpdateWithFields), ctx, e, keys)
}

// UpdateWithFieldsTx mocks base method.
func (m *MockRoleRepository) UpdateWithFieldsTx(arg0 context.Context, arg1 domain.IEntity, arg2 []string, arg3 domain.ITransactionEvent) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateWithFieldsTx", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateWithFieldsTx indicates an expected call of UpdateWithFieldsTx.
func (mr *MockRoleRepositoryMockRecorder) UpdateWithFieldsTx(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateWithFieldsTx", reflect.TypeOf((*MockRoleRepository)(nil).UpdateWithFieldsTx), arg0, arg1, arg2, arg3)
}