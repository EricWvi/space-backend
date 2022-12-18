package model

import (
	"errors"
	log "github.com/sirupsen/logrus"
	"github.com/space-backend/service"
	"gorm.io/gorm"
)

type Doc struct {
	gorm.Model
	DocField
}

const (
	Doc_Table        = "docs"
	Doc_Sid          = "sid"
	Doc_CollectionId = "collection_id"
	Doc_Title        = "title"
	Doc_Version      = "version"
)

type DocField struct {
	Sid          int64  `json:"sid"`
	CollectionId int64  `json:"collectionId"`
	Title        string `json:"title"`
	Version      int    `json:"version"`
}

func (d *Doc) Create(db *gorm.DB) error {
	return db.Create(d).Error
}

func GetDoc(db *gorm.DB, where map[string]any) (d *Doc, err error) {
	var docs []Doc
	db.Table(Doc_Table).Where(where).Find(&docs)
	if len(docs) != 1 {
		err = errors.New("doc not found")
		return
	}
	d = &docs[0]
	return
}

func BumpDocVersion(db *gorm.DB, sid int64) error {
	err := db.Table(Doc_Table).Where(Doc_Sid, sid).Update(Doc_Version, gorm.Expr("version + 1")).Error
	if err != nil {
		log.Error(err)
	}
	return err
}

type DocView struct {
	Sid          string `json:"sid"`
	CollectionId string `json:"collectionId"`
	DocField
}

func GetDocViewsByCollectionId(db *gorm.DB, cid int64) (views []*DocView, err error) {
	views = make([]*DocView, 0)
	var docs []DocField
	err = db.Table(Doc_Table).Where(Doc_CollectionId, cid).
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
