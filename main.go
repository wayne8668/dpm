package main

import (
	"dpm/common"
	"dpm/routers"
	"dpm/vars"
	"net/http"
)

func main() {

	router := routers.NewRouter()
	httpPort := vars.Cfg.Get("server.http_port").(string)

	common.Logger.Infof("Golang Server will started with port:[%s]", httpPort)
	common.Logger.Fatal(http.ListenAndServe(httpPort, router))
}
