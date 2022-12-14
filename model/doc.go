package model

import (
	"github.com/space-backend/config"
	"github.com/space-backend/service"
	"gorm.io/gorm"
)

type Doc struct {
	gorm.Model
	DocField
}

type DocField struct {
	Sid          int64  `json:"sid"`
	CollectionId int64  `json:"collectionId"`
	Title        string `json:"title"`
}

func (d *Doc) Create() error {
	return config.DB.Create(d).Error
}

type DocView struct {
	Sid          string `json:"sid"`
	CollectionId string `json:"collectionId"`
	DocField
}

func GetDocViewsByCollectionId(cid int64) (views []*DocView, err error) {
	var docs []DocField
	err = config.DB.Table("docs").Where("collection_id = ?", cid).
		Find(&docs).Error
	for _, d := range docs {
		sid, _ := service.ToSid(d.Sid)
		cid, _ := service.ToSid(d.CollectionId)
		views = append(views, &DocView{
			Sid:          sid,
			CollectionId: cid,
			DocField:     d,
		})
	}
	return
}
