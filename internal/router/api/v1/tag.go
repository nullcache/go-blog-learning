package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/nullcache/go-blog-learning/global"
	"github.com/nullcache/go-blog-learning/internal/service"
	"github.com/nullcache/go-blog-learning/pkg/app"
	"github.com/nullcache/go-blog-learning/pkg/errcode"
)

func NewTag(c *gin.Context) {
	param := service.CreateTagRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.ToErrorList()...))
		return
	}
	if param.Status == 0 {
		param.Status = 1
	}
	svc := service.New(c.Request.Context())
	err := svc.NewTag(&param)
	if err != nil {
		global.Logger.Errorf("svc.GetTagList err: %v", err)
		response.ToErrorResponse(errcode.ErrorCreateTagFail)
		return
	}
	response.ToResponse(nil)
}

func EditTag(c *gin.Context) {
	param := service.UpdateTagRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.ToErrorList()...))
		return
	}
	if param.Status == 0 {
		param.Status = 1
	}
	svc := service.New(c.Request.Context())
	_, err := svc.GetTag(&service.GetTagRequest{ID: param.ID})
	if err != nil {
		global.Logger.Errorf("svc.GetTag err: %v", err)
		response.ToErrRespWithStatus(errcode.TagNotFound, 404)
		return
	}
	err = svc.UpdateTag(&param)
	if err != nil {
		global.Logger.Errorf("svc.UpdateTag err: %v", err)
		response.ToErrorResponse(errcode.ErrorUpdateTagFail)
		return
	}
	response.ToResponse(nil)
}

func DelTag(c *gin.Context) {
	param := service.DeleteTagRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.ToErrorList()...))
		return
	}
	svc := service.New(c.Request.Context())
	_, err := svc.GetTag(&service.GetTagRequest{ID: param.ID})
	if err != nil {
		global.Logger.Errorf("svc.GetTag err: %v", err)
		response.ToErrRespWithStatus(errcode.TagNotFound, 404)
		return
	}
	err = svc.DelTag(&param)
	if err != nil {
		global.Logger.Errorf("svc.DeleteTag err: %v", err)

		response.ToErrorResponse(errcode.ErrorDeleteTagFail)
		return
	}
	response.ToResponse(nil)
}

func GetTagList(c *gin.Context) {
	param := service.TagListRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.ToErrorList()...))
		return
	}
	svc := service.New(c.Request.Context())
	pagerIn := app.Pager{Page: app.GetPage(c), PageSize: app.GetPageSize(c)}
	list, pagerOut, err := svc.TagList(&param, pagerIn)
	if err != nil {
		global.Logger.Errorf("svc.GetTagList err: %v", err)
		response.ToErrorResponse(errcode.ErrorGetTagListFail)
		return
	}
	response.ToResponseList(list, pagerOut.TotalPage, pagerOut.TotalCount)
}

func GetTagByID(c *gin.Context) {
	param := service.GetTagRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.ToErrorList()...))
		return
	}
	svc := service.New(c.Request.Context())
	res, err := svc.GetTag(&param)
	if err != nil {
		global.Logger.Errorf("svc.GetTag err: %v", err)
		response.ToErrRespWithStatus(errcode.TagNotFound, 404)
		return
	}
	response.ToResponse(res)
}
