package app

import (
  "fmt"
	"github.com/kelseyhightower/envconfig"
)

type RunEnv int

const (
  Development RunEnv = iota
  Staging
  Production
)

func (m *RunEnv) Decode(value string) error {
    switch {
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

type Config struct {
	Env         RunEnv `envconfig:"RUN_ENV" required:"true"`
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
