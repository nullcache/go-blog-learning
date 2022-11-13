package initDB

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/nullcache/go-blog-learning/global"
	"github.com/nullcache/go-blog-learning/pkg/setting"
)

func NewDBEngine(dbSetting *setting.DatabaseSetting) (*gorm.DB, error) {
	s := "%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local"
	db, err := gorm.Open(dbSetting.DBType, fmt.Sprintf(s,
		dbSetting.UserName,
		dbSetting.Password,
		dbSetting.Host,
		dbSetting.DBName,
		dbSetting.Charset,
		dbSetting.ParseTime,
	))
	if err != nil {
		return nil, err
	}

	if global.ServerSetting.RunMode == "debug" {
		db.LogMode(true)
	}
	db.SingularTable(true) // 如果设置为true,`User`的默认表名为`user`,使用`TableName`设置的表名不受影响
	db.DB().SetMaxIdleConns(dbSetting.MaxIdleConns)
	db.DB().SetMaxOpenConns(dbSetting.MaxOpenConns)

	return db, nil
}

func SetupDBEngine() (err error) {
	global.DBEngine, err = NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return
	}
	return
}
