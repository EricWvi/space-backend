package model

import (
	"errors"
	log "github.com/sirupsen/logrus"
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

const (
	Atom_Table   = "atoms"
	Atom_Sid     = "sid"
	Atom_Content = "content"
	Atom_Name    = "name"
	Atom_Type    = "type"
	Atom_Version = "version"
	Atom_DocId   = "doc_id"
	Atom_PrevId  = "prev_id"
)

type AtomField struct {
	Sid     int64  `json:"sid"`
	Content string `json:"content"`
	Name    string `json:"name"`
	Type    int    `json:"type"`
	Version int    `json:"version"`
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

func (a *Atom) Create(db *gorm.DB) error {
	return db.Create(a).Error
}

func GetAtom(db *gorm.DB, where map[string]any) (a *Atom, err error) {
	var atoms []Atom
	db.Table(Atom_Table).Where(where).Order(Atom_Version + " DESC").Find(&atoms)
	if len(atoms) == 0 {
		err = errors.New("atom not found")
		return
	}
	a = &atoms[0]
	return
}

//func UpdateAtomPrevId(db *gorm.DB, sid int64, prevId int64) error {
//	err := db.Table(Atom_Table).Where(Atom_Sid, sid).Update(Atom_PrevId, prevId).Error
//	if err != nil {
//		log.Error(err)
//	}
//	return err
//}

type AtomView struct {
	Sid    string `json:"sid"`
	DocId  string `json:"docId"`
	Type   string `json:"type"`
	PrevId string `json:"prevId"`
	AtomField
}

func GetAtomViewsByDoc(db *gorm.DB, docId int64, docVersion int) (views []*AtomView, err error) {
	views = make([]*AtomView, 0)
	var rows []Atom
	err = db.Raw("select atoms.*\n"+
		"from (SELECT max(version) as version, sid\n"+
		"      FROM `atoms`\n"+
		"      WHERE `doc_id` = ?	\n"+
		"        and version <= ?\n"+
		"        and deleted_at is null\n"+
		"      group by sid) t1\n"+
		"         inner join atoms on t1.version = atoms.version and t1.sid = atoms.sid\n", docId, docVersion).
		Scan(&rows).Error
	if err != nil {
		log.Error(err)
		return
	}
	var atoms []AtomField
	for i := range rows {
		atoms = append(atoms, rows[i].AtomField)
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
