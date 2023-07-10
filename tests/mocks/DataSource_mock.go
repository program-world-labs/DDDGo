// Code generated by MockGen. DO NOT EDIT.
// Source: internal/infra/datasource/interface.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	domain "github.com/program-world-labs/DDDGo/internal/domain"
	dto "github.com/program-world-labs/DDDGo/internal/infra/dto"
)

// MockIDataSource is a mock of IDataSource interface.
type MockIDataSource struct {
	ctrl     *gomock.Controller
	recorder *MockIDataSourceMockRecorder
}

// MockIDataSourceMockRecorder is the mock recorder for MockIDataSource.
type MockIDataSourceMockRecorder struct {
	mock *MockIDataSource
}

// NewMockIDataSource creates a new mock instance.
func NewMockIDataSource(ctrl *gomock.Controller) *MockIDataSource {
	mock := &MockIDataSource{ctrl: ctrl}
	mock.recorder = &MockIDataSourceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIDataSource) EXPECT() *MockIDataSourceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockIDataSource) Create(arg0 context.Context, arg1 dto.IRepoEntity) (dto.IRepoEntity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(dto.IRepoEntity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockIDataSourceMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockIDataSource)(nil).Create), arg0, arg1)
}

// CreateTx mocks base method.
func (m *MockIDataSource) CreateTx(arg0 context.Context, arg1 dto.IRepoEntity, arg2 domain.ITransactionEvent) (dto.IRepoEntity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTx", arg0, arg1, arg2)
	ret0, _ := ret[0].(dto.IRepoEntity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTx indicates an expected call of CreateTx.
func (mr *MockIDataSourceMockRecorder) CreateTx(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTx", reflect.TypeOf((*MockIDataSource)(nil).CreateTx), arg0, arg1, arg2)
}

// Delete mocks base method.
func (m *MockIDataSource) Delete(arg0 context.Context, arg1 dto.IRepoEntity) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockIDataSourceMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockIDataSource)(nil).Delete), arg0, arg1)
}

// DeleteTx mocks base method.
func (m *MockIDataSource) DeleteTx(arg0 context.Context, arg1 dto.IRepoEntity, arg2 domain.ITransactionEvent) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTx", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTx indicates an expected call of DeleteTx.
func (mr *MockIDataSourceMockRecorder) DeleteTx(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTx", reflect.TypeOf((*MockIDataSource)(nil).DeleteTx), arg0, arg1, arg2)
}

// GetAll mocks base method.
func (m *MockIDataSource) GetAll(arg0 context.Context, arg1 dto.IRepoEntity, arg2 *domain.SearchQuery) ([]dto.IRepoEntity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", arg0, arg1, arg2)
	ret0, _ := ret[0].([]dto.IRepoEntity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockIDataSourceMockRecorder) GetAll(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockIDataSource)(nil).GetAll), arg0, arg1, arg2)
}

// GetByID mocks base method.
func (m *MockIDataSource) GetByID(arg0 context.Context, arg1 dto.IRepoEntity) (dto.IRepoEntity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", arg0, arg1)
	ret0, _ := ret[0].(dto.IRepoEntity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockIDataSourceMockRecorder) GetByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockIDataSource)(nil).GetByID), arg0, arg1)
}

// Update mocks base method.
func (m *MockIDataSource) Update(arg0 context.Context, arg1 dto.IRepoEntity) (dto.IRepoEntity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(dto.IRepoEntity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockIDataSourceMockRecorder) Update(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockIDataSource)(nil).Update), arg0, arg1)
}

