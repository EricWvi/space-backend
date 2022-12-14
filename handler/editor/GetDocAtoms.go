package editor

import (
	"github.com/gin-gonic/gin"
	"github.com/space-backend/handler"
	"github.com/space-backend/model"
	"github.com/space-backend/service"
)

func (b Base) GetDocAtoms(c *gin.Context, req *GetDocAtomsRequest) *GetDocAtomsResponse {
	did, _ := service.ParseSid(req.DocId)
	atoms, err := model.GetAtomViewsByDocId(did)
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
