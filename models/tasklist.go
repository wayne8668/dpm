package models

import (
	"time"
)

type TaskList struct {
	TLid   string
	Name   string
	Order  int
	Cyclic int
}

type Task struct {
	Tid        string
	Name       string
	Desc       string
	CreateTime time.Time
	FinishTime time.Time
	DateLine   time.Time
	RemindTime time.Time
}
