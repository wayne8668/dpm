package api

import (
	"dpm/common"
)

var (
	Logger         = common.Logger
	CURRENT_USER   = common.CURRENT_USER
)

const (
	// API_KEY   = vars.PROJECT_NAME
)

type PageableRequest struct{
	Limit int64 `qval:"limit,inquery"`
	Page int64	`qval:"page,inquery"`
}