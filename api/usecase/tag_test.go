// Package usecase is application layer package
package usecase_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"

	"github.com/oshou/AwesomeMusic-api/api/domain/model"
	"github.com/oshou/AwesomeMusic-api/api/domain/repository"
	"github.com/oshou/AwesomeMusic-api/api/mock/mock_repository"
	"github.com/oshou/AwesomeMusic-api/api/usecase"
)

func TestNewTagUsecase(t *testing.T) {
	tests := []struct {
		name string
		repo repository.ITagRepository
		want usecase.ITagUsecase
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := usecase.NewTagUsecase(tt.repo); !cmp.Equal(got, tt.want) {
				t.Errorf("NewTagUsecase() = %v, want %v\ndiff=%v", got, tt.want, cmp.Diff(got, tt.want))
			}
		})
	}
}

func Test_tagUsecase_ListTags(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := []struct {
		name    string
		mock    []*model.Tag
		mockErr error
		want    []*model.Tag
		wantErr error
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mock := mock_repository.NewMockITagRepository(ctrl)
			mock.EXPECT().List().Return(tt.mock, tt.mockErr)
			tu := usecase.NewTagUsecase(mock)
			got, err := tu.ListTags()

			if err != tt.wantErr {
				t.Errorf("tagUsecase.ListTags() error (wantErr %v, gotErr %v)", tt.wantErr, err)
			}

			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("tagUsecase.ListTags() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func Test_tagUsecase_GetTagByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := []struct {
		name    string
		tagID   int
		mock    *model.Tag
		mockErr error
		want    *model.Tag
		wantErr error
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mock := mock_repository.NewMockITagRepository(ctrl)
			mock.EXPECT().GetByID(tt.tagID).Return(tt.mock, tt.mockErr)
			tu := usecase.NewTagUsecase(mock)
			got, err := tu.GetTagByID(tt.tagID)

			if err != tt.wantErr {
				t.Errorf("tagUsecase.GetTagByID() error (wantErr %v, gotErr %v)", tt.wantErr, err)
			}

			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("tagUsecase.GetTagByID() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func Test_tagUsecase_ListTagsByPostID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := []struct {
		name    string
		postID  int
		mock    []*model.Tag
		mockErr error
		want    []*model.Tag
		wantErr error
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mock := mock_repository.NewMockITagRepository(ctrl)
			mock.EXPECT().GetByPostID(tt.postID).Return(tt.mock, tt.mockErr)
			tu := usecase.NewTagUsecase(mock)
			got, err := tu.ListTagsByPostID(tt.postID)

			if err != tt.wantErr {
				t.Errorf("tagUsecase.ListTagsByPostID() error (wantErr %v, gotErr %v)", tt.wantErr, err)
			}

			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("tagUsecase.ListTagsByPostID() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func Test_tagUsecase_AddTag(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := []struct {
		name    string
		tagName string
		mock    *model.Tag
		mockErr error
		want    *model.Tag
		wantErr error
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mock := mock_repository.NewMockITagRepository(ctrl)
			mock.EXPECT().Add(tt.tagName).Return(tt.mock, tt.mockErr)
			tu := usecase.NewTagUsecase(mock)
			got, err := tu.AddTag(tt.tagName)

			if err != tt.wantErr {
				t.Errorf("tagUsecase.AddTag() error (wantErr %v, gotErr %v)", tt.wantErr, err)
			}

			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("tagUsecase.AddTag() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func Test_tagUsecase_AttachTag(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := []struct {
		name    string
		postID  int
		tagID   int
		mock    *model.PostTag
		mockErr error
		want    *model.PostTag
		wantErr error
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mock := mock_repository.NewMockITagRepository(ctrl)
			mock.EXPECT().Attach(tt.postID, tt.tagID).Return(tt.mock, tt.mockErr)
			tu := usecase.NewTagUsecase(mock)
			got, err := tu.AttachTag(tt.postID, tt.tagID)

			if err != tt.wantErr {
				t.Errorf("tagUsecase.AttachTag() error (wantErr %v, gotErr %v)", tt.wantErr, err)
			}

			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("tagUsecase.AttachTag() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
