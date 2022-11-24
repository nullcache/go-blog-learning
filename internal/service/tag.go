package service

import (
	"github.com/liangyaopei/structmap"
	"github.com/nullcache/go-blog-learning/internal/dao"
	"github.com/nullcache/go-blog-learning/internal/model"
	"github.com/nullcache/go-blog-learning/pkg/app"
)

type CreateTagRequest struct {
	Name      string `binding:"required,min=1,max=100" form:"name" json:"name"`
	CreatedBy string `binding:"required,min=1,max=100" form:"created_by" json:"created_by"`
	Status    uint8  `binding:"oneof=0 1 2" form:"status,default=1" json:"status"`
}

type UpdateTagRequest struct {
	ID        uint32 `binding:"required,gte=1" form:"id" json:"id,omitempty"`
	Name      string `binding:"omitempty,min=1,max=100" form:"name" json:"name,omitempty"`
	Status    *uint8 `binding:"omitempty,oneof=0 1 2" form:"status" json:"status"`
	UpdatedBy string `binding:"required,min=1,max=100" form:"updated_by" json:"updated_by,omitempty"`
}

type DeleteTagRequest struct {
	ID uint32 `binding:"required,gte=1" form:"id" json:"id" `
}

type TagListRequest struct {
	Name   string `binding:"omitempty,max=100" form:"name" json:"name"`
	Status uint8  `binding:"oneof=0 1 2" form:"status" json:"status"`
}

type GetTagRequest struct {
	ID uint32 `binding:"required,gte=1" form:"id" json:"id"`
}

func (s *Service) NewTag(param *CreateTagRequest) error {
	return dao.New(s.ctx).CreateTag(param.Name, param.Status, param.CreatedBy)
}

func (s *Service) UpdateTag(param *UpdateTagRequest) error {
	// use json for less writing
	updateMap, err := structmap.StructToMap(param, "json", "")
	if err != nil {
		return err
	}
	return dao.New(s.ctx).UpdateTag(param.ID, updateMap)
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
	items, err := dao.New(s.ctx).GetTagList(param.Name, param.Status, pager.Page, pager.PageSize)
	return items, p, err
}

func (s *Service) GetTag(param *GetTagRequest) (*model.Tag, error) {
	return dao.New(s.ctx).GetTag(param.ID)
}
