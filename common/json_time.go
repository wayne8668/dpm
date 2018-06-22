package common

import (
	"time"
)

type (
	JSONTime time.Time
)

const (
	timeFormart = "2006-01-02 15:04:05"
)

func (this *JSONTime) UnmarshalJSON(data []byte) (err error) {
	t, err := time.ParseInLocation(`"`+timeFormart+`"`, string(data), time.Local)
	*this = JSONTime(t)
	return
}

func (this JSONTime) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(timeFormart)+2)
	b = append(b, '"')
	b = time.Time(this).AppendFormat(b, timeFormart)
	b = append(b, '"')
	return b, nil
}

func UMStr2JSONTime(st string) (JSONTime, error) {
	t, err := time.ParseInLocation(`"`+timeFormart+`"`, st, time.Local)
	return JSONTime(t), err
}

func NowStringFormat() string {
	return time.Now().Format(timeFormart)
}

func (this JSONTime) String() string {
	return time.Time(this).Format(timeFormart)
}
