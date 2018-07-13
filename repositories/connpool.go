package repositories

import (
	bolt "github.com/johnnadratowski/golang-neo4j-bolt-driver"
)

const (
	connStr = "bolt://neo4j:xhtian@127.0.0.1:7687"
)

var (
	pool    bolt.ClosableDriverPool
)

func init() {
	var err error
	pool, err = bolt.NewClosableDriverPool(connStr, 10)
	if err != nil {
		Logger.Fatalf("cann't create the neo4j conn pool with url:[%s]", connStr)
	}
}
