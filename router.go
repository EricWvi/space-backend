package main

import (
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/jxskiss/ginregex"
	log "github.com/sirupsen/logrus"
	"github.com/space-backend/handler"
	"github.com/space-backend/handler/editor"
	"github.com/space-backend/handler/files"
	"github.com/space-backend/handler/login"
	"github.com/space-backend/handler/ping"
	"github.com/space-backend/middleware"
	"github.com/spf13/viper"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// Load loads the middlewares, routes, handlers.
func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	// Middlewares.
	g.Use(gin.Recovery())
	g.Use(mw...)
	g.Use(gzip.Gzip(gzip.DefaultCompression))

	// 404 Handler.
	g.NoRoute(func(c *gin.Context) {
		handler.ReplyError(c, http.StatusNotFound, "The incorrect API route.")
	})

	regexRouter := ginregex.New(g, nil)
	regexRouter.Any("^/.*$", func(c *gin.Context) {
		c.File(viper.GetString("route.front.index"))
	})

	dir := viper.GetString("route.front.dir")
	err := filepath.Walk(dir,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() {
				p := strings.TrimPrefix(path, dir)
				g.StaticFile(p, path)
			}
			return nil
		})
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}

	back := g.Group(viper.GetString("route.back.base"))

	back.Use(middleware.Logging)
	back.Any("/ping", ping.DefaultHandler)
	back.Any("/login", login.DefaultHandler)
	auth := back.Group("space")
	auth.Use(middleware.JWT)
	{
		auth.Any("/files", files.DefaultHandler)
		auth.Any("/editor", editor.DefaultHandler)
	}

	return g
}
