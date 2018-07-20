package api

import (
	"dpm/common"
)

var (
	Logger       = common.Logger
	CURRENT_USER = common.CURRENT_USER
)

const (
// API_KEY   = vars.PROJECT_NAME
)

///////////////////////////////////////////////////////////////////////////////

//PageableRequest
type PageableRequest struct {
	Limit int64 `qval:"limit,inquery"`
	Page  int64 `qval:"page,inquery"`
}

///////////////////////////////////////////////////////////////////////////////

//RspModel
type RspModel map[string]interface{}

func NewRspModel() RspModel {
	return make(map[string]interface{})
}

func (m RspModel) AddAttribute(k string, v interface{}) RspModel {
	if m == nil {
		m = make(map[string]interface{})
	}
	m[k] = v
	return m
}

///////////////////////////////////////////////////////////////////////////////
