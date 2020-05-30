// Package service is application layer package
package service

import (
	"reflect"
	"testing"

	"github.com/oshou/AwesomeMusic-api/domain/model"
	"github.com/oshou/AwesomeMusic-api/domain/repository"
)

func TestNewSearchService(t *testing.T) {
	type args struct {
		repo repository.ISearchRepository
	}
	tests := []struct {
		name string
		args args
		want ISearchService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSearchService(tt.args.repo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSearchService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_searchService_GetPostsByTitle(t *testing.T) {
	type args struct {
		q string
	}
	tests := []struct {
		name    string
		su      *searchService
		args    args
		want    []*model.Post
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.su.GetPostsByTitle(tt.args.q)
			if (err != nil) != tt.wantErr {
				t.Errorf("searchService.GetPostsByTitle() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("searchService.GetPostsByTitle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_searchService_GetPostsByUserName(t *testing.T) {
	type args struct {
		q string
	}
	tests := []struct {
		name    string
		su      *searchService
		args    args
		want    []*model.Post
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.su.GetPostsByUserName(tt.args.q)
			if (err != nil) != tt.wantErr {
				t.Errorf("searchService.GetPostsByUserName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("searchService.GetPostsByUserName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_searchService_GetPostsByTagName(t *testing.T) {
	type args struct {
		q string
	}
	tests := []struct {
		name    string
		su      *searchService
		args    args
		want    []*model.Post
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.su.GetPostsByTagName(tt.args.q)
			if (err != nil) != tt.wantErr {
				t.Errorf("searchService.GetPostsByTagName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("searchService.GetPostsByTagName() = %v, want %v", got, tt.want)
			}
		})
	}
}
