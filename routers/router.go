package routers

import (
	"dpm/common"
	"github.com/gorilla/mux"
	// "net/http"
)

type Route struct {
	Name        string
	Methods     []string
	Pattern     string
	HandlerFunc interface{}
}

type (
	ControllerRouter []Route
)

const (
	CONTEXT_PATH = "/dpm"
	VERSION      = "/api/v1.0"
)

var (
	prefixion = CONTEXT_PATH + VERSION

	controllerRouters = []ControllerRouter{
		userRouter,
		cvRouter,
		cvtRouter,
		// add new router here...
	}
)

func NewRouter() *mux.Router {

	muxRouter := mux.NewRouter().StrictSlash(true)
	for _, controllerRouter := range controllerRouters {
		for _, route := range controllerRouter {
			// var handler HandlerFunc
			handler := common.ValidateTokenHandlerFunc(HttpHandlerWrap(route), route.Name)
			handler = common.AppErrorHandlerFunc(handler)
			handler = CORSAllowMiddleware(route.Methods, handler)
			// handler = RouteLogMiddleware(handler, route.Name)
			muxRouter.Path(route.Pattern).Name(route.Name).Handler(handler).Methods(route.Methods...)
		}
	}
	// muxRouter.Handle("/",http.FileServer(http.Dir("dist")))
	// muxRouter.Use(mux.CORSMethodMiddleware(muxRouter))
	return muxRouter
}
