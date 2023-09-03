package app

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Env         string `envconfig:"RUN_ENV" required:"true"`
	DatabaseUrl string `envconfig:"DATABASE_URL" default:"host=localhost port=5432 user=postgres dbname=postgres sslmode=disable"`
}

func InitConfig() (Config, error) {
	var c Config
	err := envconfig.Process("app", &c)
	if err != nil {
		return Config{}, err
	} else {
		return c, nil
	}
}
