//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=../mock/$GOPACKAGE/$GOFILE
// Package usecase is application layer package
package usecase

import (
	"reflect"
	"testing"

	"github.com/oshou/AwesomeMusic-api/api/domain/model"
	"github.com/oshou/AwesomeMusic-api/api/domain/repository"
)

func TestNewCommentUsecase(t *testing.T) {
	type args struct {
		repo repository.ICommentRepository
	}
	tests := []struct {
		name string
		args args
		want ICommentUsecase
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCommentUsecase(tt.args.repo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCommentUsecase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_commentUsecase_ListComments(t *testing.T) {
	type args struct {
		postID int
	}
	tests := []struct {
		name    string
		cu      *commentUsecase
		args    args
		want    []*model.Comment
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.cu.ListComments(tt.args.postID)
			if (err != nil) != tt.wantErr {
				t.Errorf("commentUsecase.ListComments() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("commentUsecase.ListComments() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_commentUsecase_GetCommentByID(t *testing.T) {
	type args struct {
		commentID int
	}
	tests := []struct {
		name    string
		cu      *commentUsecase
		args    args
		want    *model.Comment
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.cu.GetCommentByID(tt.args.commentID)
			if (err != nil) != tt.wantErr {
				t.Errorf("commentUsecase.GetCommentByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("commentUsecase.GetCommentByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_commentUsecase_AddComment(t *testing.T) {
	type args struct {
		postID      int
		userID      int
		commentText string
	}
	tests := []struct {
		name    string
		cu      *commentUsecase
		args    args
		want    *model.Comment
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.cu.AddComment(tt.args.postID, tt.args.userID, tt.args.commentText)
			if (err != nil) != tt.wantErr {
				t.Errorf("commentUsecase.AddComment() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("commentUsecase.AddComment() = %v, want %v", got, tt.want)
			}
		})
	}
}
