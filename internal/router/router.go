package router

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/nullcache/go-blog-learning/internal/router/api/v1"
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
			articleGroup.GET("/delArticle", v1.DelArticle)
			articleGroup.GET("/getArticleList", v1.GetArticleList)
			articleGroup.GET("/getArticleByID", v1.GetArticleByID)
		}
		tagGroup := apiV1.Group("/group")
		{
			tagGroup.POST("/newTag", v1.NewTag)
			tagGroup.POST("/editTag", v1.EditTag)
			tagGroup.GET("/delTag", v1.DelTag)
			tagGroup.GET("/getTagList", v1.GetTagList)
			tagGroup.GET("/getTagListByID", v1.GetTagByID)
		}

	}

	return r
}
