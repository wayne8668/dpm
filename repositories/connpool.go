package repositories

import (
	bolt "github.com/johnnadratowski/golang-neo4j-bolt-driver"
)

const (
	connStr = "bolt://neo4j:xhtian@localhost:7687"
)

var (
	pool    bolt.ClosableDriverPool
	baseRep = &BaseRepository{}
)

func init() {
	var err error
	pool, err = bolt.NewClosableDriverPool(connStr, 10)
	if err != nil {
		Logger.Fatalf("cann't create the neo4j conn pool with url:[%s]", connStr)
	}
}
