// Package handler is ui layer http-handler package
package handler

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/oshou/AwesomeMusic-api/usecase"
)

func TestNewCommentHandler(t *testing.T) {
	type args struct {
		usecase usecase.ICommentUsecase
	}
	tests := []struct {
		name string
		args args
		want ICommentHandler
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCommentHandler(tt.args.usecase); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCommentHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_commentHandler_GetComments(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		ch   *commentHandler
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.ch.GetComments(tt.args.w, tt.args.r)
		})
	}
}

func Test_commentHandler_AddComment(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		ch   *commentHandler
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.ch.AddComment(tt.args.w, tt.args.r)
		})
	}
}

func Test_commentHandler_GetCommentByID(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		ch   *commentHandler
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.ch.GetCommentByID(tt.args.w, tt.args.r)
		})
	}
}
