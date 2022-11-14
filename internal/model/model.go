package model

import (
	"fmt"
	"github.com/nullcache/go-blog-learning/global"
	"github.com/nullcache/go-blog-learning/pkg/setting"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/soft_delete"
)

type Model struct {
	ID        uint32                `json:"id"`
	CreatedAt uint32                `json:"create_time"`
	UpdatedAt uint32                `json:"update_time"`
	DeletedAt soft_delete.DeletedAt `json:"-"` // gorm查询会自动过滤被软删除的结果，软删除即 deleted_at != 0
	CreatedBy string                `json:"create_by"`
	UpdatedBy string                `json:"update_by"`
	Status    uint8                 `json:"status"` // 1启用,2禁用,99删除

}

func NewDBEngine(dbSetting *setting.DatabaseSetting) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local",
		dbSetting.UserName,
		dbSetting.Password,
		dbSetting.Host,
		dbSetting.DBName,
		dbSetting.Charset,
		dbSetting.ParseTime)

	config := gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,    // 如果设置为true,`User`的默认表名为`user`,使用`TableName`设置的表名不受影响
			TablePrefix:   "blog_", // table name prefix, table for `User` would be `t_users`
		},
	}
	if global.ServerSetting.RunMode == "debug" {
		config.Logger = logger.Default.LogMode(logger.Info)
	}

	db, err := gorm.Open(mysql.Open(dsn), &config)
	if err != nil {
		return nil, err
	}
	dbConfig, _ := db.DB()

	dbConfig.SetMaxIdleConns(dbSetting.MaxIdleConns)
	dbConfig.SetMaxOpenConns(dbSetting.MaxOpenConns)
	return db, nil
}
