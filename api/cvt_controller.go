package api

import (
	"dpm/common"
	"dpm/models"
	"dpm/repositories"
)

var (
	cvtsRepository = repositories.NewCVTRepository()
)

//新增模板
func CreateCVT(req *common.ApiRequest) (rsp common.ApiRsponse) {
	Logger.Info("CreateCVT method is invorked....")
	var cvt models.CVTemplate
	if err := req.BindJSON(&cvt); err != nil {
		return rsp.Error(common.ErrTrace(err))
	}
	return rsp.Error(cvtsRepository.CreateNewCVTemplate(cvt))
}

type UpdateCVTRequest struct {
	Id  string            `path:"id" binding:"required"`
	CVT models.CVTemplate `struct:"json"`
}

//修改模板
func UpdateCVT(req *common.ApiRequest) (rsp common.ApiRsponse) {
	var param UpdateCVTRequest
	if err := req.BindStruct(&param); err != nil {
		return rsp.Error(common.ErrTrace(err))
	}
	param.CVT.CVTId = param.Id
	return rsp.Error(cvtsRepository.UpdateCVTemplate(param.CVT))
}

//返回所有模板-分页
func GetAllCVTS(req *common.ApiRequest) (rsp common.ApiRsponse) {
	var pr PageableRequest
	if err := req.BindStruct(&pr); err != nil {
		return rsp.Error(common.ErrTrace(err))
	}
	p, err := common.NewPageable(pr.Limit, pr.Page)
	if err != nil {
		return rsp.Error(common.ErrTrace(err))
	}
	p, err = cvtsRepository.GetAllCVTS(p)
	if err != nil {
		return rsp.Error(common.ErrTrace(err))
	}
	return rsp.AddObject(p)
}
