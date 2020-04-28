package service

import (
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/oshou/AwesomeMusic-api/domain/model"
	"github.com/oshou/AwesomeMusic-api/mock/mock_repository"
)

func TestCommentService_GetComments(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cases := []struct {
		postID   int
		comments []*model.Comment
		err      error
	}{
		{1, nil, nil},
		{1, []*model.Comment{}, nil},
		{1, []*model.Comment{{ID: 1, UserID: 1, PostID: 1, Comment: "comment1"}}, nil},
		{1, []*model.Comment{
			{ID: 1, UserID: 1, PostID: 1, Comment: "comment1"},
			{ID: 1, UserID: 1, PostID: 2, Comment: "comment2"},
			{ID: 1, UserID: 2, PostID: 3, Comment: "comment3"},
		}, nil},
	}

	for _, tc := range cases {
		mockRepo := mock_repository.NewMockICommentRepository(ctrl)
		mockRepo.EXPECT().GetAll(tc.postID).Return(tc.comments, tc.err)
		commentService := NewCommentService(mockRepo)
		comments, err := commentService.GetComments(tc.postID)

		if err != nil {
			t.Errorf("[Failed] %v", err)
		}

		if !reflect.DeepEqual(comments, tc.comments) {
			t.Errorf("[Failed] expected:%v, actual:%v", tc.comments, comments)
		}
	}
}

func TestCommentService_GetCommentByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cases := []struct {
		commentID int
		comments  *model.Comment
		err       error
	}{
		{1, nil, nil},
		{1, &model.Comment{}, nil},
		{1, &model.Comment{ID: 1, UserID: 1, PostID: 1, Comment: "comment1"}, nil},
	}

	for _, tc := range cases {
		// Actual
		mockRepo := mock_repository.NewMockICommentRepository(ctrl)
		mockRepo.EXPECT().GetByID(tc.commentID).Return(tc.comments, tc.err)
		commentService := NewCommentService(mockRepo)
		comments, err := commentService.GetCommentByID(tc.commentID)

		if err != nil {
			t.Errorf("[Failed] %v", err)
		}

		if !reflect.DeepEqual(comments, tc.comments) {
			t.Errorf("[Failed] expected:%v, actual:%v", tc.comments, comments)
		}
	}
}

func TestCommentService_AddComment(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cases := []struct {
		postID      int
		userID      int
		commentText string
		comment     *model.Comment
		err         error
	}{
		{1, 1, "aaa", nil, nil},
		{1, 1, "aaa", &model.Comment{}, nil},
		{1, 1, "aaa", &model.Comment{ID: 1, UserID: 1, PostID: 1, Comment: "comment1"}, nil},
	}

	for _, tc := range cases {
		// Actual
		mockRepo := mock_repository.NewMockICommentRepository(ctrl)
		mockRepo.EXPECT().Add(tc.postID, tc.userID, tc.commentText).Return(tc.comment, tc.err)
		commentService := NewCommentService(mockRepo)
		comments, err := commentService.AddComment(tc.postID, tc.userID, tc.commentText)

		if err != nil {
			t.Errorf("[Failed] %v", err)
		}

		if !reflect.DeepEqual(comments, tc.comment) {
			t.Errorf("[Failed] expected:%v, actual:%v", tc.comment, comments)
		}
	}
}
