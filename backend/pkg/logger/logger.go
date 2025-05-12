package logger

import (
	"fmt"
	"os"
	"portal/pkg/config"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func InitLogger(conf config.Config) *zap.Logger {
	var logLevel zapcore.Level
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.TimeEncoderOfLayout("15:04:05 02-01-2006"),
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	if conf.LogLvl == "DEBUG" {
		logLevel = zapcore.DebugLevel
	} else if conf.LogLvl == "INFO" {
		logLevel = zapcore.InfoLevel
	} else {
		logLevel = zapcore.InfoLevel // уровень по умолчанию
	}

	file, err := os.OpenFile("app/log/log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Could Not Open Log File : " + err.Error())
	}

	// Настраиваем мультиплексор вывода
	core := zapcore.NewTee(
		zapcore.NewCore(
			zapcore.NewConsoleEncoder(encoderConfig),
			zapcore.AddSync(file),
			logLevel,
		),
		zapcore.NewCore(
			zapcore.NewConsoleEncoder(encoderConfig),
			zapcore.AddSync(os.Stdout),
			logLevel,
		),
	)

	// Создаем логгер
	logger := zap.New(core, zap.AddCaller())
	defer logger.Sync()

	return logger
}
