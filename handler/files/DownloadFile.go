package files

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/space-backend/config"
	"github.com/space-backend/handler"
	"github.com/space-backend/model"
	"github.com/space-backend/service"
	"net/http"
)

func DownloadFile(c *gin.Context) {
	i := c.Request.URL.Query().Get("Sid")
	sid, err := model.ParseSid(i)
	if err != nil || sid == 0 {
		handler.ReplyError(c, http.StatusBadRequest, "arg[Sid] is invalid")
		return
	}
	file, err := model.GetFile(config.DB, map[string]any{model.File_Sid: sid})
	content, err := service.DownloadFile(file.Location, file.Link)
	if err != nil {
		log.Error(err)
		handler.ReplyError(c, http.StatusInternalServerError, "failed to download file")
		return
	}

	c.Data(http.StatusOK, file.Type, content)
}
