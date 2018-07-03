package vars

import (
	"github.com/pelletier/go-toml"
)

const (
	CYPHER_PATH = "./conf/cypher.ini"

	APP_CFG_PATH = "./conf/config.ini"
)

var (
	Cfg, CypherCfg *toml.Tree
)
