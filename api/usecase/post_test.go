//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=../mock/$GOPACKAGE/$GOFILE
// Package usecase is application layer package
package usecase

import (
	"reflect"
	"testing"

	"github.com/oshou/AwesomeMusic-api/api/domain/model"
	"github.com/oshou/AwesomeMusic-api/api/domain/repository"
)

func TestNewPostUsecase(t *testing.T) {
	type args struct {
		repo repository.IPostRepository
	}
	tests := []struct {
		name string
		args args
		want IPostUsecase
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPostUsecase(tt.args.repo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPostUsecase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_postUsecase_ListPosts(t *testing.T) {
	tests := []struct {
		name    string
		pu      *postUsecase
		want    []*model.Post
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.pu.ListPosts()
			if (err != nil) != tt.wantErr {
				t.Errorf("postUsecase.ListPosts() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("postUsecase.ListPosts() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_postUsecase_GetPostByID(t *testing.T) {
	type args struct {
		postID int
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

func Test_postUsecase_ListPostsByTagID(t *testing.T) {
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
			got, err := tt.pu.ListPostsByTagID(tt.args.tagID)
			if (err != nil) != tt.wantErr {
				t.Errorf("postUsecase.ListPostsByTagID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("postUsecase.ListPostsByTagID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_postUsecase_ListPostsByUserID(t *testing.T) {
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
			got, err := tt.pu.ListPostsByUserID(tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("postUsecase.ListPostsByUserID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("postUsecase.ListPostsByUserID() = %v, want %v", got, tt.want)
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
