package service

import (
	"github.com/nullcache/go-blog-learning/internal/dao"
	"github.com/nullcache/go-blog-learning/internal/model"
	"github.com/nullcache/go-blog-learning/pkg/app"
)

type CreateTagRequest struct {
	Name      string `form:"name" binding:"required,min=1,max=100"`
	CreatedBy string `form:"created_by" binding:"required,min=1,max=100"`
	Status    uint8  `form:"status,default=1" binding:"oneof=0 1 2"`
}

type UpdateTagRequest struct {
	ID        uint32 `form:"id" binding:"required,gte=1"`
	Name      string `form:"name" binding:"min=1,max=100"`
	Status    uint8  `form:"status,default=1" binding:"oneof=0 1 2"`
	UpdatedBy string `form:"updated_by" binding:"required,min=1,max=100"`
}

type DeleteTagRequest struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
}

type TagListRequest struct {
	Name   string `form:"name" binding:"max=100"`
	Status uint8  `form:"status,default=0" binding:"oneof=0 1 2"`
}

type GetTagRequest struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
}

func (s *Service) NewTag(param *CreateTagRequest) error {
	return dao.New(s.ctx).CreateTag(param.Name, param.Status, param.CreatedBy)
}

func (s *Service) UpdateTag(param *UpdateTagRequest) error {
	return dao.New(s.ctx).UpdateTag(param.ID, param.Name, param.Status, param.UpdatedBy)
}

func (s *Service) DelTag(param *DeleteTagRequest) error {
	return dao.New(s.ctx).DeleteTag(param.ID)
}

func (s *Service) TagList(param *TagListRequest, pager app.Pager) ([]*model.Tag, *app.Pager, error) {
	totalCount, err := dao.New(s.ctx).CountTag(param.Name, param.Status)
	if err != nil {
		return nil, nil, err
	}
	totalPage := pager.TotalCount/pager.PageSize + 1
	p := &app.Pager{TotalCount: totalCount, TotalPage: totalPage}
	items, err := dao.New(s.ctx).GetCountList(param.Name, param.Status, pager.Page, pager.PageSize)
	return items, p, err
}

func (s *Service) GetTag(param *GetTagRequest) (*model.Tag, error) {
	return dao.New(s.ctx).GetTag(param.ID)
}
