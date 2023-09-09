package app

import (
	"fmt"
	"github.com/kelseyhightower/envconfig"
)

type RunEnv string

const (
	Development = RunEnv("Development")
	Staging     = RunEnv("Staging")
	Production  = RunEnv("Production")
)

func (m *RunEnv) Decode(value string) error {
	switch {
	//FIXME: duplicated value definition
	case value == "Production":
		*m = Production
	case value == "Staging":
		*m = Production
	case value == "Development":
		*m = Development
	default:
		return fmt.Errorf("'%s' is an unrecognized env", value)
	}
	return nil
}

type AppConfig struct {
	Env         RunEnv `envconfig:"RUN_ENV" required:"true"`
	DatabaseUrl string `envconfig:"DATABASE_URL" default:"postgres://postgres:postgres@127.0.0.1:/postgres?sslmode=disable"`
}

var c AppConfig

func InitAppConfig() error {
	err := envconfig.Process("app", &c)
	if err != nil {
		return err
	}
	return nil
}

func GetAppConfig() AppConfig {
	return c
}
