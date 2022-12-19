package editor

import (
	"github.com/gin-gonic/gin"
	"github.com/space-backend/config"
	"github.com/space-backend/handler"
	"github.com/space-backend/model"
)

func (b Base) GetCollectionDocs(c *gin.Context, req *GetCollectionDocsRequest) *GetCollectionDocsResponse {
	views, err := model.GetDocViewsByCollectionId(config.DB, req.CollectionId)
	if err != nil {
		handler.Errorf(c, err.Error())
		return nil
	}

	return &GetCollectionDocsResponse{
		Docs: views,
	}
}

type GetCollectionDocsRequest struct {
	CollectionId model.Sid `json:"collectionId"`
}

type GetCollectionDocsResponse struct {
	Docs []*model.DocView `json:"docs"`
}
