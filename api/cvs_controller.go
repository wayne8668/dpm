package api

import (
	"dpm/common"
	"dpm/repositories"
	"github.com/gorilla/context"
	"net/http"
	// "github.com/goinggo/mapstructure"
)

var (
	cvsRepositories = repositories.NewCVSRepository()
)

type GetUsersCVSRequest struct {
	PageableRequest `qval:"+"`
	Uid             string `qval:"uid,inquery"`
}

//返回指定用户的所有简历
func GetUsersCVS(req GetUsersCVSRequest) (p common.Pageable, err error) {
	// swagger:operation GET /cvs cvs GetUsersCVS
	//
	//返回指定用户的所有简历
	//
	// Return User's CVS
	//
	// ---
	// Consumes:
	// - application/json
	// produces:
	// - application/json
	// parameters:
	// - name: uid
	//   in: query
	//   description: cvs of user's id
	//   required: true
	// - name: limit
	//   in: query
	//   description: per page limit
	//   required: true
	// - name: page
	//   in: query
	//   description: page number
	//   required: true
	// responses:
	//   '200':
	//     description: "返回用户简历"
	//     schema:
	//       "$ref": "#/definitions/Pageable"
	//   '400':
	//     description: "{\"rsp_msg\":errro msg} - Bad Request Error"
	//   '401':
	//     description: "{\"rsp_msg\":errro msg} - Unauthorized Error"
	//   '403':
	//     description: "{\"rsp_msg\":errro msg} - Forbidden Error"
	//   '500':
	//     description: "{\"rsp_msg\":errro msg} - Internal Server Error"

	if p, err = common.NewPageable(req.Limit, req.Page); err != nil {
		return p, common.ErrTrace(err)
	}

	return cvsRepositories.GetUsersCVS(p, req.Uid)
}

type ReSetCVTempRequest struct {
	Cvtid string `qval:"cvtid,inpath"`
	Cvid  string `qval:"cvid,inpath"`
	Uid   string `qval:"uid,inquery"`
}

//修改简历模板
func ReSetCVTemp(req ReSetCVTempRequest, r *http.Request) (err error) {
	// swagger:operation PUT /cvs/{cvid}/cvms/cvt/{cvtid} cvs ReSetCVTemp
	//
	//修改简历模板
	//
	// Reset User's CVS with template
	//
	// ---
	// Consumes:
	// - application/json
	// produces:
	// - application/json
	// parameters:
	// - name: cvtid
	//   in: path
	//   description: cvs of template id
	//   required: true
	// - name: cvid
	//   in: path
	//   description: cv's id
	//   required: true
	// - name: uid
	//   in: query
	//   description: cvs of user id
	//   required: true
	// responses:
	//   '200':
	//     description: "{\"rsp_msg\":ok}"
	//   '400':
	//     description: "{\"rsp_msg\":errro msg} - Bad Request Error"
	//   '401':
	//     description: "{\"rsp_msg\":errro msg} - Unauthorized Error"
	//   '403':
	//     description: "{\"rsp_msg\":errro msg} - Forbidden Error"
	//   '500':
	//     description: "{\"rsp_msg\":errro msg} - Internal Server Error"

	//判断是否为当前用户操作
	ut := context.Get(r, CURRENT_USER).(*common.UserToken)
	if ut.Id != req.Uid {
		return common.ErrForbidden("You can create cv only for youself...")
	}

	//判断模板是否存在
	if ok, err := cvtsRepository.IsExist(req.Cvtid); err != nil {
		return common.ErrTrace(err)
	} else if !ok {
		return common.ErrBadRequest("the path param cvtid:[%s] is not exist...")
	}

	return cvsRepositories.ReSetCVTemp(req.Uid, req.Cvid, req.Cvtid)

}

type CreateCVWithTempRequest struct {
	Cvtid string `qval:"cvtid,inpath"`
	Uid   string `qval:"uid,inquery"`
}

//新增简历
func CreateCVWithTemp(req CreateCVWithTempRequest, r *http.Request) (m RspModel, err error) {
	// swagger:operation POST /cvs/cvms/cvt/{cvtid} cvs CreateCVWithTemp
	//
	//新增简历
	//
	// create User's CVS with template
	//
	// ---
	// Consumes:
	// - application/json
	// produces:
	// - application/json
	// parameters:
	// - name: cvtid
	//   in: path
	//   description: cvs of template id
	//   required: true
	// - name: uid
	//   in: query
	//   description: cvs of user id
	//   required: true
	// responses:
	//   '200':
	//     description: "{\"cvid\":cvid}"
	//   '400':
	//     description: "{\"rsp_msg\":errro msg} - Bad Request Error"
	//   '401':
	//     description: "{\"rsp_msg\":errro msg} - Unauthorized Error"
	//   '403':
	//     description: "{\"rsp_msg\":errro msg} - Forbidden Error"
	//   '500':
	//     description: "{\"rsp_msg\":errro msg} - Internal Server Error"

	//判断是否为当前用户操作
	ut := context.Get(r, CURRENT_USER).(*common.UserToken)
	if ut.Id != req.Uid {
		return m, common.ErrForbidden("You can create cv only for youself...")
	}

	//判断模板是否存在
	if ok, err := cvtsRepository.IsExist(req.Cvtid); err != nil {
		return m, common.ErrTrace(err)
	} else if !ok {
		return m, common.ErrBadRequest("the path param cvtid:[%s] is not exist...")
	}

	cvid, err := cvsRepositories.CreateCVWithTemp(req.Uid, req.Cvtid)
	return m.AddAttribute("cvid", cvid).AddAttribute("haha", "aiyo"), err
}

/*
* 创建简历首页
 */
func CreateCVIndex(w http.ResponseWriter, r *http.Request) {

}
