// Code generated by MockGen. DO NOT EDIT.
// Source: comment.go

// Package mock_usecase is a generated GoMock package.
package mock_usecase

import (
	gomock "github.com/golang/mock/gomock"
	model "github.com/oshou/AwesomeMusic-api/api/domain/model"
	reflect "reflect"
)

// MockICommentUsecase is a mock of ICommentUsecase interface
type MockICommentUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockICommentUsecaseMockRecorder
}

// MockICommentUsecaseMockRecorder is the mock recorder for MockICommentUsecase
type MockICommentUsecaseMockRecorder struct {
	mock *MockICommentUsecase
}

// NewMockICommentUsecase creates a new mock instance
func NewMockICommentUsecase(ctrl *gomock.Controller) *MockICommentUsecase {
	mock := &MockICommentUsecase{ctrl: ctrl}
	mock.recorder = &MockICommentUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockICommentUsecase) EXPECT() *MockICommentUsecaseMockRecorder {
	return m.recorder
}

// ListComments mocks base method
func (m *MockICommentUsecase) ListComments(postID int) ([]*model.Comment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListComments", postID)
	ret0, _ := ret[0].([]*model.Comment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListComments indicates an expected call of ListComments
func (mr *MockICommentUsecaseMockRecorder) ListComments(postID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListComments", reflect.TypeOf((*MockICommentUsecase)(nil).ListComments), postID)
}

// GetCommentByID mocks base method
func (m *MockICommentUsecase) GetCommentByID(commentID int) (*model.Comment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCommentByID", commentID)
	ret0, _ := ret[0].(*model.Comment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCommentByID indicates an expected call of GetCommentByID
func (mr *MockICommentUsecaseMockRecorder) GetCommentByID(commentID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCommentByID", reflect.TypeOf((*MockICommentUsecase)(nil).GetCommentByID), commentID)
}

// AddComment mocks base method
func (m *MockICommentUsecase) AddComment(postID, userID int, comment string) (*model.Comment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddComment", postID, userID, comment)
	ret0, _ := ret[0].(*model.Comment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddComment indicates an expected call of AddComment
func (mr *MockICommentUsecaseMockRecorder) AddComment(postID, userID, comment interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddComment", reflect.TypeOf((*MockICommentUsecase)(nil).AddComment), postID, userID, comment)
}
