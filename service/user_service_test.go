// Package service is application layer package
package service

import (
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/oshou/AwesomeMusic-api/domain/repository"
	"github.com/oshou/AwesomeMusic-api/mock/mock_repository"
)

func TestNewUserService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := mock_repository.NewMockIUserRepository(ctrl)

	type args struct {
		repo repository.IUserRepository
	}
	tests := []struct {
		name string
		args args
		want IUserService
	}{
		{
			name: "aaa-test",
			args: args{repo: mockRepo},
			want: &userService{repo: mockRepo},
		},
		{
			name: "aaa-test",
			args: args{repo: mockRepo},
			want: &userService{repo: mockRepo},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUserService(tt.args.repo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserService() = %v, want %v", got, tt.want)
			}
		})
	}
}

///// func Test_userService_GetUsers(t *testing.T) {
///// 	type fields struct {
///// 		repo repository.IUserRepository
///// 	}
///// 	mockRepo := &repository.UserRepositoryMock{}
///// 	tests := []struct {
///// 		name    string
///// 		fields  fields
///// 		want    []*model.User
///// 		wantErr bool
///// 	}{
///// 		{
///// 			"test-aaa",
///// 			fields{repo: mockRepo},
///// 			[]*model.User{},
///// 			false,
///// 		},
///// 	}
///// 	for _, tt := range tests {
///// 		mockRepo.GetAllMock = func() ([]*model.User, error) {
///// 			return tt.want, nil
///// 		}
///// 		tt.fields.repo, _ = mockRepo.GetAllMock()
///// 		t.Run(tt.name, func(t *testing.T) {
///// 			uu := &userService{
///// 				repo: tt.fields.repo,
///// 			}
///// 			got, err := uu.GetUsers()
///// 			if (err != nil) != tt.wantErr {
///// 				t.Errorf("userService.GetUsers() error = %v, wantErr %v", err, tt.wantErr)
///// 				return
///// 			}
///// 			if !reflect.DeepEqual(got, tt.want) {
///// 				t.Errorf("userService.GetUsers() = %v, want %v", got, tt.want)
///// 			}
///// 		})
///// 	}
///// }
/////
///// func Test_userService_GetUserByID(t *testing.T) {
///// 	type fields struct {
///// 		repo repository.IUserRepository
///// 	}
///// 	type args struct {
///// 		userID int
///// 	}
///// 	tests := []struct {
///// 		name    string
///// 		fields  fields
///// 		args    args
///// 		want    *model.User
///// 		wantErr bool
///// 	}{
///// 		// TODO: Add test cases.
///// 	}
///// 	for _, tt := range tests {
///// 		t.Run(tt.name, func(t *testing.T) {
///// 			uu := &userService{
///// 				repo: tt.fields.repo,
///// 			}
///// 			got, err := uu.GetUserByID(tt.args.userID)
///// 			if (err != nil) != tt.wantErr {
///// 				t.Errorf("userService.GetUserByID() error = %v, wantErr %v", err, tt.wantErr)
///// 				return
///// 			}
///// 			if !reflect.DeepEqual(got, tt.want) {
///// 				t.Errorf("userService.GetUserByID() = %v, want %v", got, tt.want)
///// 			}
///// 		})
///// 	}
///// }
/////
///// func Test_userService_AddUser(t *testing.T) {
///// 	type fields struct {
///// 		repo repository.IUserRepository
///// 	}
///// 	type args struct {
///// 		name string
///// 	}
///// 	tests := []struct {
///// 		name    string
///// 		fields  fields
///// 		args    args
///// 		want    *model.User
///// 		wantErr bool
///// 	}{
///// 		// TODO: Add test cases.
///// 	}
///// 	for _, tt := range tests {
///// 		t.Run(tt.name, func(t *testing.T) {
///// 			uu := &userService{
///// 				repo: tt.fields.repo,
///// 			}
///// 			got, err := uu.AddUser(tt.args.name)
///// 			if (err != nil) != tt.wantErr {
///// 				t.Errorf("userService.AddUser() error = %v, wantErr %v", err, tt.wantErr)
///// 				return
///// 			}
///// 			if !reflect.DeepEqual(got, tt.want) {
///// 				t.Errorf("userService.AddUser() = %v, want %v", got, tt.want)
///// 			}
///// 		})
///// 	}
///// }
