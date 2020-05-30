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
			name: "new",
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

func Test_userService_GetUsers(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := []struct {
		name    string
		prepare func(*mock_repository.MockIUserRepository)
		want    []*model.User
		wantErr bool
	}{
		{
			name: "no-data",
			prepare: func(mockRepo *mock_repository.MockIUserRepository) {
				mockRepo.EXPECT().GetAll().Return([]*model.User{}, nil)
			},
			want:    []*model.User{},
			wantErr: false,
		},
		{
			name: "3-data",
			prepare: func(mockRepo *mock_repository.MockIUserRepository) {
				mockRepo.EXPECT().GetAll().Return([]*model.User{
					{ID: 1, Name: "Mike"},
					{ID: 2, Name: "Jane"},
					{ID: 3, Name: "John"},
				}, nil)
			},
			want: []*model.User{
				{ID: 1, Name: "Mike"},
				{ID: 2, Name: "Jane"},
				{ID: 3, Name: "John"},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := mock_repository.NewMockIUserRepository(ctrl)
			tt.prepare(mockRepo)
			uu := &userService{repo: mockRepo}
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
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type args struct {
		userID int
	}
	tests := []struct {
		name    string
		args    args
		prepare func(*mock_repository.MockIUserRepository)
		want    *model.User
		wantErr bool
	}{
		{
			name: "no-data",
			args: args{userID: 0},
			prepare: func(mockRepo *mock_repository.MockIUserRepository) {
				mockRepo.EXPECT().GetByID(0).Return(&model.User{}, nil)
			},
			want:    &model.User{},
			wantErr: false,
		},
		{
			name: "1-data",
			args: args{userID: 1},
			prepare: func(mockRepo *mock_repository.MockIUserRepository) {
				mockRepo.EXPECT().GetByID(1).Return(&model.User{ID: 1, Name: "Mike"}, nil)
			},
			want:    &model.User{ID: 1, Name: "Mike"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := mock_repository.NewMockIUserRepository(ctrl)
			tt.prepare(mockRepo)
			uu := &userService{repo: mockRepo}
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
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type args struct {
		name string
	}
	tests := []struct {
		name    string
		args    args
		prepare func(*mock_repository.MockIUserRepository)
		want    *model.User
		wantErr bool
	}{
		{
			name: "no-data",
			args: args{name: "unknown"},
			prepare: func(mockRepo *mock_repository.MockIUserRepository) {
				mockRepo.EXPECT().Add("unknown").Return(&model.User{}, nil)
			},
			want:    &model.User{},
			wantErr: false,
		},
		{
			name: "1-data",
			args: args{name: "Mike"},
			prepare: func(mockRepo *mock_repository.MockIUserRepository) {
				mockRepo.EXPECT().Add("Mike").Return(&model.User{ID: 1, Name: "Mike"}, nil)
			},
			want:    &model.User{ID: 1, Name: "Mike"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := mock_repository.NewMockIUserRepository(ctrl)
			tt.prepare(mockRepo)
			uu := &userService{repo: mockRepo}
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
