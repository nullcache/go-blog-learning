package global

import (
	"github.com/nullcache/go-blog-learning/pkg/setting"
)

var (
	ServerSetting   *setting.ServerSetting
	AppSetting      *setting.AppSetting
	LogSetting      *setting.LogSetting
	DatabaseSetting *setting.DatabaseSetting
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
