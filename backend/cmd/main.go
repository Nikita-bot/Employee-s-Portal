package main

import (
	"portal/pkg/config"
	"portal/pkg/database"

	"portal/pkg/logger"

	"github.com/pressly/goose"
)

func main() {
	config, err := config.New()
	if err != nil {
		panic(err)
	}
	logger := logger.InitLogger(config)
	logger.Debug("Логер готов")

	logger.Info("Подключение к БД")
	db := database.InitDB(logger, config)
	defer db.Close()

	if err := goose.SetDialect("postgres"); err != nil {
		logger.Error(err.Error())
	}

	err = goose.Up(db.DB, "app/migration")
	if err != nil {
		logger.Error(err.Error())
		panic("Не удалось загрузить файлы миграции")
	}
	logger.Debug("БД готова")
}
