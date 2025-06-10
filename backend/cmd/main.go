package main

import (
	"net/http"
	"portal/internal/handlers"
	"portal/internal/repository"
	"portal/internal/service"
	"portal/pkg/config"
	"portal/pkg/database"
	"portal/pkg/server"

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

	err = goose.Up(db.DB, "/app/migrations")
	if err != nil {
		logger.Error(err.Error())
		panic("Не удалось загрузить файлы миграции")
	}
	logger.Debug("БД готова")

	redis := database.InitRedis(config, logger)

	server, err := server.InitServer(config)
	if err != nil {
		logger.Error(err.Error())
		panic("Не удалось инициализировать сервер")
	}

	repo := repository.NewRepository(db, logger, redis)

	service := service.NewService(logger, repo)

	handler := handlers.NewHandler(logger, service, server)

	err = handler.Handle(config.Port)
	if err != http.ErrServerClosed {
		logger.Error(err.Error())
		panic("Не удалось запустить сервер")
	}
}
