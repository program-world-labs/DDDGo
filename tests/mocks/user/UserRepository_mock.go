// Code generated by MockGen. DO NOT EDIT.
// Source: internal/domain/user/repository/user_repository.go

// Package user is a generated GoMock package.
package user

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	domain "github.com/program-world-labs/DDDGo/internal/domain"
)

// MockUserRepository is a mock of UserRepository interface.
type MockUserRepository struct {
	ctrl     *gomock.Controller
	recorder *MockUserRepositoryMockRecorder
}

// MockUserRepositoryMockRecorder is the mock recorder for MockUserRepository.
type MockUserRepositoryMockRecorder struct {
	mock *MockUserRepository
}

// NewMockUserRepository creates a new mock instance.
func NewMockUserRepository(ctrl *gomock.Controller) *MockUserRepository {
	mock := &MockUserRepository{ctrl: ctrl}
	mock.recorder = &MockUserRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserRepository) EXPECT() *MockUserRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockUserRepository) Create(ctx context.Context, e domain.IEntity) (domain.IEntity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, e)
	ret0, _ := ret[0].(domain.IEntity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockUserRepositoryMockRecorder) Create(ctx, e interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockUserRepository)(nil).Create), ctx, e)
}

// CreateTx mocks base method.
func (m *MockUserRepository) CreateTx(arg0 context.Context, arg1 domain.IEntity, arg2 domain.ITransactionEvent) (domain.IEntity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTx", arg0, arg1, arg2)
	ret0, _ := ret[0].(domain.IEntity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTx indicates an expected call of CreateTx.
func (mr *MockUserRepositoryMockRecorder) CreateTx(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTx", reflect.TypeOf((*MockUserRepository)(nil).CreateTx), arg0, arg1, arg2)
}

// Delete mocks base method.
func (m *MockUserRepository) Delete(ctx context.Context, e domain.IEntity) (domain.IEntity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, e)
	ret0, _ := ret[0].(domain.IEntity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockUserRepositoryMockRecorder) Delete(ctx, e interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockUserRepository)(nil).Delete), ctx, e)
}

// DeleteTx mocks base method.
func (m *MockUserRepository) DeleteTx(arg0 context.Context, arg1 domain.IEntity, arg2 domain.ITransactionEvent) (domain.IEntity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTx", arg0, arg1, arg2)
	ret0, _ := ret[0].(domain.IEntity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteTx indicates an expected call of DeleteTx.
func (mr *MockUserRepositoryMockRecorder) DeleteTx(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTx", reflect.TypeOf((*MockUserRepository)(nil).DeleteTx), arg0, arg1, arg2)
}

// GetAll mocks base method.
func (m *MockUserRepository) GetAll(ctx context.Context, sq *domain.SearchQuery, e domain.IEntity) (*domain.List, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", ctx, sq, e)
	ret0, _ := ret[0].(*domain.List)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockUserRepositoryMockRecorder) GetAll(ctx, sq, e interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockUserRepository)(nil).GetAll), ctx, sq, e)
}

// GetByID mocks base method.
func (m *MockUserRepository) GetByID(ctx context.Context, e domain.IEntity) (domain.IEntity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", ctx, e)
	ret0, _ := ret[0].(domain.IEntity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockUserRepositoryMockRecorder) GetByID(ctx, e interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockUserRepository)(nil).GetByID), ctx, e)
}

// Update mocks base method.
func (m *MockUserRepository) Update(ctx context.Context, e domain.IEntity) (domain.IEntity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, e)
	ret0, _ := ret[0].(domain.IEntity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockUserRepositoryMockRecorder) Update(ctx, e interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockUserRepository)(nil).Update), ctx, e)
}

// UpdateTx mocks base method.
func (m *MockUserRepository) UpdateTx(arg0 context.Context, arg1 domain.IEntity, arg2 domain.ITransactionEvent) (domain.IEntity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTx", arg0, arg1, arg2)
	ret0, _ := ret[0].(domain.IEntity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateTx indicates an expected call of UpdateTx.
func (mr *MockUserRepositoryMockRecorder) UpdateTx(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTx", reflect.TypeOf((*MockUserRepository)(nil).UpdateTx), arg0, arg1, arg2)
}

// UpdateWithFields mocks base method.
func (m *MockUserRepository) UpdateWithFields(ctx context.Context, e domain.IEntity, keys []string) (domain.IEntity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateWithFields", ctx, e, keys)
	ret0, _ := ret[0].(domain.IEntity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateWithFields indicates an expected call of UpdateWithFields.
func (mr *MockUserRepositoryMockRecorder) UpdateWithFields(ctx, e, keys interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateWithFields", reflect.TypeOf((*MockUserRepository)(nil).UpdateWithFields), ctx, e, keys)
}

// UpdateWithFieldsTx mocks base method.
func (m *MockUserRepository) UpdateWithFieldsTx(arg0 context.Context, arg1 domain.IEntity, arg2 []string, arg3 domain.ITransactionEvent) (domain.IEntity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateWithFieldsTx", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(domain.IEntity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateWithFieldsTx indicates an expected call of UpdateWithFieldsTx.
func (mr *MockUserRepositoryMockRecorder) UpdateWithFieldsTx(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateWithFieldsTx", reflect.TypeOf((*MockUserRepository)(nil).UpdateWithFieldsTx), arg0, arg1, arg2, arg3)
}
