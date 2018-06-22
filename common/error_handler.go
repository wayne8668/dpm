package common

import (
	"runtime/debug"
	"github.com/juju/errors"
	"net/http"
)

func AppErrorHandlerFunc(inner http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				var errMsg string
				var httpCode int = http.StatusInternalServerError
				if oe, ok := err.(error); ok {
					errMsg = errors.ErrorStack(oe)
					if ae, ok := err.(*AppError); ok {
						errMsg = errors.ErrorStack(ae.error)
						httpCode = ae.HttpStatusCode
					}
					Logger.Error(errMsg)
					debug.PrintStack()
					JsonResponseMsg(w, httpCode, oe.Error())
				}
			}
		}()
		inner.ServeHTTP(w, r)
	})
}
