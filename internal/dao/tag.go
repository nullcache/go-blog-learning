package dao

import (
	"github.com/nullcache/go-blog-learning/internal/model"
	"github.com/nullcache/go-blog-learning/pkg/app"
)

func (d *Dao) CountTag(name string, status uint8) (int, error) {
	m := model.Tag{Name: name, Model: &model.Model{
		Status: status,
	}}
	return m.Count(d.db)
}

func (d *Dao) GetCountList(name string, status uint8, page, pageSize int) ([]*model.Tag, error) {
	m := model.Tag{Name: name, Model: &model.Model{
		Status: status,
	}}
	return m.List(d.db, app.GetPageOffset(page, pageSize), pageSize)
}

func (d *Dao) CreateTag(name string, status uint8, createdBy string) error {
	m := model.Tag{Name: name, Model: &model.Model{
		Status:    status,
		CreatedBy: createdBy,
	}}
	return m.Create(d.db)
}

func (d *Dao) UpdateTag(id uint32, name string, status uint8, updatedBy string) error {
	updateMap := map[string]any{
		"updated_by": updatedBy,
	}
	if name != "" {
		updateMap["name"] = name
	}
	m := model.Tag{
		Name: name,
		Model: &model.Model{
			ID:        id,
			UpdatedBy: updatedBy,
			Status:    status,
		},
	}
	return m.Update(d.db, updateMap)
}

func (d *Dao) DeleteTag(id uint32) error {
	m := model.Tag{Model: &model.Model{ID: id}}
	return m.Delete(d.db)
}

func (d *Dao) GetTag(id uint32) (*model.Tag, error) {
	m := model.Tag{Model: &model.Model{ID: id}}
	return m.Get(d.db)
}
