// Package usecase is application layer package
package usecase

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
	"github.com/oshou/AwesomeMusic-api/domain/model"
	"github.com/oshou/AwesomeMusic-api/mock/mock_repository"
)

func Test_userUsecase_GetUsers(t *testing.T) {
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
				{ID: 3, Name: "John"},
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
			mock.EXPECT().GetAll().Return(tt.mock, tt.mockErr)
			uu := &userUsecase{repo: mock}
			got, err := uu.GetUsers()

			if err != tt.wantErr {
				t.Errorf("userUsecase.GetUsers() error = %v, wantErr %v", err, tt.wantErr)
			}

			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("userUsecase.GetUsers() mismatch (-want +got):\n%s", diff)
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
			uu := &userUsecase{repo: mock}
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
			uu := &userUsecase{repo: mock}
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
