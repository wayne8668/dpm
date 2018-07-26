package middleware

import (
	"dpm/common"
	"github.com/gin-gonic/gin"
	"github.com/juju/errors"
	"net/http"
	"runtime/debug"
)

func Recovery() gin.HandlerFunc {
	return func(cxt *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				var errMsg string
				var httpCode int = http.StatusInternalServerError
				if oe, ok := err.(error); ok {
					errMsg = errors.ErrorStack(oe)
					if ae, ok := err.(*common.AppError); ok {
						errMsg = errors.ErrorStack(ae.GetError())
						httpCode = ae.HttpStatusCode
					}
					common.Logger.Error(errMsg)
					debug.PrintStack()
					cxt.JSON(httpCode, common.RspMsg(oe.Error()))
				}
			}
		}()
		cxt.Next()
	}
}
