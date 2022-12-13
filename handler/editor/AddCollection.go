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
		Sid:  service.NextId(),
		Name: req.Name,
	}
	err := coll.Create()
	var sid string
	if err == nil {
		sid, err = service.ToSid(coll.Sid)
	}
	if err != nil {
		log.Error(err)
		handler.Errorf(c, "failed to add collection")
		return nil
	}
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
