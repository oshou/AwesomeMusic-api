//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=../mock/$GOPACKAGE/$GOFILE
// Package usecase is application layer package
package usecase

import (
	"reflect"
	"testing"

	"github.com/oshou/AwesomeMusic-api/api/domain/model"
	"github.com/oshou/AwesomeMusic-api/api/domain/repository"
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
			if got := NewSearchUsecase(tt.args.repo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSearchUsecase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_searchUsecase_ListPostsByTitle(t *testing.T) {
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
			got, err := tt.su.ListPostsByTitle(tt.args.q)
			if (err != nil) != tt.wantErr {
				t.Errorf("searchUsecase.ListPostsByTitle() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("searchUsecase.ListPostsByTitle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_searchUsecase_ListPostsByUserName(t *testing.T) {
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
			got, err := tt.su.ListPostsByUserName(tt.args.q)
			if (err != nil) != tt.wantErr {
				t.Errorf("searchUsecase.ListPostsByUserName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("searchUsecase.ListPostsByUserName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_searchUsecase_ListPostsByTagName(t *testing.T) {
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
			got, err := tt.su.ListPostsByTagName(tt.args.q)
			if (err != nil) != tt.wantErr {
				t.Errorf("searchUsecase.ListPostsByTagName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("searchUsecase.ListPostsByTagName() = %v, want %v", got, tt.want)
			}
		})
	}
}
