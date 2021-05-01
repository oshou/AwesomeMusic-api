// Package usecase is application layer package
package usecase_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"

	"github.com/oshou/AwesomeMusic-api/api/domain/model"
	"github.com/oshou/AwesomeMusic-api/api/domain/repository"
	"github.com/oshou/AwesomeMusic-api/api/mock/mock_repository"
	"github.com/oshou/AwesomeMusic-api/api/usecase"
)

func TestNewPostUsecase(t *testing.T) {
	tests := []struct {
		name string
		repo repository.IPostRepository
		want usecase.IPostUsecase
	}{
		// TODO: Add test cases.
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := usecase.NewPostUsecase(tt.repo); !cmp.Equal(got, tt.want) {
				t.Errorf("NewPostUsecase() = %v, want %v\ndiff=%v", got, tt.want, cmp.Diff(got, tt.want))
			}
		})
	}
}
func Test_postUsecase_ListPosts(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := []struct {
		name    string
		mock    []*model.Post
		mockErr error
		want    []*model.Post
		wantErr error
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mock := mock_repository.NewMockIPostRepository(ctrl)
			mock.EXPECT().List().Return(tt.mock, tt.mockErr)
			pu := usecase.NewPostUsecase(mock)
			got, err := pu.ListPosts()

			if err != tt.wantErr {
				t.Errorf("postUsecase.ListPosts() error (wantErr %v, gotErr %v)", tt.wantErr, err)
			}

			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("postUsecase.ListPosts() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func Test_postUsecase_GetPostByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := []struct {
		name    string
		postID  int
		mock    *model.Post
		mockErr error
		want    *model.Post
		wantErr error
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mock := mock_repository.NewMockIPostRepository(ctrl)
			mock.EXPECT().GetByID(tt.postID).Return(tt.mock, tt.mockErr)
			pu := usecase.NewPostUsecase(mock)
			got, err := pu.GetPostByID(tt.postID)

			if err != tt.wantErr {
				t.Errorf("postUsecase.GetPostByID() error (wantErr %v, gotErr %v)", tt.wantErr, err)
			}

			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("postUsecase.GetPostByID() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func Test_postUsecase_ListPostsByTagID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := []struct {
		name    string
		tagID   int
		mock    []*model.Post
		mockErr error
		want    []*model.Post
		wantErr error
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mock := mock_repository.NewMockIPostRepository(ctrl)
			mock.EXPECT().GetByTagID(tt.tagID).Return(tt.mock, tt.mockErr)
			pu := usecase.NewPostUsecase(mock)
			got, err := pu.ListPostsByTagID(tt.tagID)

			if err != tt.wantErr {
				t.Errorf("postUsecase.GetPostByTagID() error (wantErr %v, gotErr %v)", tt.wantErr, err)
			}

			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("postUsecase.GetPostByTagID() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func Test_postUsecase_ListPostsByUserID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := []struct {
		name    string
		userID  int
		mock    []*model.Post
		mockErr error
		want    []*model.Post
		wantErr error
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mock := mock_repository.NewMockIPostRepository(ctrl)
			mock.EXPECT().GetByUserID(tt.userID).Return(tt.mock, tt.mockErr)
			pu := usecase.NewPostUsecase(mock)
			got, err := pu.ListPostsByUserID(tt.userID)

			if err != tt.wantErr {
				t.Errorf("postUsecase.GetPostByUserID() error (wantErr %v, gotErr %v)", tt.wantErr, err)
			}

			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("postUsecase.GetPostByUserID() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func Test_postUsecase_AddPost(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := []struct {
		name    string
		userID  int
		title   string
		url     string
		message string
		mock    *model.Post
		mockErr error
		want    *model.Post
		wantErr error
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mock := mock_repository.NewMockIPostRepository(ctrl)
			mock.EXPECT().Add(tt.userID, tt.title, tt.url, tt.message).Return(tt.mock, tt.mockErr)
			pu := usecase.NewPostUsecase(mock)
			got, err := pu.AddPost(tt.userID, tt.title, tt.url, tt.message)

			if err != tt.wantErr {
				t.Errorf("postUsecase.AddPost() error (wantErr %v, gotErr %v)", tt.wantErr, err)
			}

			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("postUsecase().AddPost() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func Test_postUsecase_DeletePostByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := []struct {
		name    string
		postID  int
		mockErr error
		wantErr error
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mock := mock_repository.NewMockIPostRepository(ctrl)
			mock.EXPECT().DeleteByID(tt.postID).Return(tt.mockErr)
			pu := usecase.NewPostUsecase(mock)
			err := pu.DeletePostByID(tt.postID)

			if err != tt.wantErr {
				t.Errorf("postUsecase.AddPost() error (wantErr %v, gotErr %v)", tt.wantErr, err)
			}
		})
	}
}
