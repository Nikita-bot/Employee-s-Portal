package config

import "github.com/caarlos0/env/v9"

type Config struct {
	Port    string `env:"HTTP_BACKEND_PORT" envDefault:"8080"`
	LogLvl  string `env:"LOGGER_LVL" envDefault:""`
	Pg_user string `env:"DB_USER" envDefault:"admin"`
	Pg_pass string `env:"DB_PASSWORD" envDefault:"password"`
	Pg_host string `env:"DB_HOST" envDefault:"median_map_postgres"`
	Pg_port string `env:"DB_PORT" envDefault:"5432"`
	Db_name string `env:"DB_NAME" envDefault:"MedianMap"`
	Token   string `env:"TOKEN_BOT" envDefault:""`
}

func New() (Config, error) {
	var c Config

	err := env.ParseWithOptions(&c, env.Options{RequiredIfNoDef: true})
	if err != nil {
		return Config{}, err
	}

	return c, nil
}
