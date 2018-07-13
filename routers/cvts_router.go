package routers

import (
	"dpm/api"
)

var (
	cvtRouter = ControllerRouter{
		Route{
			Name:        "cvts_create",
			Methods:     []string{"POST", "OPTIONS"},
			Pattern:     prefixion + "/cvts",
			HandlerFunc: api.CreateCVT,
		},
		Route{
			Name:        "cvts_update",
			Methods:     []string{"PUT", "OPTIONS"},
			Pattern:     prefixion + "/cvts/{id}",
			HandlerFunc: api.UpdateCVT,
		},
		Route{
			Name:        "cvts_all_page",
			Methods:     []string{"GET", "OPTIONS"},
			Pattern:     prefixion + "/cvts",
			HandlerFunc: api.NewGetAllCVTS,
		},
	}
)
