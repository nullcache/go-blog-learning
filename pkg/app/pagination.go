package app

import (
	"github.com/gin-gonic/gin"
	"github.com/nullcache/go-blog-learning/global"
	"github.com/nullcache/go-blog-learning/pkg/convert"
)

func GetPage(c *gin.Context) int {
	page := convert.StrTo(c.Query("page")).MustInt()
	if page <= 0 {
		page = 1
	}
	return page
}

func GetPageSize(c *gin.Context) int {
	pageSize := convert.StrTo(c.Query("page_size")).MustInt()
	if pageSize <= 0 {
		return global.AppSetting.DefaultPageSize
	}
	if pageSize > global.AppSetting.MaxPageSize {
		return global.AppSetting.MaxPageSize
	}
	return pageSize
}

func GetPageOffset(page, pageSize int) int {
	offset := 0
	if page > 0 {
		offset = (page - 1) * pageSize
	}

	return offset
}
