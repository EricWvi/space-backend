package model

import (
	log "github.com/sirupsen/logrus"
	"github.com/space-backend/config"
	"github.com/space-backend/service"
	"gorm.io/gorm"
)

const (
	MalAtomType = iota
	Text
	Image
	Audio
	Video
)

type Atom struct {
	gorm.Model
	AtomField
}

type AtomField struct {
	Sid     int64  `json:"sid"`
	Content string `json:"content"`
	Link    string `json:"link"`
	Type    int    `json:"type"`
	DocId   int64  `json:"docId"`
	PrevId  int64  `json:"prevId"`
}

func (a AtomField) Prev() int64 {
	return a.PrevId
}

func (a AtomField) Curr() int64 {
	return a.Sid
}

func (a AtomField) IsHead() bool {
	return a.PrevId == 0
}

func ParseAtomType(t string) int {
	switch t {
	case "text":
		return Text
	case "image":
		return Image
	case "audio":
		return Audio
	case "video":
		return Video
	}
	return MalAtomType
}

func FormatAtomType(aType int) string {
	switch aType {
	case Text:
		return "text"
	case Image:
		return "image"
	case Audio:
		return "audio"
	case Video:
		return "video"
	}
	return ""
}

func (a *Atom) Create() error {
	return config.DB.Create(a).Error
}

func UpdateAtomPrevId(sid int64, prevId int64) error {
	err := config.DB.Table("atoms").Where("sid = ?", sid).Update("prev_id", prevId).Error
	if err != nil {
		log.Error(err)
	}
	return err
}

type AtomView struct {
	Sid    string `json:"sid"`
	DocId  string `json:"docId"`
	Type   string `json:"type"`
	PrevId string `json:"prevId"`
	AtomField
}

func GetAtomViewsByDocId(docId int64) (views []*AtomView, err error) {
	var atoms []AtomField
	err = config.DB.Table("atoms").Where("doc_id = ?", docId).
		Find(&atoms).Error
	if err != nil {
		log.Error(err)
	}
	var nodes []ListNode[int64]
	for i := range atoms {
		nodes = append(nodes, &atoms[i])
	}
	for _, c := range Sort(nodes) {
		a := c.(*AtomField)
		sid, _ := service.ToSid(a.Sid)
		did, _ := service.ToSid(a.DocId)
		pid, _ := service.ToSid(a.PrevId)
		views = append(views, &AtomView{
			Sid:       sid,
			DocId:     did,
			Type:      FormatAtomType(a.Type),
			PrevId:    pid,
			AtomField: *a,
		})
	}
	return
}
