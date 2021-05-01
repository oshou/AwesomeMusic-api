// Package handler is ui layer http-handler package
package handler

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/oshou/AwesomeMusic-api/api/usecase"
)

func TestNewHealthHandler(t *testing.T) {
	type args struct {
		u usecase.IHealthUsecase
	}
	tests := []struct {
		name string
		args args
		want IHealthHandler
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewHealthHandler(tt.args.u); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHealthHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_healthHandler_GetHealth(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		hu   healthHandler
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.hu.GetHealth(tt.args.w, tt.args.r)
		})
	}
}
