package repositories

import (
	"bytes"
	bolt "github.com/johnnadratowski/golang-neo4j-bolt-driver"
	"io"
)

type Cypher struct {
	cypher bytes.Buffer
	params map[string]interface{}
}

func NewCypher() Cypher {
	return Cypher{}
}

func (this Cypher) String() string {
	return this.cypher.String()
}

func (this Cypher) Params(args map[string]interface{}) Cypher {
	this.params = args
	return this
}

func (this Cypher) Match(s string) Cypher {
	this.cypher.WriteString(" match ")
	this.cypher.WriteString(s)
	this.cypher.WriteString(" ")
	return this
}

func (this Cypher) Where(s string) (r Cypher) {
	this.cypher.WriteString(" where ")
	this.cypher.WriteString(s)
	this.cypher.WriteString(" ")
	return this
}

func (this Cypher) Set(s string) (r Cypher) {
	this.cypher.WriteString(" set ")
	this.cypher.WriteString(s)
	this.cypher.WriteString(" ")
	return this
}

func (this Cypher) Create(s string) (r Cypher) {
	this.cypher.WriteString(" create ")
	this.cypher.WriteString(s)
	this.cypher.WriteString(" ")
	return this
}

func (this Cypher) Return(s string) (r Cypher) {
	this.cypher.WriteString(" return ")
	this.cypher.WriteString(s)
	this.cypher.WriteString(" ")
	return this
}

func (this Cypher) OrderBy(s string) (r Cypher) {
	this.cypher.WriteString(" order by ")
	this.cypher.WriteString(s)
	this.cypher.WriteString(" ")
	return this
}

func (this Cypher) Skip(s string) (r Cypher) {
	this.cypher.WriteString(" skip ")
	this.cypher.WriteString(s)
	this.cypher.WriteString(" ")
	return this
}

func (this Cypher) Limit(s string) (r Cypher) {
	this.cypher.WriteString(" limit ")
	this.cypher.WriteString(s)
	this.cypher.WriteString(" ")
	return this
}

func (this Cypher) Delete(s string) (r Cypher) {
	this.cypher.WriteString(" delete ")
	this.cypher.WriteString(s)
	this.cypher.WriteString(" ")
	return this
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func DoExecNeo(callback func(bolt.Conn) error) error {
	conn, err := pool.OpenPool()
	defer conn.Close()

	if err != nil {
		return err
	}

	tx, err := conn.Begin()

	if err != nil {
		return err
	}

	defer func() {
		if err := recover(); err != nil {
			tx.Rollback()
			panic(err)
		}
	}()

	err = callback(conn)

	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return err
}

func ExecNeo(cyphers ...Cypher) error {
	return DoExecNeo(func(conn bolt.Conn) (err error) {
		if len(cyphers) == 0 {
			cypher := cyphers[0]
			stmt, err := conn.PrepareNeo(cypher.String())
			defer stmt.Close()

			if err != nil {
				return err
			}
			_, err = stmt.ExecNeo(cypher.params)
		} else {
			queries := make([]string, len(cyphers))
			params := make([]map[string]interface{}, len(cyphers))

			for idx, cypher := range cyphers {
				queries[idx] = cypher.String()
				params[idx] = cypher.params
			}

			pstmt, err := conn.PreparePipeline(queries...)
			defer pstmt.Close()

			if err != nil {
				return err
			}
			_, err = pstmt.ExecPipeline(params...)
		}
		return err
	})
}

func QueryNeo(callback func([]interface{}), cypher Cypher) error {
	return DoExecNeo(func(conn bolt.Conn) error {

		cypherStr := cypher.String()

		Logger.Infof("Cypher string is -> [%s]", cypherStr)

		stmt, err := conn.PrepareNeo(cypherStr)
		defer stmt.Close()

		if err != nil {
			return err
		}

		r, err := stmt.QueryNeo(cypher.params)
		defer r.Close()

		if err != nil {
			return err
		}

		for {
			row, _, err := r.NextNeo()
			if err != nil || row == nil {
				if err == io.EOF {
					return nil
				}
				return err
			}
			callback(row)
		}
	})
}
