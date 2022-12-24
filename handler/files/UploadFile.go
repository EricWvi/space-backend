package files

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/space-backend/config"
	"github.com/space-backend/handler"
	"github.com/space-backend/model"
	"github.com/space-backend/service"
	"io"
	"net/http"
)

func UploadFile(c *gin.Context) {
	name := c.Request.URL.Query().Get("Name")
	if name == "" {
		handler.ReplyError(c, http.StatusBadRequest, "arg[Name] is not set")
		return
	}
	l := c.Request.URL.Query().Get("Location")
	location := model.ParseFileLocation(l)
	if location == model.MalFileLocation {
		handler.ReplyError(c, http.StatusBadRequest, "arg[Location] is invalid")
		return
	}

	requestBody, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Error("failed to read request body file")
		handler.ReplyError(c, http.StatusInternalServerError, "failed to read request body")
		return
	}
	sid := service.NextId()
	link, err := service.UploadFile(location, sid.String(), requestBody)
	if err != nil {
		log.Error(err)
		handler.ReplyError(c, http.StatusInternalServerError, "failed to save file")
		return
	}

	fileType := c.Request.Header.Get("Content-Type")
	file := model.File{
		FileField: model.FileField{
			Sid:      sid,
			Name:     name,
			Size:     int64(len(requestBody)),
			Link:     link,
			Location: location,
			Type:     fileType,
		},
	}
	err = file.Create(config.DB)
	if err != nil {
		log.Error(err)
		handler.ReplyError(c, http.StatusInternalServerError, "failed to save file")
		return
	}

	handler.ReplyString(c, http.StatusOK, sid.String())
}
