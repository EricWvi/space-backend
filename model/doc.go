package model

import (
	"github.com/space-backend/config"
	"gorm.io/gorm"
)

type Doc struct {
	gorm.Model
	Sid          int64  `json:"sid"`
	CollectionId int64  `json:"collectionId"`
	Title        string `json:"title"`
}

func (d *Doc) Create() error {
	return config.DB.Create(d).Error
}
