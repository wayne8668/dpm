package routers

import (
	"dpm/api"
)

var (
	cvRouter = ControllerRouter{
		Route{																
			Name:        "cvs_update_cvt",								//修改简历模板
			Methods:     []string{"PUT", "OPTIONS"},
			Pattern:     prefixion + "/cvs/{cvid}/users/{uid}/cvt",
			HandlerFunc: api.UpdateCVTemp,
		},
		Route{																
			Name:        "cvs_create_cvt",								//新增简历模板
			Methods:     []string{"POST", "OPTIONS"},
			Pattern:     prefixion + "/cvs/users/{uid}/cvt",
			HandlerFunc: api.CreateCVTemp,
		},
		Route{
			Name:        "cvs_users",									//返回指定用户的所有简历
			Methods:     []string{"GET", "OPTIONS"},
			Pattern:     prefixion + "/cvs/users/{uid}",
			HandlerFunc: api.GetUsersCVS,
		},
		Route{
			Name:        "cvs_users_index",								//返回指定用户的指定简历
			Methods:     []string{"GET", "OPTIONS"},
			Pattern:     prefixion + "/cvs/{cvid}/users/{uid}",
			HandlerFunc: api.GetUsersCVS,
		},
	}
)
