package service

import (
	"github.com/liangyaopei/structmap"
	"github.com/nullcache/go-blog-learning/internal/dao"
	"github.com/nullcache/go-blog-learning/internal/model"
	"github.com/nullcache/go-blog-learning/pkg/app"
)

type NewArticleRequest struct {
	Title         string `binding:"required,max=100" form:"title" json:"title"`
	Desc          string `binding:"max=255" form:"desc" json:"desc" `
	Content       string `binding:"required" form:"content" json:"content"`
	CoverImageUrl string `binding:"omitempty,url,max=255" form:"cover_image_url" json:"cover_image_url"`
	TagID         uint32 `binding:"omitempty,gte=1,max=100" form:"tag_id" json:"tag_id"`
	CreatedBy     string `binding:"required,max=100" form:"created_by" json:"created_by"`
	Status        uint8  `binding:"oneof=0 1 2" form:"status,default=1" json:"status"`
}

type UpdateArticleRequest struct {
	ID            uint32  `binding:"required,max=100" form:"id" json:"id,omitempty"`
	Title         string  `binding:"max=100" form:"title" json:"title,omitempty"`
	Desc          *string `binding:"omitempty,max=255" form:"desc" json:"desc"`
	Content       string  `form:"content" json:"content,omitempty"`
	CoverImageUrl *string `binding:"omitempty,url,max=255" form:"cover_image_url" json:"cover_image_url"`
	TagID         *uint32 `binding:"omitempty,gte=0,max=100" form:"tag_id" json:"tag_id"`
	UpdatedBy     string  `binding:"required,max=100" form:"updated_by" json:"updated_by,omitempty"`
	Status        *uint8  `binding:"omitempty,oneof=0 1 2" form:"status" json:"status"`
}

type DelArticleRequest struct {
	ID uint32 `binding:"required,gte=1" form:"id" json:"id"`
}

type ArticleListRequest struct {
	Title  string `binding:"omitempty,max=100" form:"title" json:"title"`
	Status uint8  `binding:"oneof=0 1 2" form:"status" json:"status"`
	TagID  uint32 `binding:"omitempty,gte=1,max=100" form:"tag_id" json:"tag_id"`
}

type GetArticleRequest struct {
	ID uint32 `binding:"required,gte=1" form:"id" json:"id"`
}

func (s *Service) NewArticle(r *NewArticleRequest) error {
	return dao.New(s.ctx).NewArticle(&model.Article{
		Model: &model.Model{
			CreatedBy: r.CreatedBy,
			Status:    r.Status,
		},
		Title:         r.Title,
		Desc:          r.Desc,
		Content:       r.Content,
		CoverImageUrl: r.CoverImageUrl,
		TagID:         r.TagID,
	})
}

func (s *Service) DelArticle(param *DelArticleRequest) error {
	return dao.New(s.ctx).DeleteArticle(param.ID)
}

func (s *Service) UpdateArticle(r *UpdateArticleRequest) error {
	// use json for less writing
	updateMap, err := structmap.StructToMap(r, "json", "")
	if err != nil {
		return err
	}
	return dao.New(s.ctx).UpdateArticle(r.ID, updateMap)
}

func (s *Service) ArticleList(param *ArticleListRequest, pager app.Pager) ([]*model.Article, *app.Pager, error) {
	totalCount, err := dao.New(s.ctx).CountArticle(param.Title, param.Status, param.TagID)
	if err != nil {
		return nil, nil, err
	}
	totalPage := pager.TotalCount/pager.PageSize + 1
	p := &app.Pager{TotalCount: totalCount, TotalPage: totalPage}
	items, err := dao.New(s.ctx).GetArticleList(param.Title, param.Status, param.TagID, pager.Page, pager.PageSize)
	for _, item := range items {
		if item.TagID != 0 {
			var tag *model.Tag
			tag, err = dao.New(s.ctx).GetTag(item.TagID)
			item.Tag = tag
		}
	}
	return items, p, err
}

func (s *Service) GetArticle(param *GetArticleRequest) (*model.Article, error) {
	article, err := dao.New(s.ctx).GetArticle(param.ID)
	if err != nil {
		return nil, err
	}
	if article.TagID != 0 {
		var tag *model.Tag
		tag, err = dao.New(s.ctx).GetTag(article.TagID)
		article.Tag = tag
	}
	return article, err
}
