package routers

import (
	"dpm/api"
)

var (
	cvRouter = ControllerRouter{
		Route{
			Name:        "cvs_update_cvm_cvt", //修改简历模板
			Methods:     []string{"PUT", "OPTIONS"},
			Pattern:     prefixion + "/cvs/{cvid}/cvms/cvt/{cvtid}",
			HandlerFunc: api.ReSetCVTemp,
		},
		Route{
			Name:        "cvs_create_cvm_cvt", //新增简历
			Methods:     []string{"POST", "OPTIONS"},
			Pattern:     prefixion + "/cvs/cvms/cvt/{cvtid}",
			HandlerFunc: api.CreateCVWithTemp,
		},
		Route{
			Name:        "cvs_users", //返回指定用户的所有简历
			Methods:     []string{"GET", "OPTIONS"},
			Pattern:     prefixion + "/cvs",
			HandlerFunc: api.GetUsersCVS,
		},
		Route{
			Name:        "cvs_users_index", //返回指定用户的指定简历
			Methods:     []string{"GET", "OPTIONS"},
			Pattern:     prefixion + "/cvs/{cvid}/users/{uid}",
			HandlerFunc: api.GetUsersCVS,
		},
	}
)
