package vars

import (
	"path/filepath"
	"github.com/pelletier/go-toml"
)

const (

	PROJECT_NAME = "dpm"

	CYPHER_PATH = "conf/cypher.ini"

	APP_CFG_PATH = "conf/config.ini"
)

var (
	Cfg, CypherCfg *toml.Tree
)

func LoadCfg() {

	filePath, err := filepath.Abs(APP_CFG_PATH)

	if err != nil {
		panic(err)
	}

	Cfg, err = toml.LoadFile(filePath)

	if err != nil {
		panic(err)
	}
}

func LoadCypherCfg() {

	filePath, err := filepath.Abs(CYPHER_PATH)

	if err != nil {
		panic(err)
	}

	CypherCfg, err = toml.LoadFile(filePath)

	if err != nil {
		panic(err)
	}
}

func init() {

	LoadCfg()

	LoadCypherCfg()
}