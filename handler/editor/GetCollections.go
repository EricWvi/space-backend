package editor

import (
	"github.com/gin-gonic/gin"
	"github.com/space-backend/config"
	"github.com/space-backend/handler"
	"github.com/space-backend/model"
)

func (b Base) GetCollections(c *gin.Context, req *GetCollectionsRequest) *GetCollectionsResponse {
	views, err := model.GetCollectionViews(config.DB)
	if err != nil {
		handler.Errorf(c, err.Error())
		return nil
	}

	return &GetCollectionsResponse{
		Collections: views,
	}
}

type GetCollectionsRequest struct {
}

type GetCollectionsResponse struct {
	Collections []*model.CollectionView `json:"collections"`
}
