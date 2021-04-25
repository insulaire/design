package routers

import (
	"design/internal/handles"
	"design/internal/middlewares"
	"fmt"

	"github.com/gin-gonic/gin"
)

func InitRouters(g *gin.Engine) {
	{
		group := g.Group("auth")
		limiter := middlewares.NewLimiter()
		group.Use(middlewares.NewMiddleware(limiter))
		group.GET("/download", handles.Download)
		LimitPost(limiter, group, "/gettoken", handles.GenerateToken())

	}
	g.Use(middlewares.NewMiddleware(middlewares.NewLimiter()))
	g.Use(middlewares.GWT)
	g.POST("/ping", handles.Ping())
	g.POST("/getpath", handles.GetPath())
}

func LimitPost(limiter *middlewares.Limiter, group *gin.RouterGroup, relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	limiter.Add(fmt.Sprintf("%s%s", group.BasePath(), relativePath))
	return group.POST(relativePath, handlers...)
}
