//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=../mock/$GOPACKAGE/$GOFILE
// Package usecase is application layer package
package usecase

import (
	"reflect"
	"testing"

	"github.com/oshou/AwesomeMusic-api/api/domain/model"
	"github.com/oshou/AwesomeMusic-api/api/domain/repository"
)

func TestNewUserUsecase(t *testing.T) {
	type args struct {
		repo repository.IUserRepository
	}
	tests := []struct {
		name string
		args args
		want IUserUsecase
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUserUsecase(tt.args.repo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserUsecase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userUsecase_Authenticate(t *testing.T) {
	type args struct {
		username string
		password string
	}
	tests := []struct {
		name    string
		uu      *userUsecase
		args    args
		want    *model.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.uu.Authenticate(tt.args.username, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("userUsecase.Authenticate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userUsecase.Authenticate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userUsecase_ListUsers(t *testing.T) {
	tests := []struct {
		name    string
		uu      *userUsecase
		want    []*model.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.uu.ListUsers()
			if (err != nil) != tt.wantErr {
				t.Errorf("userUsecase.ListUsers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userUsecase.ListUsers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userUsecase_GetUserByID(t *testing.T) {
	type args struct {
		userID int
	}
	tests := []struct {
		name    string
		uu      *userUsecase
		args    args
		want    *model.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.uu.GetUserByID(tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("userUsecase.GetUserByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userUsecase.GetUserByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userUsecase_AddUser(t *testing.T) {
	type args struct {
		name     string
		password string
	}
	tests := []struct {
		name    string
		uu      *userUsecase
		args    args
		want    *model.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.uu.AddUser(tt.args.name, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("userUsecase.AddUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userUsecase.AddUser() = %v, want %v", got, tt.want)
			}
		})
	}
}
