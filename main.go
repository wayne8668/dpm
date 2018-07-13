package main

import (
	"dpm/common"
	"dpm/routers"
	"dpm/vars"
	"github.com/pelletier/go-toml"
	"net/http"
	"path/filepath"
)

func LoadCfg() {

	filePath, err := filepath.Abs(vars.APP_CFG_PATH)

	if err != nil {
		panic(err)
	}

	vars.Cfg, err = toml.LoadFile(filePath)

	if err != nil {
		panic(err)
	}
}

func LoadCypherCfg() {

	filePath, err := filepath.Abs(vars.CYPHER_PATH)

	if err != nil {
		panic(err)
	}

	vars.CypherCfg, err = toml.LoadFile(filePath)

	if err != nil {
		panic(err)
	}
}

func init() {

	LoadCfg()

	LoadCypherCfg()
}

func main() {

	router := routers.NewRouter()
	httpPort := vars.Cfg.Get("server.http_port").(string)

	common.Logger.Infof("Golang Server will started with port:[%s]", httpPort)
	common.Logger.Fatal(http.ListenAndServe(httpPort, router))
}
