package editor

import (
	"github.com/gin-gonic/gin"
	"github.com/space-backend/config"
	"github.com/space-backend/handler"
	"github.com/space-backend/model"
	"github.com/space-backend/service"
)

func (b Base) GetDocAtoms(c *gin.Context, req *GetDocAtomsRequest) *GetDocAtomsResponse {
	doc, err := model.GetDoc(config.DB, map[string]any{model.Doc_Sid: req.DocId})
	if err != nil {
		handler.Errorf(c, err.Error())
		return nil
	}
	version := doc.Version
	if req.Version != 0 {
		version = req.Version
	}
	atoms, err := model.GetAtomViewsByDoc(config.DB, doc.Sid, version)
	if err != nil {
		handler.Errorf(c, "failed to get doc atoms")
		return nil
	}

	for _, a := range atoms {
		if a.Type != model.FormatAtomType(model.Text) {
			a.Link = service.GetFileLink(a.AtomField.Sid)
		}
	}

	return &GetDocAtomsResponse{
		Atoms: atoms,
	}
}

type GetDocAtomsRequest struct {
	DocId   model.Sid `json:"docId"`
	Version int       `json:"version"`
}

type GetDocAtomsResponse struct {
	Atoms []*model.AtomView `json:"atoms"`
}
