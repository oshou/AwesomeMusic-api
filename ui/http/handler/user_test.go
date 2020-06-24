package handler_test

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
	"github.com/oshou/AwesomeMusic-api/domain/model"
	"github.com/oshou/AwesomeMusic-api/mock/mock_usecase"
	"github.com/oshou/AwesomeMusic-api/ui/http/handler"
	"github.com/oshou/AwesomeMusic-api/usecase"
)

func Test_userHandler_GetUsers(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := []struct {
		name     string
		mock     []*model.User
		mockErr  error
		wantCode int
		wantBody string
	}{
		// TODO: Add test pattern
		{
			name:     "success",
			mock:     []*model.User{},
			mockErr:  nil,
			wantCode: http.StatusOK,
			wantBody: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mock := mock_usecase.NewMockIUserUsecase(ctrl)
			mock.EXPECT().GetUsers().Return(tt.mock, tt.mockErr)

			r := httptest.NewRequest(http.MethodGet, "/v1/users", nil)
			rr := httptest.NewRecorder()
			h := handler.NewUserHandler(mock)

			h.GetUsers(rr, r)

			if diff := cmp.Diff(tt.wantCode, rr.Code); diff != "" {
				t.Errorf("userHandler.GetUsers() mismatch status code (-want +got):\n%s", diff)
			}

			if diff := cmp.Diff(tt.wantBody, rr.Body); diff != "" {
				t.Errorf("userHandler.GetUsers() mismatch body (-want +got):\n%s", diff)
			}
		})
	}
}

func Test_userHandler_AddUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := []struct {
		name     string
		mock     *model.User
		mockErr  error
		wantCode int
		wantBody string
	}{
		// TODO: Add test pattern
		{
			name:     "success",
			mock:     &model.User{},
			mockErr:  nil,
			wantCode: http.StatusOK,
			wantBody: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			mock := mock_usecase.NewMockIUserUsecase(ctrl)
			mock.EXPECT().AddUser(tt.name).Return(tt.mock, tt.mockErr)

			r := httptest.NewRequest(http.MethodPost, "/v1/users", nil)
			rr := httptest.NewRecorder()
			h := handler.NewUserHandler(mock)

			h.AddUser(rr, r)

			if diff := cmp.Diff(tt.wantCode, rr.Code); diff != "" {
				t.Errorf("userHandler.GetUsers() mismatch status code (-want +got):\n%s", diff)
			}

			if diff := cmp.Diff(tt.wantBody, rr.Body); diff != "" {
				t.Errorf("userHandler.GetUsers() mismatch body (-want +got):\n%s", diff)
			}
		})
	}
}

func Test_userHandler_GetUserByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := []struct {
		name     string
		userID   int
		mock     *model.User
		wantCode int
		wantBody string
	}{
		name:     "success",
		userID:   1,
		mock:     &model.User{},
		wantCode: http.StatusOK,
		wantBody: "",
	}
	t.Run(tt.name, func(t *testing.T) {
		t.Parallel()
		mock := mock_usecase.NewMockIUserUsecase(ctrl)
		mock.EXPECT().GetUserByID(tt.userID).Return(tt.wantCode, tt.wantBody)

		r := httptest.NewRequest(http.MethodGet, "/v1/users/1", nil)
		rr := httptest.NewRecorder()
		h := handler.NewUserHandler(mock)

		h.GetUserByID(tt.userID)

		if diff := cmp.Diff(tt.wantCode, rr.Code); diff != "" {
			t.Errorf("userHandler.GetUsers() mismatch status code (-want +got):\n%s", diff)
		}

		if diff := cmp.Diff(tt.wantBody, rr.Body); diff != "" {
			t.Errorf("userHandler.GetUsers() mismatch body (-want +got):\n%s", diff)
		}

	})

}

func TestNewUserHandler(t *testing.T) {
	type args struct {
		usecase usecase.IUserUsecase
	}
	tests := []struct {
		name string
		args args
		want IUserHandler
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUserHandler(tt.args.usecase); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}
