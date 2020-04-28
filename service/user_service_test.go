// Package service is application layer package
package service

import (
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/oshou/AwesomeMusic-api/domain/model"
	"github.com/oshou/AwesomeMusic-api/domain/repository"
	"github.com/oshou/AwesomeMusic-api/mock/mock_repository"
)

func TestUserService_GetUsers(t *testing.T) {
	ctrl := gomock.NewController(t)
	ctrl.Finish()

	tests := []struct {
		users []*model.User
		err   error
	}{}

	for _, tt := range tests {
		mockRepo := mock_repository.NewMockIUserRepository(ctrl)
		mockRepo.EXPECT().GetAll().Return(tt.users, tt.err)
		us := NewUserService(mockRepo)
		users, err := us.GetUsers()

		if err != nil {
			t.Errorf("[Failed] %v", err)
		}

		if !reflect.DeepEqual(tt.users, users) {
			t.Errorf("[Failed] want:%v, got:%v", tt.users, users)
		}
	}
}

func TestUserService_GetUserByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	ctrl.Finish()

	tests := []struct {
		userID int
		user   *model.User
		err    error
	}{}

	for _, tt := range tests {
		mockRepo := mock_repository.NewMockIUserRepository(ctrl)
		mockRepo.EXPECT().GetByID(tt.userID).Return(tt.user, tt.err)
		us := NewUserService(mockRepo)
		user, err := us.GetUserByID(tt.userID)

		if err != nil {
			t.Errorf("[Failed] %v", err)
		}

		if !reflect.DeepEqual(tt.user, user) {
			t.Errorf("[Failed] want:%v, got:%v", tt.user, user)
		}
	}
}

func TestUserService_AddUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	ctrl.Finish()

	tests := []struct {
		userName string
		user     *model.User
		err      error
	}{}

	for _, tt := range tests {
		mockRepo := mock_repository.NewMockIUserRepository(ctrl)
		mockRepo.EXPECT().Add(tt.userName).Return(tt.user, tt.err)
		us := NewUserService(mockRepo)
		user, err := us.AddUser(tt.userName)

		if err != nil {
			t.Errorf("[Failed] %v", err)
		}

		if !reflect.DeepEqual(tt.user, user) {
			t.Errorf("[Failed] want:%v, got:%v", tt.user, user)
		}
	}
}

func TestNewUserService(t *testing.T) {
	type args struct {
		repo repository.IUserRepository
	}
	tests := []struct {
		name string
		args args
		want IUserService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUserService(tt.args.repo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userService_GetUsers(t *testing.T) {
	type fields struct {
		repo repository.IUserRepository
	}
	tests := []struct {
		name    string
		fields  fields
		want    []*model.User
		wantErr bool
	}{
		{"GetUser", fields{repo: repository.IUserRepository}, nil, false},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uu := &userService{
				repo: tt.fields.repo,
			}
			got, err := uu.GetUsers()
			if (err != nil) != tt.wantErr {
				t.Errorf("userService.GetUsers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userService.GetUsers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userService_GetUserByID(t *testing.T) {
	type fields struct {
		repo repository.IUserRepository
	}
	type args struct {
		userID int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uu := &userService{
				repo: tt.fields.repo,
			}
			got, err := uu.GetUserByID(tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("userService.GetUserByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userService.GetUserByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userService_AddUser(t *testing.T) {
	type fields struct {
		repo repository.IUserRepository
	}
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uu := &userService{
				repo: tt.fields.repo,
			}
			got, err := uu.AddUser(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("userService.AddUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userService.AddUser() = %v, want %v", got, tt.want)
			}
		})
	}
}
