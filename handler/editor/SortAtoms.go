package editor

import (
	"github.com/gin-gonic/gin"
	"github.com/space-backend/config"
	"github.com/space-backend/handler"
	"github.com/space-backend/model"
)

func (b Base) SortAtoms(c *gin.Context, req *SortAtomsRequest) *SortAtomsResponse {
	tx := config.DB.Begin()
	defer tx.Commit()

	var doc *model.Doc
	errList := ""
	for _, v := range req.Atoms {
		a, err := model.GetAtom(tx, map[string]any{model.Atom_Sid: v.Sid})
		if a != nil && doc == nil {
			doc, _ = model.GetDoc(tx, map[string]any{model.Doc_Sid: a.DocId})
		}
		_, err = AddAtom(tx, a.Content, v.Sid, a.Name, a.Type, a.DocId, v.PrevId, doc.Version)

		if err != nil {
			errList += err.Error() + "\n"
			break
		}
	}
	if errList == "" && doc != nil {
		err := model.BumpDocVersion(tx, doc.Sid)
		if err != nil {
			errList += err.Error() + "\n"
		}
	}
	if errList != "" {
		tx.Rollback()
		handler.Errorf(c, errList)
		return nil
	}

	return &SortAtomsResponse{}
}

type SortAtomsRequest struct {
	Atoms []struct {
		Sid    model.Sid `json:"sid"`
		PrevId model.Sid `json:"prevId"`
	} `json:"atoms"`
}

type SortAtomsResponse struct {
}
