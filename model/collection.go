package model

import (
	"errors"
	"gorm.io/gorm"
)

type Collection struct {
	gorm.Model
	CollectionField
}

const (
	Collection_Table = "collections"
	Collection_Sid   = "sid"
	Collection_Name  = "name"
)

type CollectionField struct {
	Sid  Sid    `json:"sid"`
	Name string `json:"name"`
}

func (c *Collection) Create(db *gorm.DB) error {
	return db.Create(c).Error
}

func GetCollection(db *gorm.DB, where map[string]any) (c *Collection, err error) {
	var colls []Collection
	db.Table(Collection_Table).Where(where).Find(&colls)
	if len(colls) != 1 {
		err = errors.New("collection not found")
		return
	}
	c = &colls[0]
	return
}

type CollectionView struct {
	Sid string `json:"sid"`
	CollectionField
}

func GetCollectionViews(db *gorm.DB) (views []*CollectionView, err error) {
	var collections []CollectionField
	err = db.Table(Collection_Table).
		Find(&collections).Error
	for _, c := range collections {
		views = append(views, &CollectionView{
			Sid:             c.Sid.String(),
			CollectionField: c,
		})
	}
	return
}
