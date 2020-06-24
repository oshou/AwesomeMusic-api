// Package usecase is application layer package
package usecase

import (
	"testing"

	"github.com/oshou/AwesomeMusic-api/domain/model"
	"github.com/oshou/AwesomeMusic-api/domain/repository"

	"github.com/google/go-cmp/cmp"
)

func TestNewSearchUsecase(t *testing.T) {
	type args struct {
		repo repository.ISearchRepository
	}
	tests := []struct {
		name string
		args args
		want ISearchUsecase
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := NewSearchUsecase(tt.args.repo); !cmp.Equal(got, tt.want) {
				t.Errorf("NewSearchUsecase() = %v, want %v\ndiff=%v", got, tt.want, cmp.Diff(got, tt.want))
			}
		})
	}
}

func Test_searchUsecase_GetPostsByTitle(t *testing.T) {
	type args struct {
		q string
	}
	tests := []struct {
		name    string
		su      *searchUsecase
		args    args
		want    []*model.Post
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := tt.su.GetPostsByTitle(tt.args.q)
			if (err != nil) != tt.wantErr {
				t.Errorf("searchUsecase.GetPostsByTitle() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !cmp.Equal(got, tt.want) {
				t.Errorf("searchUsecase.GetPostsByTitle() = %v, want %v\ndiff=%v", got, tt.want, cmp.Diff(got, tt.want))
			}
		})
	}
}

func Test_searchUsecase_GetPostsByUserName(t *testing.T) {
	type args struct {
		q string
	}
	tests := []struct {
		name    string
		su      *searchUsecase
		args    args
		want    []*model.Post
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := tt.su.GetPostsByUserName(tt.args.q)
			if (err != nil) != tt.wantErr {
				t.Errorf("searchUsecase.GetPostsByUserName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !cmp.Equal(got, tt.want) {
				t.Errorf("searchUsecase.GetPostsByUserName() = %v, want %v\ndiff=%v", got, tt.want, cmp.Diff(got, tt.want))
			}
		})
	}
}

func Test_searchUsecase_GetPostsByTagName(t *testing.T) {
	type args struct {
		q string
	}
	tests := []struct {
		name    string
		su      *searchUsecase
		args    args
		want    []*model.Post
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := tt.su.GetPostsByTagName(tt.args.q)
			if (err != nil) != tt.wantErr {
				t.Errorf("searchUsecase.GetPostsByTagName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !cmp.Equal(got, tt.want) {
				t.Errorf("searchUsecase.GetPostsByTagName() = %v, want %v\ndiff=%v", got, tt.want, cmp.Diff(got, tt.want))
			}
		})
	}
}
