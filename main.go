package main

import (
	"dpm/common"
	"dpm/routers"
	"net/http"
)

func main() {
	router := routers.NewRouter()
	common.Logger.Infof("Golang Server will started with port:[%s]","8080")
	common.Logger.Fatal(http.ListenAndServe(":8080", router))
}
