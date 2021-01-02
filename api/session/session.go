package session

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"github.com/pkg/errors"
	"go.uber.org/zap"

	"github.com/oshou/AwesomeMusic-api/log"
)

const (
	AuthTokenName    = "AUTH_TOKEN"
	CacheKeyTemplate = "CACHED_AUTH_TOKEN_%d_%s"
)

// PGStore implements gorilla/sessions Store interface
// https://github.com/gorilla/sessions/blob/master/store.go#L22
type PGStore struct {
	Codecs  []securecookie.Codec
	Options *sessions.Options
	Path    string
	DbPool  *sql.DB
}

type PGSession struct {
	ID         int64
	Key        string
	Data       string
	CreatedOn  time.Time
	ModifiedOn time.Time
	ExpiresOn  time.Time
}

func NewPGStore(dbURL string, keyPairs ...[]byte) (*PGStore, error) {
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		return nil, err
	}
	return NewPGStoreFromPool(db, keyPairs...)
}

func NewPGStoreFromPool(db *sql.DB, keyPairs ...[]byte) (*PGStore, error) {
	dbStore := &PGStore{
		Codecs: []securecookie.CodecsFromPairs(keyPairs),
		Options: &sessions.Options{
			Path:   "/",
			MaxAge: 86400 * 30,
		},
		DbPool: db,
	}

	err := dbStore.createSessionsTable()
	if err != nil {
		return nil, err
	}

	return dbStore, nil
}

func (s *Store) Save(r *http.Request, w http.ResponseWriter, userID int) error {
	ss, err := s.Store.New(r, AuthTokenName)
	if err != nil {
		log.Logger.Error("failed to new session store", zap.Error(err))
		return errors.WithStack(err)
	}

	ss.Values["uid"] = userID
	ss.Values["authenticated"] = true
	if err := ss.Save(r, w); err != nil {
		log.Logger.Error("failed to save session", zap.Error(err))
		return errors.WithStack(err)
	}

	cacheKey := fmt.Sprintf(CacheKeyTemplate, userID)
	s.Cache.SetDefault(cacheKey, ss)
	return nil
}

func (s *Store) Clear(r *http.Request, w http.ResponseWriter) (int, error) {
	ss, err := s.Store.Get(r, AuthTokenName)
	if err != nil {
		return 0, nil
	}

	uid, _ := ss.Values["uid"].(int)
	if !ss.IsNew {
		ss.Options.MaxAge = -1
		if err := ss.Save(r, w); err != nil {
			log.Logger.Error("failed to save session", zap.Error(err))
			return 0, errors.WithStack(err)
		}
		cacheKey := fmt.Sprintf(CacheKeyTemplate, uid)
		s.Cache.Delete(cacheKey)
	}
	return uid, nil
}
