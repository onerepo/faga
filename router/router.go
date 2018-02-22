package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/onerepo/faga/router/middleware/header"
	"github.com/onerepo/faga/server"
)

func Load(middleware ...gin.HandlerFunc) http.Handler {
	e := gin.New()
	e.Use(gin.Recovery())

	e.Use(header.NoCache)
	e.Use(header.Options)
	e.Use(header.Secure)

	e.GET("/version", server.GetVersion)
	return e
}
