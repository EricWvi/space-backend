package editor

import (
	"github.com/gin-gonic/gin"
	"github.com/space-backend/config"
	"github.com/space-backend/handler"
	"github.com/space-backend/model"
	"github.com/space-backend/service"
)

func (b Base) GetDocAtoms(c *gin.Context, req *GetDocAtomsRequest) *GetDocAtomsResponse {
	docId, _ := service.ParseSid(req.DocId)
	doc, err := model.GetDoc(config.DB, map[string]any{model.Doc_Sid: docId})
	if err != nil {
		handler.Errorf(c, err.Error())
		return nil
	}
	atoms, err := model.GetAtomViewsByDoc(config.DB, docId, doc.Version)
	if err != nil {
		handler.Errorf(c, "failed to get doc atoms")
		return nil
	}

	return &GetDocAtomsResponse{
		Atoms: atoms,
	}
}

type GetDocAtomsRequest struct {
	DocId string `json:"docId"`
}

type GetDocAtomsResponse struct {
	Atoms []*model.AtomView `json:"atoms"`
}
