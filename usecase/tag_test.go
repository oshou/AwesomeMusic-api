// Package usecase is application layer package
package usecase

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"

	"github.com/oshou/AwesomeMusic-api/domain/model"
	"github.com/oshou/AwesomeMusic-api/mock/mock_repository"
)

func Test_tagUsecase_GetTags(t *testing.T) {
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
			mock.EXPECT().GetAll().Return(tt.mock, tt.mockErr)
			tu := &tagUsecase{repo: mock}
			got, err := tu.GetTags()

			if err != tt.wantErr {
				t.Errorf("tagUsecase.GetTags() error (wantErr %v, gotErr %v)", tt.wantErr, err)
			}

			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("tagUsecase.GetTags() mismatch (-want +got):\n%s", diff)
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
			tu := &tagUsecase{repo: mock}
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

func Test_tagUsecase_GetTagsByPostID(t *testing.T) {
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
			tu := &tagUsecase{repo: mock}
			got, err := tu.GetTagsByPostID(tt.postID)

			if err != tt.wantErr {
				t.Errorf("tagUsecase.GetTagByPostID() error (wantErr %v, gotErr %v)", tt.wantErr, err)
			}

			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("tagUsecase.GetTagByPostID() mismatch (-want +got):\n%s", diff)
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
			tu := &tagUsecase{repo: mock}
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
			tu := &tagUsecase{repo: mock}
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
