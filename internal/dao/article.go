package dao

import (
	"github.com/nullcache/go-blog-learning/internal/model"
	"github.com/nullcache/go-blog-learning/pkg/app"
)

func (d *Dao) NewArticle(a *model.Article) error {
	m := model.Article{
		Title:         a.Title,
		Desc:          a.Desc,
		Content:       a.Content,
		CoverImageUrl: a.CoverImageUrl,
		TagID:         a.TagID,
		Model: &model.Model{
			CreatedBy: a.CreatedBy,
			Status:    a.Status,
		},
	}
	return m.Create(d.db)
}

func (d *Dao) DeleteArticle(id uint32) error {
	m := model.Article{Model: &model.Model{ID: id}}
	return m.Delete(d.db)
}

func (d *Dao) UpdateArticle(a *model.Article) error {
	updateMap := map[string]any{
		"updated_by":    a.UpdatedBy,
		"desc":          a.Desc,
		"CoverImageUrl": a.CoverImageUrl,
		"tag_id":        a.TagID,
		"status":        a.Status,
	}
	if a.Title != "" {
		updateMap["title"] = a.Title
	}
	if a.Content != "" {
		updateMap["content"] = a.Content
	}

	m := model.Article{
		Model: &model.Model{
			ID: a.ID,
		},
	}
	return m.Update(d.db, updateMap)
}

func (d *Dao) GetArticle(id uint32) (*model.Article, error) {
	m := model.Article{Model: &model.Model{ID: id}}
	return m.Get(d.db)
}

func (d *Dao) CountArticle(title string, status uint8, tagId uint32) (int, error) {
	m := model.Article{Title: title, TagID: tagId, Model: &model.Model{
		Status: status,
	}}
	return m.Count(d.db)
}

func (d *Dao) GetArticleList(title string, status uint8, tagId uint32, page, pageSize int) ([]*model.Article, error) {
	m := model.Article{
		Title: title,
		TagID: tagId,
		Model: &model.Model{
			Status: status,
		}}
	return m.List(d.db, app.GetPageOffset(page, pageSize), pageSize)
}
