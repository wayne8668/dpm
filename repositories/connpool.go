package repositories

import (
	bolt "github.com/johnnadratowski/golang-neo4j-bolt-driver"
	"log"
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
		log.Fatalf("don't create the neo4j conn pool with url:[%s]", connStr)
	}
}

func GetConn() bolt.Conn {
	conn, _ := pool.OpenPool()
	return conn
}
