package api

import (
	"dpm/common"
	"dpm/middleware"
)

var (
	Logger       = common.Logger
	CURRENT_USER = middleware.CURRENT_USER
)

const (
// API_KEY   = vars.PROJECT_NAME
)

///////////////////////////////////////////////////////////////////////////////

//PageableRequest
type PageableRequest struct {
	Limit int64 `query:"limit" binding:"required,numeric"`
	Page  int64 `query:"page" binding:"required,numeric"`
}
