// Package service is application layer package
package service

import (
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/oshou/AwesomeMusic-api/domain/model"
	"github.com/oshou/AwesomeMusic-api/mock/mock_repository"
)

func TestPostService_GetPosts(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := []struct {
		posts []*model.Post
		err   error
	}{}

	for _, tt := range tests {
		mockRepo := mock_repository.NewMockIPostRepository(ctrl)
		mockRepo.EXPECT().GetAll().Return(tt.posts, tt.err)
		ps := NewPostService(mockRepo)
		posts, err := ps.GetPosts()

		if err != nil {
			t.Errorf("[Failed] %v", err)
		}

		if !reflect.DeepEqual(tt.posts, posts) {
			t.Errorf("[Failed] want:%v, got:%v", tt.posts, posts)
		}
	}
}

func TestPostService_GetPostByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := []struct {
		postID int
		post   *model.Post
		err    error
	}{}

	for _, tt := range tests {
		mockRepo := mock_repository.NewMockIPostRepository(ctrl)
		mockRepo.EXPECT().GetByID(tt.postID).Return(tt.post, tt.err)
		ps := NewPostService(mockRepo)
		post, err := ps.GetPostByID(tt.postID)

		if err != nil {
			t.Errorf("[Failed] %v", err)
		}

		if !reflect.DeepEqual(tt.post, post) {
			t.Errorf("[Failed] want:%v, got:%v", tt.post, post)
		}
	}
}

func TestPostService_GetPostsByTagID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := []struct {
		tagID int
		posts []*model.Post
		err   error
	}{}

	for _, tt := range tests {
		mockRepo := mock_repository.NewMockIPostRepository(ctrl)
		mockRepo.EXPECT().GetByTagID(tt.tagID).Return(tt.posts, tt.err)
		ps := NewPostService(mockRepo)
		posts, err := ps.GetPostsByTagID(tt.tagID)

		if err != nil {
			t.Errorf("[Failed] %v", err)
		}

		if !reflect.DeepEqual(tt.posts, posts) {
			t.Errorf("[Failed] want:%v, got:%v", tt.posts, posts)
		}
	}
}

func TestPostService_GetPostsByUserID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := []struct {
		userID int
		posts  []*model.Post
		err    error
	}{}

	for _, tt := range tests {
		mockRepo := mock_repository.NewMockIPostRepository(ctrl)
		mockRepo.EXPECT().GetByUserID(tt.userID).Return(tt.posts, tt.err)
		ps := NewPostService(mockRepo)
		posts, err := ps.GetPostsByUserID(tt.userID)

		if err != nil {
			t.Errorf("[Failed] %v", err)
		}

		if !reflect.DeepEqual(tt.posts, posts) {
			t.Errorf("[Failed] want:%v, got:%v", tt.posts, posts)
		}
	}
}

func TestPostService_AddPost(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := []struct {
		userID  int
		title   string
		url     string
		message string
		post    *model.Post
		err     error
	}{}

	for _, tc := range tests {
		mockRepo := mock_repository.NewMockIPostRepository(ctrl)
		mockRepo.EXPECT().Add(tc.userID, tc.title, tc.url, tc.message).Return(tc.post, tc.err)
		ps := NewPostService(mockRepo)
		post, err := ps.AddPost(tc.userID, tc.title, tc.url, tc.message)

		if err != nil {
			t.Errorf("[Failed] %v", err)
		}

		if !reflect.DeepEqual(tc.post, post) {
			t.Errorf("[Failed] want:%v, got:%v", tc.post, post)
		}
	}
}
