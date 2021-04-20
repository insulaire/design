package routers

import (
	"design/internal/handles"
	"design/internal/middlewares"

	"github.com/gin-gonic/gin"
)

func InitRouters(g *gin.Engine) {
	{
		group := g.Group("auth")
		group.POST("/gettoken", handles.GenerateToken())
	}
	g.Use(middlewares.GWT)
	g.POST("/ping", handles.Ping())
	g.POST("/getpath", handles.GetPath())
}
