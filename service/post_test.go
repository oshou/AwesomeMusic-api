// Package service is application layer package
package service

import (
	"reflect"
	"testing"

	"github.com/oshou/AwesomeMusic-api/domain/model"
	"github.com/oshou/AwesomeMusic-api/domain/repository"
)

func TestNewPostService(t *testing.T) {
	type args struct {
		repo repository.IPostRepository
	}

	tests := []struct {
		name string
		args args
		want IPostService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPostService(tt.args.repo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPostService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_postService_GetPosts(t *testing.T) {
	tests := []struct {
		name    string
		pu      *postService
		want    []*model.Post
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.pu.GetPosts()
			if (err != nil) != tt.wantErr {
				t.Errorf("postService.GetPosts() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("postService.GetPosts() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_postService_GetPostByID(t *testing.T) {
	type args struct {
		postID int
	}
	tests := []struct {
		name    string
		pu      *postService
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
				t.Errorf("postService.GetPostByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("postService.GetPostByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_postService_GetPostsByTagID(t *testing.T) {
	type args struct {
		tagID int
	}
	tests := []struct {
		name    string
		pu      *postService
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
				t.Errorf("postService.GetPostsByTagID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("postService.GetPostsByTagID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_postService_GetPostsByUserID(t *testing.T) {
	type args struct {
		userID int
	}
	tests := []struct {
		name    string
		pu      *postService
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
				t.Errorf("postService.GetPostsByUserID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("postService.GetPostsByUserID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_postService_AddPost(t *testing.T) {
	type args struct {
		userID  int
		title   string
		url     string
		message string
	}
	tests := []struct {
		name    string
		pu      *postService
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
				t.Errorf("postService.AddPost() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("postService.AddPost() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_postService_DeletePostByID(t *testing.T) {
	type args struct {
		postID int
	}
	tests := []struct {
		name    string
		pu      *postService
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.pu.DeletePostByID(tt.args.postID); (err != nil) != tt.wantErr {
				t.Errorf("postService.DeletePostByID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
