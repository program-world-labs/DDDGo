// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/logger/logger.go

// Package mock is a generated GoMock package.
package mock

import (
	io "io"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	logger "github.com/program-world-labs/DDDGo/pkg/logger"
	trace "go.opentelemetry.io/otel/trace"
)

// MockInterface is a mock of Interface interface.
type MockInterface struct {
	ctrl     *gomock.Controller
	recorder *MockInterfaceMockRecorder
}

// MockInterfaceMockRecorder is the mock recorder for MockInterface.
type MockInterfaceMockRecorder struct {
	mock *MockInterface
}

// NewMockInterface creates a new mock instance.
func NewMockInterface(ctrl *gomock.Controller) *MockInterface {
	mock := &MockInterface{ctrl: ctrl}
	mock.recorder = &MockInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInterface) EXPECT() *MockInterfaceMockRecorder {
	return m.recorder
}

// Debug mocks base method.
func (m *MockInterface) Debug() *logger.Event {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Debug")
	ret0, _ := ret[0].(*logger.Event)
	return ret0
}

// Debug indicates an expected call of Debug.
func (mr *MockInterfaceMockRecorder) Debug() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Debug", reflect.TypeOf((*MockInterface)(nil).Debug))
}

// Err mocks base method.
func (m *MockInterface) Err(arg0 error) *logger.Event {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Err", arg0)
	ret0, _ := ret[0].(*logger.Event)
	return ret0
}

// Err indicates an expected call of Err.
func (mr *MockInterfaceMockRecorder) Err(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Err", reflect.TypeOf((*MockInterface)(nil).Err), arg0)
}

// Error mocks base method.
func (m *MockInterface) Error() *logger.Event {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Error")
	ret0, _ := ret[0].(*logger.Event)
	return ret0
}

// Error indicates an expected call of Error.
func (mr *MockInterfaceMockRecorder) Error() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Error", reflect.TypeOf((*MockInterface)(nil).Error))
}

// ErrorSpan mocks base method.
func (m *MockInterface) ErrorSpan(arg0 trace.Span) *logger.Event {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ErrorSpan", arg0)
	ret0, _ := ret[0].(*logger.Event)
	return ret0
}

// ErrorSpan indicates an expected call of ErrorSpan.
func (mr *MockInterfaceMockRecorder) ErrorSpan(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ErrorSpan", reflect.TypeOf((*MockInterface)(nil).ErrorSpan), arg0)
}

// Fatal mocks base method.
func (m *MockInterface) Fatal() *logger.Event {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Fatal")
	ret0, _ := ret[0].(*logger.Event)
	return ret0
}

// Fatal indicates an expected call of Fatal.
func (mr *MockInterfaceMockRecorder) Fatal() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fatal", reflect.TypeOf((*MockInterface)(nil).Fatal))
}

// Info mocks base method.
func (m *MockInterface) Info() *logger.Event {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Info")
	ret0, _ := ret[0].(*logger.Event)
	return ret0
}

// Info indicates an expected call of Info.
func (mr *MockInterfaceMockRecorder) Info() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Info", reflect.TypeOf((*MockInterface)(nil).Info))
}

// Log mocks base method.
func (m *MockInterface) Log() *logger.Event {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Log")
	ret0, _ := ret[0].(*logger.Event)
	return ret0
}

// Log indicates an expected call of Log.
func (mr *MockInterfaceMockRecorder) Log() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Log", reflect.TypeOf((*MockInterface)(nil).Log))
}

// Output mocks base method.
func (m *MockInterface) Output(arg0 io.Writer) logger.Logger {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Output", arg0)
	ret0, _ := ret[0].(logger.Logger)
	return ret0
}

// Output indicates an expected call of Output.
func (mr *MockInterfaceMockRecorder) Output(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Output", reflect.TypeOf((*MockInterface)(nil).Output), arg0)
}

// Panic mocks base method.
func (m *MockInterface) Panic() *logger.Event {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Panic")
	ret0, _ := ret[0].(*logger.Event)
	return ret0
}

// Panic indicates an expected call of Panic.
func (mr *MockInterfaceMockRecorder) Panic() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Panic", reflect.TypeOf((*MockInterface)(nil).Panic))
}

// Print mocks base method.
func (m *MockInterface) Print(arg0 ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Print", varargs...)
}

// Print indicates an expected call of Print.
func (mr *MockInterfaceMockRecorder) Print(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Print", reflect.TypeOf((*MockInterface)(nil).Print), arg0...)
}

// Printf mocks base method.
func (m *MockInterface) Printf(arg0 string, arg1 ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Printf", varargs...)
}

// Printf indicates an expected call of Printf.
func (mr *MockInterfaceMockRecorder) Printf(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Printf", reflect.TypeOf((*MockInterface)(nil).Printf), varargs...)
}

// Span mocks base method.
func (m *MockInterface) Span(arg0 trace.Span) *logger.Event {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Span", arg0)
	ret0, _ := ret[0].(*logger.Event)
	return ret0
}

// Span indicates an expected call of Span.
func (mr *MockInterfaceMockRecorder) Span(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Span", reflect.TypeOf((*MockInterface)(nil).Span), arg0)
}

// Trace mocks base method.
func (m *MockInterface) Trace() *logger.Event {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Trace")
	ret0, _ := ret[0].(*logger.Event)
	return ret0
}

// Trace indicates an expected call of Trace.
func (mr *MockInterfaceMockRecorder) Trace() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Trace", reflect.TypeOf((*MockInterface)(nil).Trace))
}

// Warn mocks base method.
func (m *MockInterface) Warn() *logger.Event {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Warn")
	ret0, _ := ret[0].(*logger.Event)
	return ret0
}

// Warn indicates an expected call of Warn.
func (mr *MockInterfaceMockRecorder) Warn() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Warn", reflect.TypeOf((*MockInterface)(nil).Warn))
}

// Write mocks base method.
func (m *MockInterface) Write(p []byte) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Write", p)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Write indicates an expected call of Write.
func (mr *MockInterfaceMockRecorder) Write(p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Write", reflect.TypeOf((*MockInterface)(nil).Write), p)
}
