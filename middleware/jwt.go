package middleware

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/space-backend/handler"
	"github.com/space-backend/service"
	"net/http"
)

func JWT(c *gin.Context) {
	token := c.Request.Header.Get("X-API-Key")
	if len(token) == 0 {
		handler.ReplyError(c, http.StatusBadRequest, "token is missing")
		c.Abort()
		return
	}

	id, err := service.ParseToken(token)
	if err != nil {
		log.Error(err)
		handler.ReplyError(c, http.StatusInternalServerError, "failed to parse token")
		c.Abort()
		return
	}
	if id != 1010 {
		handler.ReplyError(c, http.StatusBadRequest, "没有权限访问")
		c.Abort()
		return
	}
}
