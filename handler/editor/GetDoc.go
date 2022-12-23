package editor

import (
	"github.com/gin-gonic/gin"
	"github.com/space-backend/config"
	"github.com/space-backend/handler"
	"github.com/space-backend/model"
)

func (b Base) GetDoc(c *gin.Context, req *GetDocRequest) (rsp *GetDocResponse) {
	doc, err := model.GetDoc(config.DB, map[string]any{model.Doc_Sid: req.DocId})
	if err != nil {
		handler.Errorf(c, err.Error())
		return nil
	}
	return &GetDocResponse{
		Title:   doc.Title,
		Version: doc.Version,
	}
}

type GetDocRequest struct {
	DocId model.Sid `json:"docId"`
}

type GetDocResponse struct {
	Title   string `json:"title"`
	Version int    `json:"version"`
}
