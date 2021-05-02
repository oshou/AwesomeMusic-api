//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=../mock/$GOPACKAGE/$GOFILE
// Package usecase is application layer package
package usecase

import (
	"reflect"
	"testing"

	"github.com/oshou/AwesomeMusic-api/api/domain/model"
	"github.com/oshou/AwesomeMusic-api/api/domain/repository"
)

func TestNewTagUsecase(t *testing.T) {
	type args struct {
		repo repository.ITagRepository
	}
	tests := []struct {
		name string
		args args
		want ITagUsecase
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTagUsecase(tt.args.repo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTagUsecase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_tagUsecase_ListTags(t *testing.T) {
	tests := []struct {
		name    string
		tu      *tagUsecase
		want    []*model.Tag
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.tu.ListTags()
			if (err != nil) != tt.wantErr {
				t.Errorf("tagUsecase.ListTags() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("tagUsecase.ListTags() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_tagUsecase_ListTagsByPostID(t *testing.T) {
	type args struct {
		postID int
	}
	tests := []struct {
		name    string
		tu      *tagUsecase
		args    args
		want    []*model.Tag
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.tu.ListTagsByPostID(tt.args.postID)
			if (err != nil) != tt.wantErr {
				t.Errorf("tagUsecase.ListTagsByPostID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("tagUsecase.ListTagsByPostID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_tagUsecase_GetTagByID(t *testing.T) {
	type args struct {
		tagID int
	}
	tests := []struct {
		name    string
		tu      *tagUsecase
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
				t.Errorf("tagUsecase.GetTagByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("tagUsecase.GetTagByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_tagUsecase_AddTag(t *testing.T) {
	type args struct {
		tagName string
	}
	tests := []struct {
		name    string
		tu      *tagUsecase
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
				t.Errorf("tagUsecase.AddTag() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("tagUsecase.AddTag() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_tagUsecase_AttachTag(t *testing.T) {
	type args struct {
		postID int
		tagID  int
	}
	tests := []struct {
		name    string
		tu      *tagUsecase
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
				t.Errorf("tagUsecase.AttachTag() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("tagUsecase.AttachTag() = %v, want %v", got, tt.want)
			}
		})
	}
}
