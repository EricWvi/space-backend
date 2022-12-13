package editor

import (
	"github.com/gin-gonic/gin"
	"github.com/space-backend/handler"
)

type Base struct{}

func DefaultHandler(c *gin.Context) {
	handler.Dispatch(c, Base{})
}
