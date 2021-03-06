// Code generated by MockGen. DO NOT EDIT.
// Source: domain/repository/comment.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"

	model "github.com/oshou/AwesomeMusic-api/api/domain/model"
)

// MockICommentRepository is a mock of ICommentRepository interface
type MockICommentRepository struct {
	ctrl     *gomock.Controller
	recorder *MockICommentRepositoryMockRecorder
}

// MockICommentRepositoryMockRecorder is the mock recorder for MockICommentRepository
type MockICommentRepositoryMockRecorder struct {
	mock *MockICommentRepository
}

// NewMockICommentRepository creates a new mock instance
func NewMockICommentRepository(ctrl *gomock.Controller) *MockICommentRepository {
	mock := &MockICommentRepository{ctrl: ctrl}
	mock.recorder = &MockICommentRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockICommentRepository) EXPECT() *MockICommentRepositoryMockRecorder {
	return m.recorder
}

// GetAll mocks base method
func (m *MockICommentRepository) GetAll(postID int) ([]*model.Comment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", postID)
	ret0, _ := ret[0].([]*model.Comment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll
func (mr *MockICommentRepositoryMockRecorder) GetAll(postID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockICommentRepository)(nil).GetAll), postID)
}

// GetByID mocks base method
func (m *MockICommentRepository) GetByID(commentID int) (*model.Comment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", commentID)
	ret0, _ := ret[0].(*model.Comment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID
func (mr *MockICommentRepositoryMockRecorder) GetByID(commentID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockICommentRepository)(nil).GetByID), commentID)
}

// Add mocks base method
func (m *MockICommentRepository) Add(postID, userID int, comment string) (*model.Comment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", postID, userID, comment)
	ret0, _ := ret[0].(*model.Comment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Add indicates an expected call of Add
func (mr *MockICommentRepositoryMockRecorder) Add(postID, userID, comment interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockICommentRepository)(nil).Add), postID, userID, comment)
}
