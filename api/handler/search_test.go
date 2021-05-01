// Package handler is ui layer http-handler package
package handler

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/oshou/AwesomeMusic-api/api/usecase"
)

func TestNewSearchHandler(t *testing.T) {
	type args struct {
		usecase usecase.ISearchUsecase
	}
	tests := []struct {
		name string
		args args
		want ISearchHandler
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSearchHandler(tt.args.usecase); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSearchHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_searchHandler_SearchByType(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		sh   *searchHandler
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.sh.SearchByType(tt.args.w, tt.args.r)
		})
	}
}
