package repositories

import (
	"bytes"
	bolt "github.com/johnnadratowski/golang-neo4j-bolt-driver"
)

type Cypher struct {
	cypher bytes.Buffer
}

func NewCypher() Cypher {
	return Cypher{}
}

func (this Cypher) String() string {
	return this.cypher.String()
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

type Neo4GoTemplate struct {
	cypher string
}

func NewNeo4GoTemplate(c Cypher) *Neo4GoTemplate {
	t := &Neo4GoTemplate{}
	t.cypher = c.String()
	return t
}

func (this *Neo4GoTemplate) execute(params map[string]interface{}, callback func(bolt.Stmt) error) error {
	conn := GetConn()
	defer conn.Close()

	tx, err := conn.Begin()

	if err != nil {
		return err
	}

	stmt, err := conn.PrepareNeo(this.cypher)
	defer stmt.Close()

	if err != nil {
		tx.Rollback()
		return err
	}

	err = callback(stmt)

	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return err
}

func (this *Neo4GoTemplate) ExecNeo(params map[string]interface{}) error {
	return this.execute(params, func(stmt bolt.Stmt) error {
		_, err := stmt.ExecNeo(params)
		return err
	})
}

func (this *Neo4GoTemplate) QueryNeoNextOne(params map[string]interface{}, callback func([]interface{})) error {
	return this.execute(params, func(stmt bolt.Stmt) error {
		rows, err := stmt.QueryNeo(params)
		defer rows.Close()
		if err != nil {
			return err
		}
		r, _, err := rows.NextNeo()
		if err != nil {
			return err
		}
		callback(r)
		return err
	})
}

func (this *Neo4GoTemplate) QueryNeoAll(params map[string]interface{}, callback func([]interface{})) error {
	return this.execute(params, func(stmt bolt.Stmt) error {
		rows, err := stmt.QueryNeo(params)
		defer rows.Close()
		if err != nil {
			return err
		}
		r, _, err := rows.All()
		if err != nil {
			return err
		}
		for _, row := range r {
			callback(row)
		}
		return err
	})
}
