package api

import (
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
	// swagger:operation GET /cvs/users/{uid} cvs GetUsersCVS
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
	//   in: path
	//   description: cvs of user's id
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
	//   '500':
	//     description: "{\"rsp_msg\":errro msg} - Internal Server Error"

	vars := mux.Vars(r)
	uidpath := vars["uid"] //
	uid := strings.TrimSpace(uidpath) 

	Logger.Infof("GetUsersCVS param uid is:[%s]", uid)

	if uid =="" {
		panic(common.ErrBadRequest("Bad Request param in path:{uid} is null"))
	}

	cvs, err := cvsRepositories.GetUsersCVS(uid)

	if err != nil {
		panic(common.ErrTrace(err))
	}

	jsonResponseOK(w, &cvs)
}

//新增简历
func CreateUsersCVS(w http.ResponseWriter, r *http.Request) {
	// swagger:operation POST /cvs/{uid} cvs CreateUsersCVS
	//
	//新增简历
	//
	// create User's CVS
	//
	// ---
	// Consumes:
	// - application/json
	// produces:
	// - application/json
	// parameters:
	// - name: uid
	//   in: path
	//   description: cvs of user's id
	//   required: true
	// responses:
	//   '200':
	//     description: "创建用户简历"
	//     schema:
	//       "$ref": "#/definitions/CurriculumVitae"
	//   '400':
	//     description: "{\"rsp_msg\":errro msg} - Bad Request Error"
	//   '401':
	//     description: "{\"rsp_msg\":errro msg} - Unauthorized Error"
	//   '500':
	//     description: "{\"rsp_msg\":errro msg} - Internal Server Error"


	vars := mux.Vars(r)
	uidpath := vars["uid"]
	uid := strings.TrimSpace(uidpath)

	Logger.Infof("GetUsersCVS param uid is:[%s]", uid)

	if uid =="" {
		panic(common.ErrBadRequest("Bad Request param in path:{uid} is null"))
	}

	var cv models.CurriculumVitae
	unmarshal2Object(w, r, &cv)

	nr,err := cvsRepositories.CreateUsersCVS(cv,uid)

	if err != nil {
		panic(common.ErrTrace(err))
	}
	jsonResponseOK(w, &nr)
}

//修改简历模板
func UpdateCVTemp(w http.ResponseWriter, r *http.Request) {

}

//新增简历模板
func CreateCVTemp(w http.ResponseWriter, r *http.Request) {

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
