package api

import (
	"dpm/common"
	"strings"
	"github.com/gorilla/mux"
	"dpm/models"
	"dpm/repositories"
	"fmt"
	"net/http"
	// "github.com/goinggo/mapstructure"
)

var (
	cvsRepositories = repositories.NewCVSRepository()
)

//我的简历
 func GetUsersCVS(w http.ResponseWriter, r *http.Request) {
	// swagger:operation GET /cvs/{uid} cvs GetUsersCVS
	//
	//返回我的简历
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
	uidpath := vars["uid"]
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

/*
* 创建简历
 */
func CreateCV(w http.ResponseWriter, r *http.Request) {
	var cv models.CurriculumVitae
	unmarshal2Object(w, r, &cv)
	fmt.Println(cv)
	fmt.Println(cv.Height)
	rm := make(map[string]interface{})
	cvsRepositories.CreateCV(cv)
	jsonResponse(w, http.StatusOK, rm)
}

/*
* 根据模板创建简历
 */
func CreateCVForTempId(w http.ResponseWriter, r *http.Request) {

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
