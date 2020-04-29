package service

import (
	"reflect"
	"testing"

	"github.com/oshou/AwesomeMusic-api/domain/model"
	"github.com/oshou/AwesomeMusic-api/domain/repository"
)

func TestNewCommentService(t *testing.T) {
	type args struct {
		repo repository.ICommentRepository
	}
	tests := []struct {
		name string
		args args
		want ICommentService
	}{
		//{"nil-input", args{repo: &repository.ICommentRepository}, ICommentService},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCommentService(tt.args.repo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCommentService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_commentService_GetComments(t *testing.T) {
	type args struct {
		postID int
	}
	tests := []struct {
		name    string
		cu      *commentService
		args    args
		want    []*model.Comment
		wantErr bool
	}{
		//{
		//	"nil-input",
		//	&commentService{},
		//	args{postID: 1},
		//	[]*model.Comment{
		//		{ID: 1, UserID: 1, PostID: 1, Comment: "test01"},
		//	},
		//	false,
		//},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.cu.GetComments(tt.args.postID)
			if (err != nil) != tt.wantErr {
				t.Errorf("commentService.GetComments() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("commentService.GetComments() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_commentService_GetCommentByID(t *testing.T) {
	type args struct {
		commentID int
	}
	tests := []struct {
		name    string
		cu      *commentService
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
				t.Errorf("commentService.GetCommentByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("commentService.GetCommentByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_commentService_AddComment(t *testing.T) {
	type args struct {
		postID      int
		userID      int
		commentText string
	}
	tests := []struct {
		name    string
		cu      *commentService
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
				t.Errorf("commentService.AddComment() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("commentService.AddComment() = %v, want %v", got, tt.want)
			}
		})
	}
}
