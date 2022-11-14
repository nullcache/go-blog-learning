package global

import (
	"github.com/nullcache/go-blog-learning/pkg/logger"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
)

var (
	Logger *logger.Logger
)

func SetupLogger() error {
	Logger = logger.NewLogger(&lumberjack.Logger{
		Filename:  LogSetting.LogSavePath + "/" + LogSetting.LogFileName + LogSetting.LogFileExt,
		MaxSize:   LogSetting.LogMaxSize,
		MaxAge:    LogSetting.LogMaxAge,
		LocalTime: LogSetting.LogLocalTime,
	}, "", log.LstdFlags).WithCaller(2)
	return nil
}
