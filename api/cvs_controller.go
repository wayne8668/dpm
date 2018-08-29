package api

import (
	"dpm/common"
	"dpm/middleware"
	"dpm/repositories"
	"fmt"
	"net/http"

	"dpm/models"
)

var (
	cvsRepositories = repositories.NewCVSRepository()
)

type GetUsersCVSRequest struct {
	PageableRequest `struct:"+"`
	Uid             string `query:"uid" binding:"required"`
}

//返回用户所有简历
func GetUsersCVS(req *common.ApiRequest) (rsp common.ApiRsponse) {

	var param GetUsersCVSRequest
	if err := req.BindStruct(&param); err != nil {
		return rsp.Error(common.ErrTrace(err))
	}

	//判断是否为当前用户操作
	uti, _ := req.Get(CURRENT_USER)
	ut := uti.(*middleware.UserToken)
	if ut.Id != param.Uid {
		return rsp.Error(common.ErrForbidden("You can create cv only for youself..."))
	}

	p, err := common.NewPageable(param.Limit, param.Page)
	if err != nil {
		return rsp.Error(common.ErrTrace(err))
	}

	p, err = cvsRepositories.GetUsersCVS(p, param.Uid)
	if err != nil {
		return rsp.Error(common.ErrTrace(err))
	}
	return rsp.AddObject(p)
}

type ReSetCVTempRequest struct {
	Cvtid string `path:"cvtid" binding:"required"`
	Cvid  string `path:"cvid" binding:"required"`
	Uid   string `query:"uid" binding:"required"`
}

//修改简历模板
func ReSetCVTemp(req *common.ApiRequest) (rsp common.ApiRsponse) {
	var param ReSetCVTempRequest
	if err := req.BindStruct(&param); err != nil {
		return rsp.Error(common.ErrTrace(err))
	}

	//判断是否为当前用户操作
	uti, _ := req.Get(CURRENT_USER)
	ut := uti.(*middleware.UserToken)
	if ut.Id != param.Uid {
		return rsp.Error(common.ErrForbidden("You can create cv only for youself..."))
	}

	//判断模板是否存在
	if ok, err := cvtsRepository.IsExist(param.Cvtid); err != nil {
		return rsp.Error(common.ErrTrace(err))
	} else if !ok {
		return rsp.Error(common.ErrBadRequest("the path param cvtid:[%s] is not exist..."))
	}
	err := cvsRepositories.ReSetCVTemp(param.Uid, param.Cvid, param.Cvtid)
	return rsp.Error(err)

}

type CreateCVWithTempRequest struct {
	Cvtid string `path:"cvtid" binding:"required"`
	Uid   string `query:"uid" binding:"required"`
}

//新增简历
func CreateCVWithTemp(req *common.ApiRequest) (rsp common.ApiRsponse) {
	var param CreateCVWithTempRequest
	if err := req.BindStruct(&param); err != nil {
		return rsp.Error(common.ErrTrace(err))
	}
	//判断是否为当前用户操作
	uti, _ := req.Get(CURRENT_USER)
	ut := uti.(*middleware.UserToken)
	fmt.Println(ut.Id, param.Uid)
	if ut.Id != param.Uid {
		return rsp.Error(common.ErrForbidden("You can create cv only for youself..."))
	}

	//判断模板是否存在
	if ok, err := cvtsRepository.IsExist(param.Cvtid); err != nil {
		return rsp.Error(common.ErrTrace(err))
	} else if !ok {
		return rsp.Error(common.ErrBadRequest("the path param cvtid:[%s] is not exist..."))
	}

	cvid, err := cvsRepositories.CreateCVWithTemp(param.Uid, param.Cvtid)
	return rsp.Error(err).AddAttribute("cvid", cvid)
}

type CreateBasicInfoRequest struct {
	Cvid      string              `path:"cvid" binding:"required"`
	Uid       string              `query:"uid" binding:"required"`
	BasicInfo models.BasicInfoCVM `struct:"+"`
}

//新增基本信息
func CreateBasicInfoCVM(req *common.ApiRequest) (rsp common.ApiRsponse) {
	var param CreateBasicInfoRequest
	if err := req.BindStruct(&param); err != nil {
		return rsp.Error(common.ErrTrace(err))
	}
	//判断是否为当前用户操作
	uti, _ := req.Get(CURRENT_USER)
	ut := uti.(*middleware.UserToken)
	fmt.Println(ut.Id, param.Uid)
	if ut.Id != param.Uid {
		return rsp.Error(common.ErrForbidden("You can create cv only for youself..."))
	}

	//cvid, err := cvsRepositories.CreateCVWithTemp(param.Uid, param.Cvtid)
	//todo 服务代码没写
	return rsp.AddAttribute("cvid", "")

}

/*
* 创建简历首页
 */
func CreateCVIndex(w http.ResponseWriter, r *http.Request) {

}
