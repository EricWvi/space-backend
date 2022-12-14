package model

import (
	"github.com/space-backend/config"
	"github.com/space-backend/service"
	"gorm.io/gorm"
)

type Collection struct {
	gorm.Model
	CollectionField
}

type CollectionField struct {
	Sid  int64  `json:"sid"`
	Name string `json:"name"`
}

func (c *Collection) Create() error {
	return config.DB.Create(c).Error
}

type CollectionView struct {
	Sid string `json:"sid"`
	CollectionField
}

func GetCollectionViews() (views []*CollectionView, err error) {
	var collections []CollectionField
	err = config.DB.Table("collections").
		Find(&collections).Error
	for _, c := range collections {
		sid, _ := service.ToSid(c.Sid)
		views = append(views, &CollectionView{
			Sid:             sid,
			CollectionField: c,
		})
	}
	return
}
