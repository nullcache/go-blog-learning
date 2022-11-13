package global

import (
	"github.com/nullcache/go-blog-learning/pkg/logger"
	"github.com/nullcache/go-blog-learning/pkg/setting"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
)

var (
	ServerSetting   *setting.ServerSetting
	AppSetting      *setting.AppSetting
	LogSetting      *setting.LogSetting
	DatabaseSetting *setting.DatabaseSetting
	Logger          *logger.Logger
)

func SetupSetting() error {
	s, err := setting.NewSetting()
	if err != nil {
		return err
	}
	if err = s.ReadSection("Server", &ServerSetting); err != nil {
		return err
	}
	if err = s.ReadSection("App", &AppSetting); err != nil {
		return err
	}
	if err = s.ReadSection("Log", &LogSetting); err != nil {
		return err
	}
	if err = s.ReadSection("Database", &DatabaseSetting); err != nil {
		return err
	}
	return nil
}

func SetupLogger() error {
	Logger = logger.NewLogger(&lumberjack.Logger{
		Filename:  LogSetting.LogSavePath + "/" + LogSetting.LogFileName + LogSetting.LogFileExt,
		MaxSize:   LogSetting.LogMaxSize,
		MaxAge:    LogSetting.LogMaxAge,
		LocalTime: LogSetting.LogLocalTime,
	}, "", log.LstdFlags).WithCaller(2)
	return nil
}
