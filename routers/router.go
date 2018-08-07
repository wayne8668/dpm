package routers

import (
	"dpm/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	// "net/http"
)

type Route struct {
	Name        string
	Methods     []string
	Pattern     string
	HandlerFunc gin.HandlerFunc
}

type (
	ControllerRouter []Route
)

var (
	controllerRouters = []ControllerRouter{
		optionsRouter,
		userRouter,
		cvRouter,
		cvtRouter,
		// add new router here...
	}

	middlewares = []gin.HandlerFunc{
		gin.Logger(),
		middleware.Recovery(),
		middleware.ValidateTokenHandlerFunc(),
		CorsMW(),
	}
)

func CorsMW() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowAllOrigins:  false,
		AllowOrigins:     []string{"http://localhost:8082"},
		AllowMethods:     []string{"GET", "POST", "PUT", "OPTIONS", "DELETE", "HEAD"},
		AllowHeaders:     []string{"authorization", "content-type"},
		AllowCredentials: true,
		// MaxAge: 12 * time.Hour,
	})
}

// func NewRouter() *mux.Router {

// 	muxRouter := mux.NewRouter().StrictSlash(true)
// 	for _, controllerRouter := range controllerRouters {
// 		for _, route := range controllerRouter {
// 			// var handler HandlerFunc
// 			handler := common.ValidateTokenHandlerFunc(HttpHandlerWrap(route), route.Name)
// 			handler = common.AppErrorHandlerFunc(handler)
// 			handler = CORSAllowMiddleware(route.Methods, handler)
// 			muxRouter.Path(route.Pattern).Name(route.Name).Handler(handler).Methods(route.Methods...)
// 		}
// 	}
// 	// muxRouter.Handle("/",http.FileServer(http.Dir("dist")))
// 	// muxRouter.Use(mux.CORSMethodMiddleware(muxRouter))
// 	return muxRouter
// }

func RegisterMiddleWare(r *gin.RouterGroup) {
	for _, middleware := range middlewares {
		r.Use(middleware)
	}
}

func RegisterRouter(r *gin.RouterGroup) {
	for _, controllerRouter := range controllerRouters {
		for _, route := range controllerRouter {
			for _, v := range route.Methods {
				r.Handle(v, route.Pattern, route.HandlerFunc)
			}
		}
	}
}
