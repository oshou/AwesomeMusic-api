package config

import (
	"fmt"

	"github.com/jinzhu/configor"
)

type IConfig interface {
	GetDriver() string
	GetDSN() string
	GetSentryDSN() string
}

type config struct {
	DB     db
	Server server
	Sentry sentry
}

type db struct {
	Driver   string `required:"true" env:"DB_DRIVER" default:"postgres"`
	Host     string `required:"true" env:"DB_HOST"`
	Port     int    `required:"true" env:"DB_PORT"`
	DBName   string `required:"true" env:"DB_NAME"`
	SSLMode  string `required:"true" env:"DB_SSL_MODE"`
	User     string `required:"true" env:"DB_USER"`
	Password string `required:"true" env:"DB_PASSWORD"`
}

type server struct {
	Port             int `required:"true" env:"API_SERVER_PORT" default:"5432"`
	TimeoutSec       int `required:"true" env:"API_TIMEOUT_SECOND"`
	GzipLevel        int
	CorsMaxAgeSecond int
}

type sentry struct {
	DSN string `required:"true" env:"SENTRY_DSN"`
}

var _ IConfig = &config{}

func NewConfig(filePath string) (IConfig, error) {
	conf := &config{}
	if err := configor.Load(conf, filePath); err != nil {
		return nil, err
	}

	return conf, nil
}

func (c *config) GetDriver() string {
	return c.DB.Driver
}

func (c *config) GetDSN() string {
	var dsn string
	dsn = fmt.Sprintf(
		"host=%v port=%v user=%v password=%v dbname=%v sslmode=%v",
		c.DB.Host,
		c.DB.Port,
		c.DB.User,
		c.DB.Password,
		c.DB.DBName,
		c.DB.SSLMode,
	)
	return dsn
}

func (c *config) GetSentryDSN() string {
	return c.Sentry.DSN
}
