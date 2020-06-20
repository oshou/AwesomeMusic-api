package usecase

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"

	"github.com/oshou/AwesomeMusic-api/domain/model"
	"github.com/oshou/AwesomeMusic-api/mock/mock_repository"
)

func Test_commentUsecase_GetComments(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := []struct {
		name    string
		postID  int
		mock    []*model.Comment
		mockErr error
		want    []*model.Comment
		wantErr error
	}{
		{
			name:   "success",
			postID: 1,
			mock: []*model.Comment{
				{ID: 1, UserID: 1, PostID: 1, Comment: "sample01"},
				{ID: 2, UserID: 1, PostID: 1, Comment: "sample02"},
				{ID: 3, UserID: 2, PostID: 1, Comment: "sample03"},
			},
			mockErr: nil,
			want: []*model.Comment{
				{ID: 1, UserID: 1, PostID: 1, Comment: "sample01"},
				{ID: 2, UserID: 1, PostID: 1, Comment: "sample02"},
				{ID: 3, UserID: 2, PostID: 1, Comment: "sample03"},
			},
			wantErr: nil,
		},
		{
			name:    "no data",
			postID:  0,
			mock:    []*model.Comment{},
			mockErr: nil,
			want:    []*model.Comment{},
			wantErr: nil,
		},
		{
			name:    "repository error",
			postID:  1,
			mock:    []*model.Comment{},
			mockErr: errors.New("test"),
			want:    []*model.Comment{},
			wantErr: errors.New("test"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mock := mock_repository.NewMockICommentRepository(ctrl)
			mock.EXPECT().GetAll(tt.postID).Return(tt.mock, tt.mockErr)
			cu := &commentUsecase{repo: mock}
			got, err := cu.GetComments(tt.postID)

			if err != tt.wantErr {
				t.Errorf("commentUsecase.GetComments() error (wantErr %v, gotErr %v)", tt.wantErr, err)
			}

			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("commentUsecase.GetComments() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func Test_commentUsecase_GetCommentByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := []struct {
		name      string
		commentID int
		mock      *model.Comment
		mockErr   error
		want      *model.Comment
		wantErr   error
	}{
		{
			name:      "success",
			commentID: 1,
			mock:      &model.Comment{},
			mockErr:   nil,
			want:      &model.Comment{},
			wantErr:   nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mock := mock_repository.NewMockICommentRepository(ctrl)
			mock.EXPECT().GetByID(tt.commentID).Return(tt.mock, tt.mockErr)
			cu := &commentUsecase{repo: mock}
			got, err := cu.GetCommentByID(tt.commentID)

			if err != tt.wantErr {
				t.Errorf("commentUsecase.GetCommentByID() error (wantErr %v, gotErr %v)", tt.wantErr, err)
			}

			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("commentUsecase.GetCommentByID() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func Test_commentUsecase_AddComment(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := []struct {
		name        string
		postID      int
		userID      int
		commentText string
		mock        *model.Comment
		mockErr     error
		want        *model.Comment
		wantErr     error
	}{
		// TODO: Add test cases.
		{
			name:        "success",
			postID:      1,
			userID:      1,
			commentText: "hello",
			mock:        &model.Comment{},
			mockErr:     nil,
			want:        &model.Comment{},
			wantErr:     nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mock := mock_repository.NewMockICommentRepository(ctrl)
			mock.EXPECT().Add(tt.postID, tt.userID, tt.commentText).Return(tt.mock, tt.mockErr)
			cu := &commentUsecase{repo: mock}
			got, err := cu.AddComment(tt.postID, tt.userID, tt.commentText)

			if err != tt.wantErr {
				t.Errorf("commentUsecase.AddComment() error (wantErr %v, gotErr %v)", tt.wantErr, err)
			}

			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("commentUsecase.AddComment() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
