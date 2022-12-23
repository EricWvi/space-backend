package editor

import (
	"github.com/gin-gonic/gin"
	"github.com/space-backend/config"
	"github.com/space-backend/handler"
	"github.com/space-backend/model"
)

func (b Base) DeleteAtom(c *gin.Context, req *DeleteAtomRequest) (rsp *DeleteAtomResponse) {
	var err error
	tx := config.DB.Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
			handler.Errorf(c, err.Error())
			rsp = nil
		} else {
			tx.Commit()
		}
	}()

	var doc *model.Doc
	a, err := model.GetAtom(tx, map[string]any{model.Atom_Sid: req.Sid})
	if err != nil {
		return
	}
	doc, err = model.GetDoc(tx, map[string]any{model.Doc_Sid: a.DocId})
	if err != nil {
		return
	}
	_, err = AddAtom(tx, "", a.Sid, a.Name, a.Type, a.DocId, -1, doc.Version)
	if err != nil {
		return
	}
	prev := a.PrevId
	if req.NextId != 0 {
		a, err = model.GetAtom(tx, map[string]any{model.Atom_Sid: req.NextId})
		if err == nil {
			_, err = AddAtom(tx, a.Content, a.Sid, a.Name, a.Type, a.DocId, prev, doc.Version)
		}
	}
	err = model.BumpDocVersion(tx, doc.Sid)

	return &DeleteAtomResponse{}
}

type DeleteAtomRequest struct {
	Sid    model.Sid `json:"sid"`
	NextId model.Sid `json:"nextId"`
}

type DeleteAtomResponse struct {
}
