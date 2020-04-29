// Package service is application layer package
package service

import (
	"reflect"
	"testing"

	"github.com/oshou/AwesomeMusic-api/domain/model"
	"github.com/oshou/AwesomeMusic-api/domain/repository"
)

func TestNewTagService(t *testing.T) {
	type args struct {
		repo repository.ITagRepository
	}
	tests := []struct {
		name string
		args args
		want ITagService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTagService(tt.args.repo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTagService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_tagService_GetTags(t *testing.T) {
	tests := []struct {
		name    string
		tu      *tagService
		want    []*model.Tag
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.tu.GetTags()
			if (err != nil) != tt.wantErr {
				t.Errorf("tagService.GetTags() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("tagService.GetTags() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_tagService_GetTagByID(t *testing.T) {
	type args struct {
		tagID int
	}
	tests := []struct {
		name    string
		tu      *tagService
		args    args
		want    *model.Tag
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.tu.GetTagByID(tt.args.tagID)
			if (err != nil) != tt.wantErr {
				t.Errorf("tagService.GetTagByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("tagService.GetTagByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_tagService_GetTagsByPostID(t *testing.T) {
	type args struct {
		postID int
	}
	tests := []struct {
		name    string
		tu      *tagService
		args    args
		want    []*model.Tag
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.tu.GetTagsByPostID(tt.args.postID)
			if (err != nil) != tt.wantErr {
				t.Errorf("tagService.GetTagsByPostID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("tagService.GetTagsByPostID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_tagService_AddTag(t *testing.T) {
	type args struct {
		tagName string
	}
	tests := []struct {
		name    string
		tu      *tagService
		args    args
		want    *model.Tag
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.tu.AddTag(tt.args.tagName)
			if (err != nil) != tt.wantErr {
				t.Errorf("tagService.AddTag() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("tagService.AddTag() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_tagService_AttachTag(t *testing.T) {
	type args struct {
		postID int
		tagID  int
	}
	tests := []struct {
		name    string
		tu      *tagService
		args    args
		want    *model.PostTag
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.tu.AttachTag(tt.args.postID, tt.args.tagID)
			if (err != nil) != tt.wantErr {
				t.Errorf("tagService.AttachTag() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("tagService.AttachTag() = %v, want %v", got, tt.want)
			}
		})
	}
}
