package login

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/space-backend/handler"
	"github.com/space-backend/service"
)

type Base struct{}

func DefaultHandler(c *gin.Context) {
	handler.Dispatch(c, Base{})
}

func (b Base) LogIn(c *gin.Context, req *LogInRequest) *LogInResponse {
	if req.Username != "Eric" || req.Password != "1010521" {
		handler.Errorf(c, "用户名或密码错误")
		return nil
	}

	tokenString, err := service.Sign(1010)
	if err != nil {
		log.Error(err)
		handler.Errorf(c, "failed to sign in")
		return nil
	}

	return &LogInResponse{
		Token: tokenString,
	}
}

type LogInRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LogInResponse struct {
	Token string `json:"token"`
}
