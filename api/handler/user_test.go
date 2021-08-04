// Package handler is ui layer http-handler package
package handler

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/oshou/AwesomeMusic-api/api/usecase"
)

func TestNewUserHandler(t *testing.T) {
	type args struct {
		usecase usecase.IUserUsecase
	}
	tests := []struct {
		name string
		args args
		want IUserHandler
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUserHandler(tt.args.usecase); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userHandler_ListUsers(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		uh   *userHandler
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.uh.ListUsers(tt.args.w, tt.args.r)
		})
	}
}

func Test_userHandler_AddUser(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		uh   *userHandler
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.uh.AddUser(tt.args.w, tt.args.r)
		})
	}
}

func Test_userHandler_GetUserByID(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		uh   *userHandler
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.uh.GetUserByID(tt.args.w, tt.args.r)
		})
	}
}
