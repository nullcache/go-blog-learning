package setting

import "time"

type ServerSetting struct {
	RunMode        string
	HttpPort       string
	ReadTimeout    time.Duration
	WriteTimeout   time.Duration
	MaxHeaderBytes int
}

type AppSetting struct {
	DefaultPageSize int
	MaxPageSize     int
}

type LogSetting struct {
	LogSavePath  string
	LogFileName  string
	LogFileExt   string
	LogMaxSize   int
	LogMaxAge    int
	LogLocalTime bool
}

type DatabaseSetting struct {
	DBType       string
	UserName     string
	Password     string
	Host         string
	DBName       string
	TablePrefix  string
	Charset      string
	ParseTime    bool
	MaxIdleConns int
	MaxOpenConns int
}

func (s *Setting) ReadSection(k string, v any) error {
	return s.vp.UnmarshalKey(k, v)
}
