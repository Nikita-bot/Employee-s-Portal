package database

import (
	"fmt"
	"portal/pkg/config"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

func InitDB(log *zap.Logger, c config.Config) *sqlx.DB {
	str := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", c.Pg_host, c.Pg_port, c.Pg_user, c.Db_name, c.Pg_pass)
	db, err := sqlx.Open("postgres", str)
	if err != nil {
		log.Panic("Ошибка подключения к БД: " + err.Error())
	}
	err = db.Ping()
	if err != nil {
		log.Panic("Ошибка подключения к БД: " + err.Error())
	}

	return db
}
