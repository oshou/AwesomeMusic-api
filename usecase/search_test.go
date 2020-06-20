// Package usecase is application layer package
package usecase

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"

	"github.com/oshou/AwesomeMusic-api/domain/model"
	"github.com/oshou/AwesomeMusic-api/mock/mock_repository"
)

func Test_searchUsecase_GetPostsByTitle(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := []struct {
		name    string
		q       string
		mock    []*model.Post
		mockErr error
		want    []*model.Post
		wantErr error
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mock := mock_repository.NewMockISearchRepository(ctrl)
			mock.EXPECT().GetByTitle(tt.q).Return(tt.mock, tt.mockErr)
			su := &searchUsecase{repo: mock}
			got, err := su.GetPostsByTagName(tt.q)

			if err != tt.wantErr {
				t.Errorf("searchUsecase.GetPostsByTitle() error (wantErr %v, gotErr %v)", tt.wantErr, err)
			}

			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("searchUsecase.GetPostsByTitle() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func Test_searchUsecase_GetPostsByUserName(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := []struct {
		name    string
		q       string
		mock    []*model.Post
		mockErr error
		want    []*model.Post
		wantErr error
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mock := mock_repository.NewMockISearchRepository(ctrl)
			mock.EXPECT().GetByTitle(tt.q).Return(tt.mock, tt.mockErr)
			su := &searchUsecase{repo: mock}
			got, err := su.GetPostsByUserName(tt.q)

			if err != tt.wantErr {
				t.Errorf("searchUsecase.GetPostsUserName() error (wantErr %v, gotErr %v)", tt.wantErr, err)
			}

			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("searchUsecase.GetPostsUserName() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func Test_searchUsecase_GetPostsByTagName(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := []struct {
		name    string
		q       string
		mock    []*model.Post
		mockErr error
		want    []*model.Post
		wantErr error
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mock := mock_repository.NewMockISearchRepository(ctrl)
			mock.EXPECT().GetByTagName(tt.q).Return(tt.mock, tt.mockErr)
			su := &searchUsecase{repo: mock}
			got, err := su.GetPostsByTagName(tt.q)

			if err != tt.wantErr {
				t.Errorf("searchUsecase.GetPostsTagName() error (wantErr %v, gotErr %v)", tt.wantErr, err)
			}

			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("searchUsecase.GetPostsTagName() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
