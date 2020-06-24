// Package handler is ui layer http-handler package
package handler

import (
	"net/http"
	"testing"

	"github.com/oshou/AwesomeMusic-api/usecase"

	"github.com/google/go-cmp/cmp"
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
			if got := NewSearchHandler(tt.args.usecase); !cmp.Equal(got, tt.want) {
				t.Errorf("NewSearchHandler() = %v, want %v\ndiff=%v", got, tt.want, cmp.Diff(got, tt.want))
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
