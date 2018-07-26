package routers

import (
	"dpm/api"
	"dpm/common"
)

var (
	userRouter = ControllerRouter{
		Route{
			Name:        "users_all_page",
			Methods:     []string{"GET"},
			Pattern:     "/users",
			HandlerFunc: common.HttpFuncWrap(api.GetAllUsers),
		},
		Route{
			Name:        "users_create",
			Methods:     []string{"POST"},
			Pattern:     "/users",
			HandlerFunc: common.HttpFuncWrap(api.CreateUser),
		},
		Route{
			Name:        "users_login",
			Methods:     []string{"POST"},
			Pattern:     "/users/login",
			HandlerFunc: common.HttpFuncWrap(api.Loggin),
		},
		Route{
			Name:        "users_register",
			Methods:     []string{"POST"},
			Pattern:     "/users/register",
			HandlerFunc: common.HttpFuncWrap(api.RegisterUser),
		},
	}
)
