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
	conf, err := config.New()
	if err != nil {
		panic(err)
	}
	logger := logger.InitLogger(conf)
	logger.Debug("Логер готов")

	logger.Info("Подключение к БД")
	db := database.InitDB(logger, conf)
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

	err = config.SeedInitialData(db)
	if err != nil {
		logger.Error(err.Error())
		panic("Не удалось загрузить начальные данные")
	}

	redis := database.InitRedis(conf, logger)

	server, err := server.InitServer(conf)
	if err != nil {
		logger.Error(err.Error())
		panic("Не удалось инициализировать сервер")
	}

	repo := repository.NewRepository(db, logger, redis)

	service := service.NewService(logger, repo)

	handler := handlers.NewHandler(logger, service, server)

	err = handler.Handle(conf.Port)
	if err != http.ErrServerClosed {
		logger.Error(err.Error())
		panic("Не удалось запустить сервер")
	}
}
