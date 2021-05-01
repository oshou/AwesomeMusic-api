// Code generated by MockGen. DO NOT EDIT.
// Source: session.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	gomock "github.com/golang/mock/gomock"
	model "github.com/oshou/AwesomeMusic-api/api/domain/model"
	reflect "reflect"
)

// MockISessionRepository is a mock of ISessionRepository interface
type MockISessionRepository struct {
	ctrl     *gomock.Controller
	recorder *MockISessionRepositoryMockRecorder
}

// MockISessionRepositoryMockRecorder is the mock recorder for MockISessionRepository
type MockISessionRepositoryMockRecorder struct {
	mock *MockISessionRepository
}

// NewMockISessionRepository creates a new mock instance
func NewMockISessionRepository(ctrl *gomock.Controller) *MockISessionRepository {
	mock := &MockISessionRepository{ctrl: ctrl}
	mock.recorder = &MockISessionRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockISessionRepository) EXPECT() *MockISessionRepositoryMockRecorder {
	return m.recorder
}

// Get mocks base method
func (m *MockISessionRepository) Get(id string) (*model.Session, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", id)
	ret0, _ := ret[0].(*model.Session)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockISessionRepositoryMockRecorder) Get(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockISessionRepository)(nil).Get), id)
}

// Set mocks base method
func (m *MockISessionRepository) Set(arg0 *model.Session) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Set", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Set indicates an expected call of Set
func (mr *MockISessionRepositoryMockRecorder) Set(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockISessionRepository)(nil).Set), arg0)
}