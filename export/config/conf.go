package config

type Config struct {
	Pg_user       string `env:"DB_USER"`
	Pg_pass       string `env:"DB_PASSWORD"`
	Pg_host       string `env:"DB_HOST"`
	Pg_port       string `env:"DB_PORT"`
	Db_name       string `env:"DB_NAME"`
	FIRST_API_URL string
	API_URL       string
}

func New() *Config {
	var c Config

	c.Pg_user = "PORTALDATA"
	c.Pg_pass = "GAUSKKBSMP_3_DATA-base"
	c.Db_name = "EMPLOYERSPORTALDATABASE"
	c.Pg_port = "5432"
	c.Pg_host = "portal_database"
	c.FIRST_API_URL = "http://titan/testExport/hs/employees/getEmployees/00010101/true"
	c.API_URL = "http://titan/testExport/hs/employees/getEmployees/%s/false"
	return &c
}
