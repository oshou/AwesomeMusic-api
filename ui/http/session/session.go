package session

import (
	"encoding/hex"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/sessions"
	cache "github.com/patrickmn/go-cache"
	"github.com/pkg/errors"
	"go.uber.org/zap"

	"github.com/oshou/AwesomeMusic-api/log"
)

const (
	AuthTokenName    = "AUTH_TOKEN"
	CacheKeyTemplate = "CACHED_AUTH_TOKEN_%d_%s"
)

type Store struct {
	Store sessions.Store
	Cache *cache.Cache
}

func NewStore(secretKey string, opt *sessions.Options) (*Store, error) {
	key, err := hex.DecodeString(secretKey)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	s := sessions.NewCookieStore(key)
	if opt != nil {
		s.Options = opt
	}

	exp := int64(s.Options.MaxAge) * int64(time.Second)
	var (
		expiration      = time.Duration(exp)
		cleanupInterval = time.Duration(exp)
	)
	c := cache.New(
		expiration,
		cleanupInterval,
	)
	return &Store{
		Store: s,
		Cache: c,
	}, nil
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