// UpdateTx mocks base method.
func (m *MockIDataSource) UpdateTx(arg0 context.Context, arg1 dto.IRepoEntity, arg2 domain.ITransactionEvent) (dto.IRepoEntity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTx", arg0, arg1, arg2)
	ret0, _ := ret[0].(dto.IRepoEntity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateTx indicates an expected call of UpdateTx.
func (mr *MockIDataSourceMockRecorder) UpdateTx(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTx", reflect.TypeOf((*MockIDataSource)(nil).UpdateTx), arg0, arg1, arg2)
}

// UpdateWithFields mocks base method.
func (m *MockIDataSource) UpdateWithFields(arg0 context.Context, arg1 dto.IRepoEntity, arg2 []string) (dto.IRepoEntity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateWithFields", arg0, arg1, arg2)
	ret0, _ := ret[0].(dto.IRepoEntity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateWithFields indicates an expected call of UpdateWithFields.
func (mr *MockIDataSourceMockRecorder) UpdateWithFields(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateWithFields", reflect.TypeOf((*MockIDataSource)(nil).UpdateWithFields), arg0, arg1, arg2)
}

// UpdateWithFieldsTx mocks base method.
func (m *MockIDataSource) UpdateWithFieldsTx(arg0 context.Context, arg1 dto.IRepoEntity, arg2 []string, arg3 domain.ITransactionEvent) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateWithFieldsTx", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateWithFieldsTx indicates an expected call of UpdateWithFieldsTx.
func (mr *MockIDataSourceMockRecorder) UpdateWithFieldsTx(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateWithFieldsTx", reflect.TypeOf((*MockIDataSource)(nil).UpdateWithFieldsTx), arg0, arg1, arg2, arg3)
}

// MockICacheDataSource is a mock of ICacheDataSource interface.
type MockICacheDataSource struct {
	ctrl     *gomock.Controller
	recorder *MockICacheDataSourceMockRecorder
}

// MockICacheDataSourceMockRecorder is the mock recorder for MockICacheDataSource.
type MockICacheDataSourceMockRecorder struct {
	mock *MockICacheDataSource
}

// NewMockICacheDataSource creates a new mock instance.
func NewMockICacheDataSource(ctrl *gomock.Controller) *MockICacheDataSource {
	mock := &MockICacheDataSource{ctrl: ctrl}
	mock.recorder = &MockICacheDataSourceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockICacheDataSource) EXPECT() *MockICacheDataSourceMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockICacheDataSource) Delete(ctx context.Context, e dto.IRepoEntity) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, e)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockICacheDataSourceMockRecorder) Delete(ctx, e interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockICacheDataSource)(nil).Delete), ctx, e)
}

// Get mocks base method.
func (m *MockICacheDataSource) Get(ctx context.Context, e dto.IRepoEntity, ttl ...time.Duration) (dto.IRepoEntity, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, e}
	for _, a := range ttl {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Get", varargs...)
	ret0, _ := ret[0].(dto.IRepoEntity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockICacheDataSourceMockRecorder) Get(ctx, e interface{}, ttl ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, e}, ttl...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockICacheDataSource)(nil).Get), varargs...)
}

// Set mocks base method.
func (m *MockICacheDataSource) Set(ctx context.Context, e dto.IRepoEntity, ttl ...time.Duration) (dto.IRepoEntity, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, e}
	for _, a := range ttl {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Set", varargs...)
	ret0, _ := ret[0].(dto.IRepoEntity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Set indicates an expected call of Set.
func (mr *MockICacheDataSourceMockRecorder) Set(ctx, e interface{}, ttl ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, e}, ttl...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockICacheDataSource)(nil).Set), varargs...)
}

// MockITransactionRun is a mock of ITransactionRun interface.
type MockITransactionRun struct {
	ctrl     *gomock.Controller
	recorder *MockITransactionRunMockRecorder
}

// MockITransactionRunMockRecorder is the mock recorder for MockITransactionRun.
type MockITransactionRunMockRecorder struct {
	mock *MockITransactionRun
}

// NewMockITransactionRun creates a new mock instance.
func NewMockITransactionRun(ctrl *gomock.Controller) *MockITransactionRun {
	mock := &MockITransactionRun{ctrl: ctrl}
	mock.recorder = &MockITransactionRunMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockITransactionRun) EXPECT() *MockITransactionRunMockRecorder {
	return m.recorder
}

// RunTransaction mocks base method.
func (m *MockITransactionRun) RunTransaction(arg0 context.Context, arg1 domain.TransactionEventFunc) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RunTransaction", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RunTransaction indicates an expected call of RunTransaction.
func (mr *MockITransactionRunMockRecorder) RunTransaction(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunTransaction", reflect.TypeOf((*MockITransactionRun)(nil).RunTransaction), arg0, arg1)
}
