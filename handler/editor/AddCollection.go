package editor

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
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
	// TODO check if the name of collection exists
	err := coll.Create()
	if err != nil {
		log.Error(err)
		handler.Errorf(c, "failed to add collection")
		return nil
	}
	sid, _ := service.ToSid(coll.Sid)
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
