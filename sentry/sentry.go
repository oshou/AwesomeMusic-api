package sentry

import (
	"fmt"
	"time"

	"github.com/getsentry/sentry-go"
	sentrygo "github.com/getsentry/sentry-go"
	"github.com/pkg/errors"
)

func Init(dsn, env, name, version string) error {
	err := sentrygo.Init(sentry.ClientOptions{
		Dsn:         dsn,
		Environment: env,
		Release:     fmt.Sprintf("%s@%s", name, version),
	})
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func Recover() {
	err := recover()

	if err != nil {
		sentrygo.CurrentHub().Recover(err)
		sentrygo.Flush(2 * time.Second)
		panic(err)
	}
}
