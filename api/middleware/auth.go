package middleware

import (
	"context"
	"net/http"

	"github.com/gorilla/sessions"

	"github.com/oshou/AwesomeMusic-api/api/usecase"
)

type TokenKey string
type CtxKey string

const (
	AuthTokenKey TokenKey = "AUTH-TOKEN"
	UserCtxKey   CtxKey   = "user"
)

func Auth(store *sessions.Store, usecase usecase.IUserUsecase) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			// リクエストヘッダのAuthTokenを取得
			//authToken := r.Header.Get(AuthTokenKey)

			// store
			ctx := context.WithValue(r.Context(), UserCtxKey, usecase)
			next.ServeHTTP(w, r.WithContext(ctx))
		}
		return http.HandlerFunc(fn)
	}
}
