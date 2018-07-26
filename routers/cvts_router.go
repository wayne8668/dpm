package routers

import (
	"dpm/api"
	"dpm/common"
)

var (
	cvtRouter = ControllerRouter{
		Route{
			Name:        "cvts_create",
			Methods:     []string{"POST"},
			Pattern:     "/cvts",
			HandlerFunc: common.HttpFuncWrap(api.CreateCVT),
		},
		Route{
			Name:        "cvts_update",
			Methods:     []string{"PUT"},
			Pattern:     "/cvts/:id",
			HandlerFunc: common.HttpFuncWrap(api.UpdateCVT),
		},
		Route{
			Name:        "cvts_all_page",
			Methods:     []string{"GET"},
			Pattern:     "/cvts",
			HandlerFunc: common.HttpFuncWrap(api.GetAllCVTS),
		},
	}
)
