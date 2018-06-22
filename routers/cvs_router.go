package routers

import (
	"dpm/api"
)

var (
	cvRouter = ControllerRouter{
		Route{
			Name:        "cvs_create",
			Methods:     []string{"POST"},
			Pattern:     prefixion + "/cvs",
			HandlerFunc: api.CreateCV,
		},
		Route{
			Name:        "cvs_create_index",
			Methods:     []string{"GET"},
			Pattern:     prefixion + "/cvs",
			HandlerFunc: api.CreateCVIndex,
		},
		Route{
			Name:        "cvs_create_temp_id",
			Methods:     []string{"GET"},
			Pattern:     prefixion + "/cvs/template/{id}",
			HandlerFunc: api.CreateCVForTempId,
		},
		Route{
			Name:        "cvs_users",
			Methods:     []string{"GET", "OPTIONS"},
			Pattern:     prefixion + "/cvs/{uid}",
			HandlerFunc: api.GetUsersCVS,
		},
	}
)
