package database

import (
	"context"
	"fmt"
	"portal/pkg/config"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

var ctx = context.Background()

func InitRedis(c config.Config, l *zap.Logger) *redis.Client {

	adress := fmt.Sprintf("%s:%s", c.R_host, c.R_port)

	rDB := redis.NewClient(
		&redis.Options{
			Addr:     adress,
			Password: "",
			DB:       0,
		})

	_, err := rDB.Ping(ctx).Result()
	if err != nil {
		l.Error("Не удалось инициализировать Redis: " + err.Error())
		return nil
	}

	return rDB
}
