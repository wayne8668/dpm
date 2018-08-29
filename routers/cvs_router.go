package routers

import (
	"dpm/api"
	"dpm/common"
)

var (
	cvRouter = ControllerRouter{
		Route{
			Name:        "cvs_create_cvm_bsinfo", //新增基本信息
			Methods:     []string{"POST"},
			Pattern:     "/cvs/:cvid/cvms/bsinfo",
			HandlerFunc: common.HttpFuncWrap(api.CreateBasicInfoCVM),
		},
		Route{
			Name:        "cvs_update_cvm_cvt", //修改简历模板
			Methods:     []string{"PUT"},
			Pattern:     "/cvs/:cvid/cvms/cvt/:cvtid",
			HandlerFunc: common.HttpFuncWrap(api.ReSetCVTemp),
		},
		Route{
			Name:        "cvs_create_cvm_cvt", //新增简历
			Methods:     []string{"POST"},
			Pattern:     "/cvs/cvms/cvt/:cvtid",
			HandlerFunc: common.HttpFuncWrap(api.CreateCVWithTemp),
		},
		Route{
			Name:        "cvs_users", //返回指定用户的所有简历
			Methods:     []string{"GET"},
			Pattern:     "/cvs",
			HandlerFunc: common.HttpFuncWrap(api.GetUsersCVS),
		},
		// Route{
		// 	Name:        "cvs_users_index", //返回指定用户的指定简历
		// 	Methods:     []string{"GET"},
		// 	Pattern:     "/cvs/:cvid/users/:uid",
		// 	HandlerFunc: common.HttpFuncWrap(api.GetUsersCVS),
		// },
	}
)
