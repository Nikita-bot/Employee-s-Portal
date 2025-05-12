package config

import "github.com/caarlos0/env/v9"

type Config struct {
	Port    string `env:"PORT_BACK" envDefault:"8080"`
	R_port  string `env:"REDIS_PORT" envDefault:"6379"`
	R_host  string `env:"REDIS_HOST" envDefault:"geomap_test_redis"`
	LogLvl  string `env:"LOGGER_LVL" envDefault:""`
	Pg_user string `env:"PG_USER" envDefault:"admin"`
	Pg_pass string `env:"PG_PASS" envDefault:"password"`
	Pg_host string `env:"PG_HOST" envDefault:"median_map_postgres"`
	Pg_port string `env:"PG_PORT" envDefault:"5432"`
	Db_name string `env:"PG_DB_NAME" envDefault:"MedianMap"`
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
