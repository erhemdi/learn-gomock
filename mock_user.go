// Code generated by MockGen. DO NOT EDIT.
// Source: user.go

// Package user is a generated GoMock package.
package user

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockUserRedisIface is a mock of UserRedisIface interface.
type MockUserRedisIface struct {
	ctrl     *gomock.Controller
	recorder *MockUserRedisIfaceMockRecorder
}

// MockUserRedisIfaceMockRecorder is the mock recorder for MockUserRedisIface.
type MockUserRedisIfaceMockRecorder struct {
	mock *MockUserRedisIface
}

// NewMockUserRedisIface creates a new mock instance.
func NewMockUserRedisIface(ctrl *gomock.Controller) *MockUserRedisIface {
	mock := &MockUserRedisIface{ctrl: ctrl}
	mock.recorder = &MockUserRedisIfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserRedisIface) EXPECT() *MockUserRedisIfaceMockRecorder {
	return m.recorder
}

// GetData mocks base method.
func (m *MockUserRedisIface) GetData(id int64) (User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetData", id)
	ret0, _ := ret[0].(User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetData indicates an expected call of GetData.
func (mr *MockUserRedisIfaceMockRecorder) GetData(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetData", reflect.TypeOf((*MockUserRedisIface)(nil).GetData), id)
}

// MockUserDBIface is a mock of UserDBIface interface.
type MockUserDBIface struct {
	ctrl     *gomock.Controller
	recorder *MockUserDBIfaceMockRecorder
}

// MockUserDBIfaceMockRecorder is the mock recorder for MockUserDBIface.
type MockUserDBIfaceMockRecorder struct {
	mock *MockUserDBIface
}

// NewMockUserDBIface creates a new mock instance.
func NewMockUserDBIface(ctrl *gomock.Controller) *MockUserDBIface {
	mock := &MockUserDBIface{ctrl: ctrl}
	mock.recorder = &MockUserDBIfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserDBIface) EXPECT() *MockUserDBIfaceMockRecorder {
	return m.recorder
}

// GetData mocks base method.
func (m *MockUserDBIface) GetData(id int64) (User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetData", id)
	ret0, _ := ret[0].(User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetData indicates an expected call of GetData.
func (mr *MockUserDBIfaceMockRecorder) GetData(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetData", reflect.TypeOf((*MockUserDBIface)(nil).GetData), id)
}