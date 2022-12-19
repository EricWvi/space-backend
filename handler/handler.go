package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"reflect"
)

type Response struct {
	RequestId string `json:"requestId"`
	Code      int    `json:"code"`
	Message   any    `json:"message"`
}

func Dispatch(c *gin.Context, base any) {
	body := c.GetString("RequestBody")

	method := reflect.ValueOf(base).MethodByName(c.GetString("Action"))
	if !method.IsValid() {
		ReplyError(c, http.StatusNotFound, "request action does not exist")
	} else {
		Type := method.Type()
		ctx := reflect.ValueOf(c)
		param := Type.In(1).Elem()
		ptr, err := parse([]byte(body), param)
		if err != nil {
			ReplyError(c, http.StatusBadRequest, "failed to parse request body: "+err.Error())
			return
		}
		log.Debugf("%#v", ptr.Elem())
		rst := method.Call([]reflect.Value{ctx, ptr})[0]
		if !c.IsAborted() {
			c.JSON(http.StatusOK, Response{
				RequestId: c.GetString("RequestId"),
				Code:      http.StatusOK,
				Message:   rst.Interface(),
			})
		}
	}
}

func ReplyError(c *gin.Context, code int, msg string) {
	c.JSON(code, Response{
		RequestId: c.GetString("RequestId"),
		Code:      code,
		Message:   msg,
	})
}

func ReplyString(c *gin.Context, code int, msg string) {
	c.JSON(http.StatusOK, Response{
		RequestId: c.GetString("RequestId"),
		Code:      code,
		Message:   msg,
	})
}

func Errorf(c *gin.Context, format string, a ...any) {
	ReplyString(c, http.StatusBadRequest, fmt.Sprintf(format, a...))
	c.Abort()
}

// TODO validator
func parse(body []byte, param reflect.Type) (reflect.Value, error) {
	ptr := reflect.New(param).Interface()
	if len(body) != 0 {
		err := json.Unmarshal(body, &ptr)
		if err != nil {
			log.Error(err)
			return reflect.Value{}, err
		}
	}
	return reflect.ValueOf(ptr), nil
}
