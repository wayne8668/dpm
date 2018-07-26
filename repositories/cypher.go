package repositories

import (
	"dpm/vars"
	bolt "github.com/johnnadratowski/golang-neo4j-bolt-driver"
	"io"
)

type Cypher struct {
	cypStr string
	params map[string]interface{}
}

func NewCypher(key string) Cypher {
	Logger.Infof("load cypher.ini key:[%s]", key)
	return Cypher{
		// cypStr :strings.Replace(vars.CypherCfg.Get(key).(string),"\n", "", -1),
		cypStr: vars.CypherCfg.Get(key).(string),
	}
}

func (this Cypher) Params(args map[string]interface{}) Cypher {
	this.params = args
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

	if err = callback(conn); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func ExecNeo(cyphers ...Cypher) error {
	return DoExecNeo(func(conn bolt.Conn) (err error) {
		// if len(cyphers) == 0 {
		// 	cypher := cyphers[0]
		// 	cypherStr := cypher.String()
		// 	Logger.Infof("Cypher string is: [%s]", cypherStr)
		// 	Logger.Info("Cypher params is: ", cypher.params)
		// 	stmt, err := conn.PrepareNeo(cypherStr)
		// 	defer stmt.Close()

		// 	if err != nil {
		// 		return err
		// 	}
		// 	_, err = stmt.ExecNeo(cypher.params)
		// } else {
		queries := make([]string, len(cyphers))
		params := make([]map[string]interface{}, len(cyphers))

		for idx, cypher := range cyphers {
			queries[idx] = cypher.cypStr
			params[idx] = cypher.params
			Logger.Infof("Cypher string is: [%s]", cypher.cypStr)
			Logger.Info("Cypher params is: ", cypher.params)
		}

		pstmt, err := conn.PreparePipeline(queries...)
		defer pstmt.Close()

		if err != nil {
			return err
		}
		_, err = pstmt.ExecPipeline(params...)
		// }
		return err
	})
}

func QueryNeo(rowCallBack func([]interface{}), cypher Cypher) error {
	return DoExecNeo(func(conn bolt.Conn) error {

		cypherStr := cypher.cypStr

		Logger.Infof("Cypher string is: [%s]", cypherStr)

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
			rowCallBack(row)
		}
	})
}
