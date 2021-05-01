// Package usecase is application layer package
package usecase_test

import (
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
	"github.com/oshou/AwesomeMusic-api/api/domain/model"
	"github.com/oshou/AwesomeMusic-api/api/domain/repository"
	"github.com/oshou/AwesomeMusic-api/api/mock/mock_repository"
	"github.com/oshou/AwesomeMusic-api/api/usecase"
)

func TestNewUserUsecase(t *testing.T) {
	tests := []struct {
		name string
		repo repository.IUserRepository
		want usecase.IUserUsecase
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := usecase.NewUserUsecase(tt.repo); !cmp.Equal(got, tt.want) {
				t.Errorf("NewUserUsecase() = %v, want %v\ndiff=%v", got, tt.want, cmp.Diff(got, tt.want))
			}
		})
	}
}

func Test_userUsecase_ListUsers(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := []struct {
		name    string
		mock    []*model.User
		mockErr error
		want    []*model.User
		wantErr error
	}{
		{
			name: "success",
			mock: []*model.User{
				{ID: 1, Name: "Mike"},
				{ID: 2, Name: "Jane"},
				{ID: 3, Name: "John"},
			},
			mockErr: nil,
			want: []*model.User{
				{ID: 1, Name: "Mike"},
				{ID: 2, Name: "Jane"},
				{ID: 3, Name: "John1"},
			},
			wantErr: nil,
		},
		{
			name:    "no data",
			mock:    []*model.User{},
			mockErr: nil,
			want:    []*model.User{},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mock := mock_repository.NewMockIUserRepository(ctrl)
			mock.EXPECT().List().Return(tt.mock, tt.mockErr)
			uu := usecase.NewUserUsecase(mock)
			got, err := uu.ListUsers()

			if err != tt.wantErr {
				t.Errorf("userUsecase.ListUsers() error = %v, wantErr %v", err, tt.wantErr)
			}

			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("userUsecase.ListUsers() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func Test_userUsecase_GetUserByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := []struct {
		name    string
		userID  int
		mock    *model.User
		mockErr error
		want    *model.User
		wantErr error
	}{
		{
			name:    "success",
			userID:  1,
			mock:    &model.User{ID: 1, Name: "Mike"},
			mockErr: nil,
			want:    &model.User{ID: 1, Name: "Mike"},
			wantErr: nil,
		},
		{
			name:    "no data",
			userID:  0,
			mock:    &model.User{},
			mockErr: nil,
			want:    &model.User{},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mock := mock_repository.NewMockIUserRepository(ctrl)
			mock.EXPECT().GetByID(tt.userID).Return(tt.mock, tt.mockErr)
			uu := usecase.NewUserUsecase(mock)
			got, err := uu.GetUserByID(tt.userID)

			if err != tt.wantErr {
				t.Errorf("userUsecase.GetUserByID() error (wantErr %v, gotErr %v)", tt.wantErr, err)
			}

			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("userUsecase.GetUserByID() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func Test_userUsecase_AddUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := []struct {
		name     string
		username string
		mock     *model.User
		mockErr  error
		want     *model.User
		wantErr  error
	}{
		{
			name:     "success",
			username: "Mike",
			mock:     &model.User{ID: 1, Name: "Mike"},
			mockErr:  nil,
			want:     &model.User{ID: 1, Name: "Mike"},
			wantErr:  nil,
		},
		{
			name:     "no data",
			username: "Unknown",
			mock:     &model.User{},
			mockErr:  nil,
			want:     &model.User{},
			wantErr:  nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mock := mock_repository.NewMockIUserRepository(ctrl)
			mock.EXPECT().Add(tt.username).Return(tt.mock, tt.mockErr)
			uu := usecase.NewUserUsecase(mock)
			got, err := uu.AddUser(tt.username)

			if err != tt.wantErr {
				t.Errorf("userUsecase.AddUser() error (wantErr %v, gotErr %v)", tt.wantErr, err)
			}

			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("userUsecase.AddUser() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func Test_userUsecase_Authenticate(t *testing.T) {
	type fields struct {
		repo repository.IUserRepository
	}
	type args struct {
		username string
		password string
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
			uu := &userUsecase{
				repo: tt.fields.repo,
			}
			got, err := uu.Authenticate(tt.args.username, tt.args.password)
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
