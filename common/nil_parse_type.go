package common

import (
	"time"
)

func NilParseString(arg interface{}) (s string){
	if arg == nil {
		return s
	}
	return arg.(string)
}

func NilParseInt(arg interface{}) (i int){
	if arg == nil {
		return i
	}
	return arg.(int)
}

func NilParseJSONTime(arg interface{}) (j JSONTime){
	if arg == nil {
		return j
	}
	return JSONTime(time.Unix(0,arg.(int64)))
}