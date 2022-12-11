package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/space-backend/handler"
	"github.com/space-backend/service"
)

func JWT(c *gin.Context) {
	token := c.Request.Header.Get("X-API-Key")
	if len(token) == 0 {
		handler.Errorf(c, "token is missing")
		return
	}

	id, err := service.Parse(token)
	if err != nil {
		handler.Errorf(c, err.Error())
		return
	}
	if id != "eric" {
		handler.Errorf(c, "没有权限访问")
		return
	}
}
