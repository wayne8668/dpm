package common

import (
	"strconv"
	"github.com/sony/sonyflake"
	"time"
)

var (
	sf *sonyflake.Sonyflake
)

func MachineIdFunc() (uint16, error) {
	//uint16 max value is 65535 in 64bit machine
	return 10101, nil
}

func CheckMachineIdFunc(machineId uint16) bool {
	//machineId is the MachineIdFunc's return params 'uint64' value
	return machineId == 10101
}

func init() {
	layout := "2006-01-02 15:04:05.999"
	st := "2016-06-06 06:06:06"
	t, err := time.ParseInLocation(layout, st, time.Local)
	if err != nil {
		panic(err)
	}
	s := sonyflake.Settings{
		StartTime:      t,
		MachineID:      MachineIdFunc,
		CheckMachineID: CheckMachineIdFunc,
	}
	sf = sonyflake.NewSonyflake(s)
}

func UUID() string {
	uint64Id, err := sf.NextID()
	if err != nil {
		panic(err)
	}
	return strconv.FormatUint(uint64Id,10)
}
