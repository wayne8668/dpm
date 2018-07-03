package api

import (
	"strconv"
	"github.com/gorilla/context"
	"dpm/common"
	"strings"
	"github.com/gorilla/mux"
	"dpm/models"
	"dpm/repositories"
	"net/http"
	// "github.com/goinggo/mapstructure"
)

var (
	cvsRepositories = repositories.NewCVSRepository()
)

//返回指定用户的所有简历
 func GetUsersCVS(w http.ResponseWriter, r *http.Request) {
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


	ls := ParseQueryGet(r,"limit")
	pgs := ParseQueryGet(r,"page")
	Logger.Infof("ls:[%s];pgs:[%s]", ls,pgs)

	uid := ParseQueryGet(r,"uid")
	uid = strings.TrimSpace(uid)

	Logger.Infof("GetUsersCVS param uid is:[%s]", uid)

	if uid =="" {
		panic(common.ErrBadRequest("Bad Request param in query:{uid} is null"))
	}

	l, _ := strconv.ParseInt(ls, 10, 64)
	pg, _ := strconv.ParseInt(pgs, 10, 64)

	p, err := common.NewPageable(l, pg)

	if err != nil {
		panic(common.ErrTrace(err))
	}

	cvs, err := cvsRepositories.GetUsersCVS(p,uid)

	if err != nil {
		panic(common.ErrTrace(err))
	}

	jsonResponseOK(w, &cvs)
}

//修改简历模板
func ReSetCVTemp(w http.ResponseWriter, r *http.Request) {
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
//     schema:
//       "$ref": "#/definitions/CurriculumVitae"
//   '400':
//     description: "{\"rsp_msg\":errro msg} - Bad Request Error"
//   '401':
//     description: "{\"rsp_msg\":errro msg} - Unauthorized Error"
//   '403':
//     description: "{\"rsp_msg\":errro msg} - Forbidden Error"
//   '500':
//     description: "{\"rsp_msg\":errro msg} - Internal Server Error"

	vars := mux.Vars(r)
	cvtid := vars["cvtid"]
	cvid := vars["cvid"]
	uid := ParseQueryGet(r,"uid")
	uid = strings.TrimSpace(uid)

	//判断是否为当前用户操作
	ut := context.Get(r,CURRENT_USER).(*common.UserToken)
	if ut.Id != uid {
		panic(common.ErrForbidden("You can create cv only for youself..."))
	}

	//判断模板是否存在
	if ok,err := cvtsRepository.IsExist(cvtid);err != nil {
		panic(common.ErrTrace(err))
	}else if !ok {
		panic(common.ErrBadRequest("the path param cvtid:[%s] is not exist..."))
	}

	err := cvsRepositories.ReSetCVTemp(uid,cvid,cvtid)

	if err!=nil {
		panic(common.ErrTrace(err))
	}

	rm := make(map[string]interface{})
	rm["rsp_msg"] = "ok"
	jsonResponseOK(w, &rm)

}

//新增简历
func CreateCVWithTemp(w http.ResponseWriter, r *http.Request) {
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
//     schema:
//       "$ref": "#/definitions/CurriculumVitae"
//   '400':
//     description: "{\"rsp_msg\":errro msg} - Bad Request Error"
//   '401':
//     description: "{\"rsp_msg\":errro msg} - Unauthorized Error"
//   '403':
//     description: "{\"rsp_msg\":errro msg} - Forbidden Error"
//   '500':
//     description: "{\"rsp_msg\":errro msg} - Internal Server Error"

	vars := mux.Vars(r)
	cvtid := vars["cvtid"]
	uid := ParseQueryGet(r,"uid")
	uid = strings.TrimSpace(uid)

	//判断是否为当前用户操作
	ut := context.Get(r,CURRENT_USER).(*common.UserToken)
	if ut.Id != uid {
		panic(common.ErrForbidden("You can create cv only for youself..."))
	}

	//判断模板是否存在
	if ok,err := cvtsRepository.IsExist(cvtid);err != nil {
		panic(common.ErrTrace(err))
	}else if !ok {
		panic(common.ErrBadRequest("the path param cvtid:[%s] is not exist..."))
	}

	cvid ,err := cvsRepositories.CreateCVWithTemp(uid,cvtid)
	
	if err!=nil {
		panic(common.ErrTrace(err))
	}

	rm := make(map[string]interface{})
	rm["cvid"] = cvid
	jsonResponseOK(w, &rm)
}

/*
* 创建简历首页
 */
func CreateCVIndex(w http.ResponseWriter, r *http.Request) {
	// cv := models.NewCurriculumVitae()
	rm := make(map[string]interface{})
	rm["cv"] = models.NewCurriculumVitae()
	jsonResponse(w, http.StatusOK, rm)
}
