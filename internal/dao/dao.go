package dao

import (
	"context"
	"github.com/nullcache/go-blog-learning/internal/model"
	"gorm.io/gorm"
)

type Dao struct {
	db *gorm.DB
}

func New(ctx context.Context) *Dao {
	return &Dao{
		db: model.GetDBWithCtx(ctx),
	}
}
