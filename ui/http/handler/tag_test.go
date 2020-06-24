package handler_test

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/oshou/AwesomeMusic-api/usecase"
)

func TestNewTagHandler(t *testing.T) {
	type args struct {
		usecase usecase.ITagUsecase
	}
	tests := []struct {
		name string
		args args
		want ITagHandler
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTagHandler(tt.args.usecase); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTagHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_tagHandler_GetTags(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		th   *tagHandler
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.th.GetTags(tt.args.w, tt.args.r)
		})
	}
}

func Test_tagHandler_GetTagByID(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		th   *tagHandler
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.th.GetTagByID(tt.args.w, tt.args.r)
		})
	}
}

func Test_tagHandler_GetTagsByPostID(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		th   *tagHandler
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.th.GetTagsByPostID(tt.args.w, tt.args.r)
		})
	}
}

func Test_tagHandler_AddTag(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		th   *tagHandler
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.th.AddTag(tt.args.w, tt.args.r)
		})
	}
}

func Test_tagHandler_AttachTag(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		th   *tagHandler
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.th.AttachTag(tt.args.w, tt.args.r)
		})
	}
}
