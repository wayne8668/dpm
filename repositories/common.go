package repositories

import (
	"dpm/common"
	"github.com/satori/go.uuid"
	// "github.com/goinggo/mapstructure"
)

var (
	Logger = common.Logger
)

func NewUUID() string {
	return uuid.Must(uuid.NewV4()).String()
}
