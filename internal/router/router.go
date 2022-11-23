package router

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/nullcache/go-blog-learning/internal/router/api/v1"
	"github.com/nullcache/go-blog-learning/pkg/app"
	"github.com/nullcache/go-blog-learning/pkg/errcode"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery(), gin.Logger())
	apiV1 := r.Group("/api/v1")
	{
		articleGroup := apiV1.Group("/article")
		{
			articleGroup.POST("/newArticle", v1.NewArticle)
			articleGroup.POST("/editArticle", v1.EditArticle)
			articleGroup.POST("/delArticle", v1.DelArticle)
			articleGroup.GET("/getArticleList", v1.GetArticleList)
			articleGroup.GET("/getArticleById", v1.GetArticleByID)
		}
		tagGroup := apiV1.Group("/tag")
		{
			tagGroup.POST("/newTag", v1.NewTag)
			tagGroup.POST("/editTag", v1.EditTag)
			tagGroup.POST("/delTag", v1.DelTag)
			tagGroup.GET("/getTagList", v1.GetTagList)
			tagGroup.GET("/getTagById", v1.GetTagByID)

		}

	}
	r.NoRoute(func(c *gin.Context) {
		app.NewResponse(c).ToErrorResponse(errcode.NotFound)
	})
	return r
}
