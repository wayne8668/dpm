package repositories

import (
	"fmt"
)

type BaseRepository struct{}

func (this *BaseRepository) execNeo(connStr string, m map[string]interface{}) (numResult int64, err error) {
	conn := GetConn()
	defer conn.Close()
	result, err := conn.ExecNeo(connStr, m)
	numResult, err = result.RowsAffected()
	fmt.Println("BaseRepository ExecNeo method invork...")
	return numResult, err
}
