package editor

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/space-backend/handler"
	"github.com/space-backend/model"
	"github.com/space-backend/service"
)

func (b Base) AddDoc(c *gin.Context, req *AddDocRequest) *AddDocResponse {
	collId, err := service.ParseSid(req.CollectionId)
	if err != nil {
		handler.Errorf(c, err.Error())
		return nil
	}
	doc := model.Doc{
		Sid:          service.NextId(),
		CollectionId: collId,
		Title:        req.Title,
	}
	err = doc.Create()

	var sid string
	if err == nil {
		sid, err = service.ToSid(doc.Sid)
	}
	if err != nil {
		log.Error(err)
		handler.Errorf(c, "failed to add doc")
		return nil
	}
	return &AddDocResponse{
		Sid: sid,
	}
}

type AddDocRequest struct {
	Title        string `json:"title"`
	CollectionId string `json:"collectionId"`
}

type AddDocResponse struct {
	Sid string `json:"sid"`
}
