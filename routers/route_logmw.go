package routers

import (
	// "github.com/gorilla/mux"
	// "log"
	"dpm/common"
	"net/http"
	"time"
)

// func LoggerMiddleware(name string) mux.MiddlewareFunc {
// 	return func(next http.Handler) http.Handler {
// 		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
// 			start := time.Now()
// 			common.Logger.Infof(
// 				"%s\t%s\t%s\t%s",
// 				req.Method,
// 				req.RequestURI,
// 				name,
// 				time.Since(start),
// 			)
// 			next.ServeHTTP(w, req)
// 		})
// 	}
// }

func RouteLogMiddleware(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		start := time.Now()
		common.Logger.Infof(
			"%s\t%s\t%s\t%s",
			req.Method,
			req.RequestURI,
			name,
			time.Since(start),
		)
		inner.ServeHTTP(w, req)
	})
}
