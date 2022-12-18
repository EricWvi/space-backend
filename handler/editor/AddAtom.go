package editor

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/space-backend/config"
	"github.com/space-backend/handler"
	"github.com/space-backend/model"
	"github.com/space-backend/service"
	"gorm.io/gorm"
)

func (b Base) AddAtom(c *gin.Context, req *AddAtomRequest) *AddAtomResponse {
	docId, _ := service.ParseSid(req.DocId)
	prevId, _ := service.ParseSid(req.PrevId)

	doc, err := model.GetDoc(config.DB, map[string]any{model.Doc_Sid: docId})
	if err != nil {
		handler.Errorf(c, err.Error())
		return nil
	}

	tx := config.DB.Begin()
	defer tx.Commit()

	a, err := AddAtom(tx, req.Content, req.Sid, req.Name, model.ParseAtomType(req.Type), docId, prevId, doc.Version)
	if err == nil {
		err = model.BumpDocVersion(tx, docId)
	}
	if err != nil {
		tx.Rollback()
		log.Error(err)
		handler.Errorf(c, "failed to add atom")
		return nil
	}

	sid, _ := service.ToSid(a.Sid)
	return &AddAtomResponse{
		Sid: sid,
	}
}

type AddAtomRequest struct {
	Content    string `json:"content"`
	Sid        string `json:"sid"`
	Name       string `json:"name"`
	Type       string `json:"type"`
	DocId      string `json:"docId"`
	PrevId     string `json:"prevId"`
	DocVersion int    `json:"docVersion"`
}

type AddAtomResponse struct {
	Sid string `json:"sid"`
}

func AddAtom(db *gorm.DB,
	Content string,
	Sid string,
	Name string,
	Type int,
	DocId int64,
	PrevId int64,
	DocVersion int,
) (*model.Atom, error) {
	var sidRaw int64
	if Sid != "" {
		sidRaw, _ = service.ParseSid(Sid)
	} else {
		sidRaw = service.NextId()
	}

	a := model.Atom{
		AtomField: model.AtomField{
			Sid:     sidRaw,
			Content: Content,
			Name:    Name,
			Type:    Type,
			DocId:   DocId,
			PrevId:  PrevId,
			Version: DocVersion + 1,
		},
	}

	err := a.Create(db)
	return &a, err
}
