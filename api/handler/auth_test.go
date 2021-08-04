package handler

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/gorilla/sessions"
	"github.com/oshou/AwesomeMusic-api/api/usecase"
)

func TestNewAuthHandler(t *testing.T) {
	type args struct {
		u usecase.IUserUsecase
		s sessions.Store
	}
	tests := []struct {
		name string
		args args
		want IAuthHandler
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAuthHandler(tt.args.u, tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAuthHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_authHandler_Login(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		h    *authHandler
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.h.Login(tt.args.w, tt.args.r)
		})
	}
}

func Test_authHandler_Logout(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		h    *authHandler
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.h.Logout(tt.args.w, tt.args.r)
		})
	}
}
