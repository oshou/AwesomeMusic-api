// Package usecase is application layer package
package usecase

import (
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
	"github.com/oshou/AwesomeMusic-api/domain/model"
	"github.com/oshou/AwesomeMusic-api/mock/mock_repository"
)

func Test_postUsecase_GetPosts(t *testing.T) {
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
			mock.EXPECT().GetAll().Return(tt.mock, tt.mockErr)
			pu := &postUsecase{repo: mock}
			got, err := tt.pu.GetPosts()

			if err != tt.wantErr {
				t.Errorf("postUsecase.GetPosts() error (wantErr %v, gotErr %v)", tt.wantErr, err)
			}

			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("postUsecase.GetPosts() mismatch (-want +got):\n%s", diff)
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
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mock := mock_repository.NewMockIPostRepository(ctrl)
			mock.EXPECT().GetByID(tt.postID)
			got, err := tt.pu.GetPostByID(tt.args.postID)
			if (err != nil) != tt.wantErr {
				t.Errorf("postUsecase.GetPostByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("postUsecase.GetPostByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_postUsecase_GetPostsByTagID(t *testing.T) {
	type args struct {
		tagID int
	}
	tests := []struct {
		name    string
		pu      *postUsecase
		args    args
		want    []*model.Post
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.pu.GetPostsByTagID(tt.args.tagID)
			if (err != nil) != tt.wantErr {
				t.Errorf("postUsecase.GetPostsByTagID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("postUsecase.GetPostsByTagID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_postUsecase_GetPostsByUserID(t *testing.T) {
	type args struct {
		userID int
	}
	tests := []struct {
		name    string
		pu      *postUsecase
		args    args
		want    []*model.Post
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.pu.GetPostsByUserID(tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("postUsecase.GetPostsByUserID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("postUsecase.GetPostsByUserID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_postUsecase_AddPost(t *testing.T) {
	type args struct {
		userID  int
		title   string
		url     string
		message string
	}
	tests := []struct {
		name    string
		pu      *postUsecase
		args    args
		want    *model.Post
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.pu.AddPost(tt.args.userID, tt.args.title, tt.args.url, tt.args.message)
			if (err != nil) != tt.wantErr {
				t.Errorf("postUsecase.AddPost() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("postUsecase.AddPost() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_postUsecase_DeletePostByID(t *testing.T) {
	type args struct {
		postID int
	}
	tests := []struct {
		name    string
		pu      *postUsecase
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.pu.DeletePostByID(tt.args.postID); (err != nil) != tt.wantErr {
				t.Errorf("postUsecase.DeletePostByID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
