package model

import (
	"github.com/space-backend/config"
	"gorm.io/gorm"
)

type Collection struct {
	gorm.Model
	Sid  int64  `json:"sid"`
	Name string `json:"name"`
}

func (c *Collection) Create() error {
	return config.DB.Create(c).Error
}
