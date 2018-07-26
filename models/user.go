package models

import (
	"dpm/common"
)

// User
//
// User Entity
//
// swagger:model User
type User struct {
	UId        string           `json:"uid"`
	Name       string          `json:"name" binding:"required"`
	Pwd        string          `json:"pwd" binding:"required"`
	CreateTime common.JSONTime `json:"create_time"`
}
