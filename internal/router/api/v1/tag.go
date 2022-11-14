package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/nullcache/go-blog-learning/internal/service"
	"github.com/nullcache/go-blog-learning/pkg/app"
	"github.com/nullcache/go-blog-learning/pkg/errcode"
)

func NewTag(c *gin.Context) {

}
func EditTag(c *gin.Context) {
}
func DelTag(c *gin.Context) {
}
func GetTagList(c *gin.Context) {
	// test
	var model = service.TagListRequest{}
	valid, errs := app.BindAndValid(c, &model)
	if !valid {
		app.NewResponse(c).ToErrorResponse(errcode.InvalidParams.WithDetails(errs.ToErrorList()...))
	} else {
		app.NewResponse(c).ToResponse([]any{})
	}

}
func GetTagByID(c *gin.Context) {

}
