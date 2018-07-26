package main

import (
	"runtime"
	"dpm/routers"
	"dpm/vars"
	"github.com/gin-gonic/gin"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU() * 2)
	gin.SetMode(gin.DebugMode)
	router := gin.New()

	g := router.Group("/dpm/api/v1.0")

	routers.RegisterMiddleWare(g)
	routers.RegisterRouter(g)

	httpPort := vars.Cfg.Get("server.http_port").(string)

	router.Run(httpPort)
}
