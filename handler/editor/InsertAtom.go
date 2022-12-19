package editor

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/space-backend/config"
	"github.com/space-backend/handler"
	"github.com/space-backend/model"
)

func (b Base) InsertAtom(c *gin.Context, req *InsertAtomRequest) *InsertAtomResponse {
	doc, err := model.GetDoc(config.DB, map[string]any{model.Doc_Sid: req.DocId})
	if err != nil {
		handler.Errorf(c, err.Error())
		return nil
	}

	tx := config.DB.Begin()
	defer tx.Commit()

	a, err := AddAtom(tx, req.Content, 0, req.Name, model.ParseAtomType(req.Type), doc.Sid, req.PrevId, doc.Version)
	if req.NextId != 0 {
		n, e := model.GetAtom(config.DB, map[string]any{model.Atom_Sid: req.NextId})
		err = e
		_, err = AddAtom(tx, n.Content, n.Sid, n.Name, n.Type, n.DocId, a.Sid, doc.Version)
	}
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
	return &InsertAtomResponse{
		Sid: sid,
	}
}

type InsertAtomRequest struct {
	Content string    `json:"content"`
	Name    string    `json:"name"`
	Type    string    `json:"type"`
	DocId   model.Sid `json:"docId"`
	PrevId  model.Sid `json:"prevId"`
	NextId  model.Sid `json:"nextId"`
}

type InsertAtomResponse struct {
	Sid string `json:"sid"`
}
