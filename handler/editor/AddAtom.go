package editor

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/space-backend/handler"
	"github.com/space-backend/model"
	"github.com/space-backend/service"
)

func (b Base) AddAtom(c *gin.Context, req *AddAtomRequest) *AddAtomResponse {
	docId, _ := service.ParseSid(req.DocId)
	prevId, _ := service.ParseSid(req.PrevId)

	var sidRaw int64
	if req.Sid != "" {
		sidRaw, _ = service.ParseSid(req.Sid)
	} else {
		sidRaw = service.NextId()
	}

	a := model.Atom{
		AtomField: model.AtomField{
			Sid:     sidRaw,
			Content: req.Content,
			Link:    req.Link,
			Type:    model.ParseAtomType(req.Type),
			DocId:   docId,
			PrevId:  prevId,
		},
	}

	err := a.Create()
	if err != nil {
		log.Error(err)
		handler.Errorf(c, "failed to add atom")
		return nil
	}
	sid, _ := service.ToSid(a.Sid)
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
	PrevId  string `json:"prevId"`
}

type AddAtomResponse struct {
	Sid string `json:"sid"`
}
