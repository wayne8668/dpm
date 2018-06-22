package common

import (
	"encoding/json"
	"net/http"
)

// Response struct
// swagger:response ResponseMsg
type responseMsg struct {
	//返回信息
	RspMsg interface{} `json:"rsp_msg"`
}

func newResponseMsg(msg string) *responseMsg {
	return &responseMsg{
		RspMsg: msg,
	}
}

func JsonResponseOK(w http.ResponseWriter, m interface{}) {
	JsonResponse(w, http.StatusOK, m)
}

func JsonResponseMsg(w http.ResponseWriter, httpStatus int, msg string) {
	JsonResponse(w, httpStatus, newResponseMsg(msg))
}

func JsonResponse(w http.ResponseWriter, httpStatus int, m interface{}) {
	w.Header().Set("Content-Type", "application/json;   charset=UTF-8")
	w.WriteHeader(httpStatus)
	if err := json.NewEncoder(w).Encode(m); err != nil {
		panic(err)
	}
}
