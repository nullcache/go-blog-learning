package model

import "gorm.io/gorm"

type Article struct {
	*Model
	Title         string `json:"title"`
	Desc          string `json:"desc"`
	Content       string `json:"content"`
	CoverImageUrl string `json:"cover_image_url"`
	TagID         uint32 `json:"-"`
	Tag           *Tag   `json:"tag" gorm:"-"`
}

func (a *Article) Create(db *gorm.DB) error {
	err := db.Create(a).Error
	return err
}

func (a *Article) Delete(db *gorm.DB) error {
	return db.Where("id = ?", a.ID).Delete(a).Error
}

func (a *Article) Update(db *gorm.DB, updateMap map[string]any) error {
	// 文档：updates只能传入map或结构体
	return db.Model(&Article{}).Where("id = ?", a.ID).Updates(updateMap).Error
}

func (a *Article) Get(db *gorm.DB) (*Article, error) {
	var article Article
	if err := db.First(&article, "id=?", a.ID).Error; err != nil {
		return nil, err
	}
	return &article, nil
}

func (a *Article) Count(db *gorm.DB) (int, error) {
	var count int64
	if a.Title != "" {
		db = db.Where("title = ?", a.Title)
	}
	if a.Status != 0 {
		db = db.Where("status = ?", a.Status)
	}
	if a.TagID != 0 {
		db = db.Where("tag_id = ?", a.TagID)
	}
	if err := db.Model(a).Count(&count).Error; err != nil {
		return 0, err
	}
	return int(count), nil
}

func (a *Article) List(db *gorm.DB, pageOffset, pageSize int) ([]*Article, error) {
	var list = make([]*Article, 0, pageSize)
	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}
	if a.Title != "" {
		db = db.Where("title like ?", "%"+a.Title+"%")
	}
	if a.Status != 0 {
		db = db.Where("status = ?", a.Status)
	}
	if a.TagID != 0 {
		db = db.Where("tag_id = ?", a.TagID)
	}
	if err := db.Model(a).Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}
