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
	doc, err := model.GetDoc(config.DB, map[string]any{model.Doc_Sid: req.DocId})
	if err != nil {
		handler.Errorf(c, err.Error())
		return nil
	}

	tx := config.DB.Begin()
	defer tx.Commit()

	a, err := AddAtom(tx, req.Content, req.Sid, req.Name, model.ParseAtomType(req.Type), req.DocId, req.PrevId, doc.Version)
	if err == nil {
		err = model.BumpDocVersion(tx, doc.Sid)
	}
	if err != nil {
		tx.Rollback()
		log.Error(err)
		handler.Errorf(c, "failed to add atom")
		return nil
	}

	sid := a.Sid.String()
	return &AddAtomResponse{
		Sid: sid,
	}
}

type AddAtomRequest struct {
	Content string    `json:"content"`
	Sid     model.Sid `json:"sid"`
	Name    string    `json:"name"`
	Type    string    `json:"type"`
	DocId   model.Sid `json:"docId"`
	PrevId  model.Sid `json:"prevId"`
}

type AddAtomResponse struct {
	Sid string `json:"sid"`
}

func AddAtom(db *gorm.DB,
	Content string,
	Sid model.Sid,
	Name string,
	Type int,
	DocId model.Sid,
	PrevId model.Sid,
	DocVersion int,
) (*model.Atom, error) {
	if Sid == 0 {
		Sid = service.NextId()
	}

	a := model.Atom{
		AtomField: model.AtomField{
			Sid:     Sid,
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
