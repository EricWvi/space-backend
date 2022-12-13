package main

import (
	"github.com/gin-gonic/gin"
	"github.com/space-backend/handler"
	"github.com/space-backend/handler/editor"
	"github.com/space-backend/handler/files"
	"github.com/space-backend/handler/ping"
	"github.com/space-backend/middleware"
	"net/http"
)

const ApiBase = "/space"

// Load loads the middlewares, routes, handlers.
func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	// Middlewares.
	g.Use(gin.Recovery())
	g.Use(mw...)

	// 404 Handler.
	g.NoRoute(func(c *gin.Context) {
		handler.ReplyString(c, http.StatusNotFound, "The incorrect API route.")
	})

	g.Use(middleware.Logging)
	g.Any("/ping", ping.DefaultHandler)
	u := g.Group(ApiBase)
	u.Use(middleware.JWT)
	{
		u.Any("/files", files.DefaultHandler)
		u.Any("/editor", editor.DefaultHandler)
	}

	return g
}
