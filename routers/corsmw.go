package routers

import (
	"strings"
	// "encoding/json"
	"dpm/common"
	// "github.com/gorilla/mux"
	"net/http"
)

// func CORSMethodMiddleware() mux.MiddlewareFunc {
// 	return func(next http.Handler) http.Handler {
// 		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
// 			common.Logger.Infof("CORS Middleware is working...")
// 			w.Header().Set("Access-Control-Allow-Origin", "*")
// 			w.Header().Set("Access-Control-Allow-Headers", "content-type")
// 			// w.Header().Set("Content-Type", "application/json;   charset=UTF-8")
// 			// w.Header().Set("Access-Control-Allow-Methods", "POST,OPTIONS")
// 			if req.Method == "OPTIONS" {
// 				w.WriteHeader(http.StatusOK)
// 				return
// 			}
// 			next.ServeHTTP(w, req)
// 		})
// 	}
// }

func CORSAllowMiddleware(m []string, inner http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		common.Logger.Infof("CORS Middleware is working...")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "authorization,content-type")
		w.Header().Set("Content-Type", "application/json;   charset=UTF-8")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Methods", strings.Join(m, ","))
		if req.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		inner.ServeHTTP(w, req)
	})
}
