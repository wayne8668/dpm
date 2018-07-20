package common

import (
	"io"
	"io/ioutil"
	"encoding/json"
	"net/http"
	"dpm/vars"
)

//反序列化http.Request.body至Object(传址)
func Unmarshal2Object(r *http.Request, obj interface{}) error {
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, vars.Cfg.Get("request.limit_byte").(int64)))
	if err != nil {
		panic(ErrBadRequest(err.Error()))
	}
	if err := r.Body.Close(); err != nil {
		panic(ErrInternalServer(err))
	}

	// Logger.Debug("http request boy json:",string(body))

	return json.Unmarshal(body, obj)
}
