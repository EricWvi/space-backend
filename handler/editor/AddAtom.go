package editor

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/space-backend/handler"
	"github.com/space-backend/model"
	"github.com/space-backend/service"
)

func (b Base) AddAtom(c *gin.Context, req *AddAtomRequest) *AddAtomResponse {
	docId, err := service.ParseSid(req.DocId)
	if err != nil {
		handler.Errorf(c, err.Error())
		return nil
	}
	a := model.Atom{
		Sid:     service.NextId(),
		Content: req.Content,
		Link:    req.Link,
		Type:    model.FormatAtomType(req.Type),
		DocId:   docId,
	}
	err = a.Create()
	var sid string
	if err == nil {
		sid, err = service.ToSid(a.Sid)
	}
	if err != nil {
		log.Error(err)
		handler.Errorf(c, "failed to add atom")
		return nil
	}
	return &AddAtomResponse{
		Sid: sid,
	}
}

type AddAtomRequest struct {
	Content string `json:"content"`
	Link    string `json:"link"`
	Type    string `json:"type"`
	DocId   string `json:"docId"`
}

type AddAtomResponse struct {
	Sid string `json:"sid"`
}
