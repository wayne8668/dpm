package models

import (
	"dpm/common"
)

type Project struct {
	Id         string          `json:"id"`
	Name       string          `json:"name"`
	Desc       string          `json:"desc"`
	Visibility bool            `json:"visibility"`
	CreateTime common.JSONTime `json:"create_time"`
}
