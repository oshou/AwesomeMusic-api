package model

import (
	"time"
)

type Session struct {
	SessionID int
	Data      SessionData
	ExpiredAt time.Time
}

type SessionData struct {
	UserID int
}

func (s *Session) IsExpired(t time.Time) bool {
	return t.After(s.ExpiredAt)
}
