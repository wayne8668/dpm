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
	Id         string          `json:"id"`
	Name       string          `json:"name"`
	Pwd        string          `json:"pwd"`
	CreateTime common.JSONTime `json:"create_time"`
}
