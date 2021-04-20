package main

import (
	"design/global"
	"design/internal/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(global.GlbServer.Env)
	g := gin.New()
	routers.InitRouters(g)

	g.Run(global.GlbServer.AddressString())
}
