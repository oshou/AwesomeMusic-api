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

	tests := []struct {
		postID  int
		want    []*model.Comment
		wantErr error
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

	for _, tt := range tests {
		mockRepo := mock_repository.NewMockICommentRepository(ctrl)
		mockRepo.EXPECT().GetAll(tt.postID).Return(tt.want, tt.err)
		cs := NewCommentService(mockRepo)
		comments, err := cs.GetComments(tt.postID)

		if err != nil {
			t.Errorf("[Failed] %v", err)
		}

		if !reflect.DeepEqual(tt.comments, comments) {
			t.Errorf("[Failed] want:%v, got:%v", tt.comments, comments)
		}
	}
}

func TestCommentService_GetCommentByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := []struct {
		commentID int
		comments  *model.Comment
		err       error
	}{
		{1, nil, nil},
		{1, &model.Comment{}, nil},
		{1, &model.Comment{ID: 1, UserID: 1, PostID: 1, Comment: "comment1"}, nil},
	}

	for _, tt := range tests {
		// Actual
		mockRepo := mock_repository.NewMockICommentRepository(ctrl)
		mockRepo.EXPECT().GetByID(tt.commentID).Return(tt.comments, tt.err)
		cs := NewCommentService(mockRepo)
		comment, err := cs.GetCommentByID(tt.commentID)

		if err != nil {
			t.Errorf("[Failed] %v", err)
		}

		if !reflect.DeepEqual(tt.comments, comment) {
			t.Errorf("[Failed] want:%v, got:%v", tt.comments, comment)
		}
	}
}

func TestCommentService_AddComment(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := []struct {
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

	for _, tt := range tests {
		// Actual
		mockRepo := mock_repository.NewMockICommentRepository(ctrl)
		mockRepo.EXPECT().Add(tt.postID, tt.userID, tt.commentText).Return(tt.comment, tt.err)
		cs := NewCommentService(mockRepo)
		comment, err := cs.AddComment(tt.postID, tt.userID, tt.commentText)

		if err != nil {
			t.Errorf("[Failed] %v", err)
		}

		if !reflect.DeepEqual(tt.comment, comment) {
			t.Errorf("[Failed] want:%v, got:%v", tt.comment, comment)
		}
	}
}
