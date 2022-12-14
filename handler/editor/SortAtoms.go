package editor

import (
	"github.com/gin-gonic/gin"
	"github.com/space-backend/handler"
	"github.com/space-backend/model"
	"github.com/space-backend/service"
)

func (b Base) SortAtoms(c *gin.Context, req *SortAtomsRequest) *SortAtomsResponse {
	errList := ""
	for _, v := range req.Atoms {
		sid, _ := service.ParseSid(v.Sid)
		pid, _ := service.ParseSid(v.PrevId)
		err := model.UpdateAtomPrevId(sid, pid)
		if err != nil {
			errList += err.Error() + "\n"
		}
	}
	if errList != "" {
		handler.Errorf(c, errList)
		return nil
	}
	return &SortAtomsResponse{}
}

type SortAtomsRequest struct {
	Atoms []struct {
		Sid    string `json:"sid"`
		PrevId string `json:"prevId"`
	} `json:"atoms"`
}

type SortAtomsResponse struct {
}
