package ping

import (
	"github.com/gin-gonic/gin"
	"github.com/space-backend/handler"
)

type Base struct{}

func DefaultHandler(c *gin.Context) {
	handler.Dispatch(c, Base{})
}

func (b Base) Ping(c *gin.Context, req *PingRequest) *PingResponse {
	return &PingResponse{
		Value: "Pong",
		Echo:  req.Echo,
	}
}

type PingRequest struct {
	Echo string `json:"echo"`
}

type PingResponse struct {
	Value string `json:"value"`
	Echo  string `json:"echo"`
}
