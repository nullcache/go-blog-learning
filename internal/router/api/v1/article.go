package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/nullcache/go-blog-learning/pkg/app"
	"github.com/nullcache/go-blog-learning/pkg/errcode"
)

func NewArticle(c *gin.Context) {
	// test
	app.NewResponse(c).ToErrorResponse(errcode.ServerError)
}
func EditArticle(c *gin.Context) {
}
func DelArticle(c *gin.Context) {
}
func GetArticleList(c *gin.Context) {
}
func GetArticleByID(c *gin.Context) {
}
