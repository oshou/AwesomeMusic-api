package session_test

import (
	"testing"

	"github.com/gorilla/sessions"
)

const sskey = "aaaa"

func TestNewStore(t *testing.T) {
	opt := &sessions.Options{
		Path:     "/",
		Domain:   "example.com",
		MaxAge:   60 * 60 * 24,
		Secure:   true,
		HttpOnly: true,
	}
	tests := []struct {
		name    string
		key     string
		opt     *sessions.Options
		want    *sessions.Store
		wantErr error
	}{
		{
			name: "Success",
			key:  sskey,
			opt:  opt,
			want: &sessions.Store{
				Store: &sessions.CookieStore{
					Options: opt,
				},
			},
			wantErr: nil,
		},
	}
}
