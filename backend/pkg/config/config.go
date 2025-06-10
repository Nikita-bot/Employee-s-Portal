package config

import "github.com/caarlos0/env/v9"

type Config struct {
	Port    string `env:"HTTP_BACKEND_PORT"`
	LogLvl  string `env:"LOGGER_LVL"`
	Pg_user string `env:"DB_USER"`
	Pg_pass string `env:"DB_PASSWORD"`
	Pg_host string `env:"DB_HOST"`
	Pg_port string `env:"DB_PORT"`
	Db_name string `env:"DB_NAME"`
	R_host  string `env:"REDIS_HOST"`
	R_port  string `env:"REDIS_PORT"`
	Token   string `env:"TOKEN_BOT"`
	Secret  string `env:"SECRET_KEY"`
}

func New() (Config, error) {
	var c Config
	err := env.ParseWithOptions(&c, env.Options{RequiredIfNoDef: true})
	if err != nil {
		return Config{}, err
	}

	return c, nil
}
