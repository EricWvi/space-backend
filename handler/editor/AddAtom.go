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

	var sidRaw int64
	if req.Sid != "" {
		sidRaw, err = service.ParseSid(req.Sid)
		if err != nil {
			handler.Errorf(c, err.Error())
			return nil
		}
	} else {
		sidRaw = service.NextId()
	}
	
	a := model.Atom{
		Sid:     sidRaw,
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
	Sid     string `json:"sid"`
	Link    string `json:"link"`
	Type    string `json:"type"`
	DocId   string `json:"docId"`
}

type AddAtomResponse struct {
	Sid string `json:"sid"`
}
