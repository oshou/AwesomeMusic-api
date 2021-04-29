package config

import (
	"fmt"

	"github.com/jinzhu/configor"
)

type IConfig interface {
	GetDSN() string
	GetDriver() string
}

type config struct {
	DB     DB
	Server Server
}

type DB struct {
	Host     string `required:"true" env:"DB_HOST"`
	Port     int    `required:"true" env:"DB_PORT"`
	DBName   string `required:"true" env:"DB_NAME"`
	SSLMode  string `required:"true" env:"DB_SSL_MODE"`
	User     string `required:"true" env:"DB_USER"`
	Password string `required:"true" env:"DB_PASSWORD"`
	Driver   string `required:"true" env:"DB_DRIVER" default:"postgres"`
}

type Server struct {
	Port             int `required:"true" env:"API_SERVER_PORT" default:"5432"`
	TimeoutSec       int `required:"true" env:"API_TIMEOUT_SECOND"`
	GzipLevel        int
	CorsMaxAgeSecond int
}

var _ IConfig = &config{}

func NewConfig(filePath string) (IConfig, error) {
	conf := &config{}
	if err := configor.Load(conf, filePath); err != nil {
		return nil, err
	}

	return conf, nil
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

func (c *config) GetDriver() string {
	return c.DB.Driver
}
