package editor

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/space-backend/config"
	"github.com/space-backend/handler"
	"github.com/space-backend/model"
	"github.com/space-backend/service"
)

func (b Base) AddCollection(c *gin.Context, req *AddCollectionRequest) *AddCollectionResponse {
	coll := model.Collection{
		CollectionField: model.CollectionField{
			Sid:  service.NextId(),
			Name: req.Name,
		},
	}
	_, err := model.GetCollection(config.DB, map[string]any{model.Collection_Name: req.Name})
	if err == nil {
		handler.Errorf(c, "Collection %s already exists", req.Name)
		return nil
	}
	err = nil
	err = coll.Create(config.DB)
	if err != nil {
		log.Error(err)
		handler.Errorf(c, "failed to add collection")
		return nil
	}
	sid := coll.Sid.String()
	return &AddCollectionResponse{
		Sid: sid,
	}
}

type AddCollectionRequest struct {
	Name string `json:"name"`
}

type AddCollectionResponse struct {
	Sid string `json:"sid"`
}
