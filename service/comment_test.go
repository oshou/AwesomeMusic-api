package service

import (
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/oshou/AwesomeMusic-api/domain/model"
	"github.com/oshou/AwesomeMusic-api/domain/repository"
	"github.com/oshou/AwesomeMusic-api/mock/mock_repository"
)

func TestNewCommentService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := mock_repository.NewMockICommentRepository(ctrl)

	type args struct {
		repo repository.ICommentRepository
	}

	tests := []struct {
		name string
		args args
		want ICommentService
	}{
		{
			name: "new",
			args: args{repo: mockRepo},
			want: &commentService{repo: mockRepo},
		},
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
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type args struct {
		postID int
	}

	tests := []struct {
		name    string
		args    args
		prepare func(*mock_repository.MockICommentRepository)
		want    []*model.Comment
		wantErr bool
	}{
		{
			name: "no-data",
			args: args{postID: 0},
			prepare: func(mockRepo *mock_repository.MockICommentRepository) {
				mockRepo.EXPECT().GetAll(0).Return([]*model.Comment{}, nil)
			},
			want:    []*model.Comment{},
			wantErr: false,
		},
		{
			name: "3-data",
			args: args{postID: 1},
			prepare: func(mockRepo *mock_repository.MockICommentRepository) {
				mockRepo.EXPECT().GetAll(1).Return([]*model.Comment{
					{ID: 1, UserID: 1, PostID: 1, Comment: "sample01"},
					{ID: 2, UserID: 1, PostID: 1, Comment: "sample02"},
					{ID: 3, UserID: 2, PostID: 1, Comment: "sample03"},
				}, nil)
			},
			want: []*model.Comment{
				{ID: 1, UserID: 1, PostID: 1, Comment: "sample01"},
				{ID: 2, UserID: 1, PostID: 1, Comment: "sample02"},
				{ID: 3, UserID: 2, PostID: 1, Comment: "sample03"},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := mock_repository.NewMockICommentRepository(ctrl)
			tt.prepare(mockRepo)
			cu := &commentService{repo: mockRepo}
			got, err := cu.GetComments(tt.args.postID)

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

//func Test_commentService_GetCommentByID(t *testing.T) {
//	type args struct {
//		commentID int
//	}
//
//	tests := []struct {
//		name    string
//		args    args
//		want    *model.Comment
//		wantErr bool
//	}{}
//
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			return aaa()
//		})
//	}
//}

func Test_commentService_GetCommentByID(t *testing.T) {
	type args struct {
		commentID int
		args
	}

	tests := []struct {
		name    string
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
