// Package handler is ui layer http-handler package
package handler

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/oshou/AwesomeMusic-api/api/usecase"
)

func TestNewPostHandler(t *testing.T) {
	type args struct {
		usecase usecase.IPostUsecase
	}
	tests := []struct {
		name string
		args args
		want IPostHandler
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPostHandler(tt.args.usecase); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPostHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_postHandler_ListPosts(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		ph   *postHandler
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.ph.ListPosts(tt.args.w, tt.args.r)
		})
	}
}

func Test_postHandler_GetPostByID(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		ph   *postHandler
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.ph.GetPostByID(tt.args.w, tt.args.r)
		})
	}
}

func Test_postHandler_ListPostsByTagID(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		ph   *postHandler
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.ph.ListPostsByTagID(tt.args.w, tt.args.r)
		})
	}
}

func Test_postHandler_ListPostsByUserID(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		ph   *postHandler
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.ph.ListPostsByUserID(tt.args.w, tt.args.r)
		})
	}
}

func Test_postHandler_AddPost(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		ph   *postHandler
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.ph.AddPost(tt.args.w, tt.args.r)
		})
	}
}

func Test_postHandler_DeletePostByID(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		ph   *postHandler
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.ph.DeletePostByID(tt.args.w, tt.args.r)
		})
	}
}
