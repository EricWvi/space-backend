package editor

import (
	"github.com/gin-gonic/gin"
	"github.com/space-backend/config"
	"github.com/space-backend/handler"
	"github.com/space-backend/model"
	"github.com/space-backend/service"
)

func (b Base) GetCollectionDocs(c *gin.Context, req *GetCollectionDocsRequest) *GetCollectionDocsResponse {
	cid, err := service.ParseSid(req.CollectionId)
	var views []*model.DocView
	if err == nil {
		views, err = model.GetDocViewsByCollectionId(config.DB, cid)
	}
	if err != nil {
		handler.Errorf(c, err.Error())
		return nil
	}

	return &GetCollectionDocsResponse{
		Docs: views,
	}
}

type GetCollectionDocsRequest struct {
	CollectionId string `json:"collectionId"`
}

type GetCollectionDocsResponse struct {
	Docs []*model.DocView `json:"docs"`
}
