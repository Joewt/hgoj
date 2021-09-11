package router

import "github.com/gin-gonic/gin"

func ApiRouter(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	g.Use(gin.Recovery())
	g.Use(mw...)

	api := g.Group("/api")
	{
		v1 := api.Group("v1")
		v1.POST("login")
	}

	return g
}
