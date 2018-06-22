package routers

import (
	"dpm/api"
)

var (
	userRouter = ControllerRouter{
		Route{
			Name:        "users_all_page",
			Methods:     []string{"GET"},
			Pattern:     prefixion + "/users",
			HandlerFunc: api.GetAllUsers,
		},
		Route{
			Name:        "users_create",
			Methods:     []string{"POST", "OPTIONS"},
			Pattern:     prefixion + "/users",
			HandlerFunc: api.CreateUser,
		},
		Route{
			Name:        "users_login",
			Methods:     []string{"POST", "OPTIONS"},
			Pattern:     prefixion + "/users/login",
			HandlerFunc: api.Loggin,
		},
		Route{
			Name:        "users_register",
			Methods:     []string{"POST", "OPTIONS"},
			Pattern:     prefixion + "/users/register",
			HandlerFunc: api.RegisterUser,
		},
	}
)
