package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/nullcache/go-blog-learning/global"
	"github.com/nullcache/go-blog-learning/internal/service"
	"github.com/nullcache/go-blog-learning/pkg/app"
	"github.com/nullcache/go-blog-learning/pkg/errcode"
)

func NewArticle(c *gin.Context) {
	param := service.NewArticleRequest{}
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
	if param.TagID != 0 {
		_, err := svc.GetTag(&service.GetTagRequest{ID: param.TagID})
		if err != nil {
			global.Logger.Errorf("svc.GetTag err: %v", err)
			response.ToErrRespWithStatus(errcode.TagNotFound, 404)
			return
		}
	}
	err := svc.NewArticle(&param)
	if err != nil {
		global.Logger.Errorf("svc.NewArticle err: %v", err)
		response.ToErrorResponse(errcode.ErrorCreateArticleFail)
		return
	}
	response.ToResponse(nil)
}

func EditArticle(c *gin.Context) {
	param := service.UpdateArticleRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.ToErrorList()...))
		return
	}
	svc := service.New(c.Request.Context())
	_, err := svc.GetArticle(&service.GetArticleRequest{ID: param.ID})
	if err != nil {
		global.Logger.Errorf("svc.GetArticle err: %v", err)
		response.ToErrRespWithStatus(errcode.ArticleNotFound, 404)
		return
	}
	err = svc.UpdateArticle(&param)
	if err != nil {
		global.Logger.Errorf("svc.UpdateArticle err: %v", err)
		response.ToErrorResponse(errcode.ErrorUpdateArticleFail)
		return
	}
	response.ToResponse(nil)
}

func DelArticle(c *gin.Context) {
	param := service.DelArticleRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.ToErrorList()...))
		return
	}
	svc := service.New(c.Request.Context())
	_, err := svc.GetArticle(&service.GetArticleRequest{ID: param.ID})
	if err != nil {
		global.Logger.Errorf("svc.GetArticle err: %v", err)
		response.ToErrRespWithStatus(errcode.ArticleNotFound, 404)
		return
	}
	err = svc.DelArticle(&param)
	if err != nil {
		global.Logger.Errorf("svc.DeleteArticle err: %v", err)
		response.ToErrorResponse(errcode.ErrorDelArticleFail)
		return
	}
	response.ToResponse(nil)
}

func GetArticleList(c *gin.Context) {
	param := service.ArticleListRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.ToErrorList()...))
		return
	}
	svc := service.New(c.Request.Context())
	pagerIn := app.Pager{Page: app.GetPage(c), PageSize: app.GetPageSize(c)}
	list, pagerOut, err := svc.ArticleList(&param, pagerIn)
	if err != nil {
		global.Logger.Errorf("svc.GetArticleList err: %v", err)
		response.ToErrorResponse(errcode.ErrorGetArticleListFail)
		return
	}
	response.ToResponseList(list, pagerOut.TotalPage, pagerOut.TotalCount)
}

func GetArticleByID(c *gin.Context) {
	param := service.GetArticleRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.ToErrorList()...))
		return
	}
	svc := service.New(c.Request.Context())
	res, err := svc.GetArticle(&param)
	if err != nil {
		global.Logger.Errorf("svc.GetArticle err: %v", err)
		response.ToErrRespWithStatus(errcode.ArticleNotFound, 404)
		return
	}
	response.ToResponse(res)
}
