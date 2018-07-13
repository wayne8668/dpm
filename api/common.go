package api

import (
	"strings"
	"net/url"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"dpm/vars"
	// "strings"
	"dpm/common"
	// "io"
	// "io/ioutil"
	// "net/url"
	// "github.com/gorilla/mux"
	// "encoding/json"
	// "fmt"
	// "net/http"
)

var (
	Logger         = common.Logger
	jsonResponse   = common.JsonResponse
	jsonResponseOK = common.JsonResponseOK
	CURRENT_USER   = common.CURRENT_USER
)

const (
	API_KEY   = vars.PROJECT_NAME
)

func ParseQueryGet(r *http.Request, key string) string {
	vs, err := url.ParseQuery(r.URL.RawQuery)
	if err == nil && len(vs[key]) > 0 {
		return strings.TrimSpace(vs[key][0])
	}
	return ""
}

//反序列化http.Request.body至Object(传址)
func Unmarshal2Object(w http.ResponseWriter, r *http.Request, obj interface{}) {
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	// Logger.Debug("http request boy json:",string(body))

	if err := json.Unmarshal(body, obj); err != nil {
		m := map[string]interface{}{
			"err": "err",
		}
		jsonResponse(w, http.StatusUnprocessableEntity, m)
	}
}

// func unmarshal2JSON(w http.ResponseWriter, r *http.Request) (httpJson string) {
// 	body, err := ioutil.ReadAll(io.LimitReader(r.Body, LimitByte))
// 	if err != nil {
// 		m := map[string]interface{}{
// 			"err": "err",
// 		}
// 		jsonResponse(w, http.StatusUnprocessableEntity, m)
// 		panic(err)
// 	}
// 	if err := r.Body.Close(); err != nil {
// 		panic(err)
// 	}
// 	fmt.Println("req_json:", string(body))
// 	return string(body)
// }
