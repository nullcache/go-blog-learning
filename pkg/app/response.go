package app

import (
	"github.com/gin-gonic/gin"
	"github.com/nullcache/go-blog-learning/pkg/errcode"
	"net/http"
)

type Response struct {
	ctx     *gin.Context
	Code    int      `json:"code"`
	Data    any      `json:"data,omitempty"`
	Msg     string   `json:"msg"`
	Details []string `json:"details,omitempty"`
}

type Pager struct {
	Items      []any `json:"items"`
	Page       int   `json:"page"`
	PageSize   int   `json:"page_size"`
	TotalCount int   `json:"total_count"`
	TotalPage  int   `json:"total_page"`
}

func NewResponse(ctx *gin.Context) *Response {
	return &Response{ctx: ctx}
}

func (r *Response) ToResponse(data any) {
	r.Code = errcode.Success.Code()
	r.Msg = "ok"
	r.Data = data
	if data == nil {
		r.Data = gin.H{}
	}
	r.ctx.JSON(http.StatusOK, r)

}

func (r *Response) ToResponseList(list []any, totalPage, totalCount int) {
	r.Code = errcode.Success.Code()
	r.Msg = "ok"
	var pager = Pager{
		Items:      list,
		Page:       GetPage(r.ctx),
		PageSize:   GetPageSize(r.ctx),
		TotalCount: totalCount,
		TotalPage:  totalPage,
	}
	r.Data = pager
	r.ctx.JSON(http.StatusOK, r)
}

func (r *Response) ToErrorResponse(err errcode.Error) {
	r.Code = err.Code()
	r.Msg = err.Msg()
	details := err.Details()
	if len(details) > 0 {
		r.Details = details
	}

	r.ctx.JSON(err.ToStatusCode(), r)
}
