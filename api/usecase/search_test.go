// Package usecase is application layer package
package usecase_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"

	"github.com/oshou/AwesomeMusic-api/api/domain/model"
	"github.com/oshou/AwesomeMusic-api/api/domain/repository"
	"github.com/oshou/AwesomeMusic-api/api/usecase"
	"github.com/oshou/AwesomeMusic-api/mock/mock_repository"
)

func TestNewSearchUsecase(t *testing.T) {
	tests := []struct {
		name string
		repo repository.ISearchRepository
		want usecase.ISearchUsecase
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := usecase.NewSearchUsecase(tt.repo); !cmp.Equal(got, tt.want) {
				t.Errorf("NewSearchUsecase() = %v, want %v\ndiff=%v", got, tt.want, cmp.Diff(got, tt.want))
			}
		})
	}
}

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
			su := usecase.NewSearchUsecase(mock)
			got, err := su.GetPostsByTitle(tt.q)

			if err != tt.wantErr {
				t.Errorf("seearchUsecase.AddPost() error (wantErr %v, gotErr %v)", tt.wantErr, err)
			}

			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("searchUsecase().AddPost() mismatch (-want +got):\n%s", diff)
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
			mock.EXPECT().GetByUserName(tt.q).Return(tt.mock, tt.mockErr)
			su := usecase.NewSearchUsecase(mock)
			got, err := su.GetPostsByUserName(tt.q)

			if err != tt.wantErr {
				t.Errorf("searchUsecase.GetPostByUserName() error (wantErr %v, gotErr %v)", tt.wantErr, err)
			}

			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("searchUsecase.GetPostByUserName() mismatch (-want +got):\n%s", diff)
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
			su := usecase.NewSearchUsecase(mock)
			got, err := su.GetPostsByTagName(tt.q)

			if err != tt.wantErr {
				t.Errorf("searchUsecase.GetPostByTagName() error (wantErr %v, gotErr %v)", tt.wantErr, err)
			}

			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("searchUsecase.GetPostByTagName() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
