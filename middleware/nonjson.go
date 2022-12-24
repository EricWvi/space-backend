package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"time"
)

func NonJsonLogging(c *gin.Context) {
	start := time.Now().UTC()
	path := c.Request.URL.Path
	queries := c.Request.URL.RawQuery
	headers := c.Request.Header

	requestId := uuid.New().String()
	c.Set("RequestId", requestId)

	method := c.Request.Method
	ip := c.ClientIP()

	log.Infof("---------------------- %s ----------------------", requestId)
	log.WithFields(log.Fields{
		"requestId": requestId,
		"method":    method,
		"path":      path,
		"queries":   queries,
		"headers":   headers,
		"ip":        ip,
	}).Info()

	c.Next()

	// Calculates the latency.
	end := time.Now().UTC()
	latency := end.Sub(start)

	log.WithFields(log.Fields{
		"requestId": requestId,
		"latency":   latency,
	}).Info()
	log.Info("------------------------------------------------------------------")
}
