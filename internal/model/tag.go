package model

import (
	"gorm.io/gorm"
)

type Tag struct {
	*Model
	Name string `json:"name"`
}

func (t *Tag) Count(db *gorm.DB) (int, error) {
	var count int64
	if t.Name != "" {
		db = db.Where("name = ?", t.Name)
	}
	if t.Status != 0 {
		db = db.Where("status = ?", t.Status)
	}
	if err := db.Model(t).Count(&count).Error; err != nil {
		return 0, err
	}
	return int(count), nil
}

func (t *Tag) List(db *gorm.DB, pageOffset, pageSize int) ([]*Tag, error) {
	var list = make([]*Tag, 0, pageSize)
	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}
	if t.Name != "" {
		db = db.Where("name = ?", t.Name)
	}
	if t.Status != 0 {
		db = db.Where("status = ?", t.Status)
	}

	if err := db.Model(t).Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

func (t *Tag) Create(db *gorm.DB) error {
	return db.Create(t).Error
}

func (t *Tag) Update(db *gorm.DB, updateMap map[string]any) error {
	// 文档：updates只能传入map或结构体
	return db.Model(&Tag{}).Where("id = ?", t.ID).Updates(updateMap).Error
}

func (t *Tag) Delete(db *gorm.DB) error {
	return db.Where("id = ?", t.ID).Delete(t).Error
}

func (t *Tag) Get(db *gorm.DB) (*Tag, error) {
	var tag Tag
	if err := db.First(&tag, "id=?", t.ID).Error; err != nil {
		return nil, err
	}
	return &tag, nil
}
