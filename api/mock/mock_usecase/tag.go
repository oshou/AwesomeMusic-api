// Code generated by MockGen. DO NOT EDIT.
// Source: tag.go

// Package mock_usecase is a generated GoMock package.
package mock_usecase

import (
	gomock "github.com/golang/mock/gomock"
	model "github.com/oshou/AwesomeMusic-api/api/domain/model"
	reflect "reflect"
)

// MockITagUsecase is a mock of ITagUsecase interface
type MockITagUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockITagUsecaseMockRecorder
}

// MockITagUsecaseMockRecorder is the mock recorder for MockITagUsecase
type MockITagUsecaseMockRecorder struct {
	mock *MockITagUsecase
}

// NewMockITagUsecase creates a new mock instance
func NewMockITagUsecase(ctrl *gomock.Controller) *MockITagUsecase {
	mock := &MockITagUsecase{ctrl: ctrl}
	mock.recorder = &MockITagUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockITagUsecase) EXPECT() *MockITagUsecaseMockRecorder {
	return m.recorder
}

// ListTags mocks base method
func (m *MockITagUsecase) ListTags() ([]*model.Tag, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTags")
	ret0, _ := ret[0].([]*model.Tag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTags indicates an expected call of ListTags
func (mr *MockITagUsecaseMockRecorder) ListTags() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTags", reflect.TypeOf((*MockITagUsecase)(nil).ListTags))
}

// ListTagsByPostID mocks base method
func (m *MockITagUsecase) ListTagsByPostID(postID int) ([]*model.Tag, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTagsByPostID", postID)
	ret0, _ := ret[0].([]*model.Tag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTagsByPostID indicates an expected call of ListTagsByPostID
func (mr *MockITagUsecaseMockRecorder) ListTagsByPostID(postID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTagsByPostID", reflect.TypeOf((*MockITagUsecase)(nil).ListTagsByPostID), postID)
}

// GetTagByID mocks base method
func (m *MockITagUsecase) GetTagByID(tagID int) (*model.Tag, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTagByID", tagID)
	ret0, _ := ret[0].(*model.Tag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTagByID indicates an expected call of GetTagByID
func (mr *MockITagUsecaseMockRecorder) GetTagByID(tagID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTagByID", reflect.TypeOf((*MockITagUsecase)(nil).GetTagByID), tagID)
}

// AddTag mocks base method
func (m *MockITagUsecase) AddTag(tagName string) (*model.Tag, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddTag", tagName)
	ret0, _ := ret[0].(*model.Tag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddTag indicates an expected call of AddTag
func (mr *MockITagUsecaseMockRecorder) AddTag(tagName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddTag", reflect.TypeOf((*MockITagUsecase)(nil).AddTag), tagName)
}

// AttachTag mocks base method
func (m *MockITagUsecase) AttachTag(postID, tagID int) (*model.PostTag, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AttachTag", postID, tagID)
	ret0, _ := ret[0].(*model.PostTag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AttachTag indicates an expected call of AttachTag
func (mr *MockITagUsecaseMockRecorder) AttachTag(postID, tagID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AttachTag", reflect.TypeOf((*MockITagUsecase)(nil).AttachTag), postID, tagID)
}