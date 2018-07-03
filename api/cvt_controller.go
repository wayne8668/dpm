package api

import (
	"strconv"
	"dpm/common"
	"github.com/gorilla/mux"
	// "fmt"
	"dpm/models"
	"dpm/repositories"
	"net/http"
)

var (
	cvtsRepository = repositories.NewCVTRepository()
)

//新增模板
func CreateCVT(w http.ResponseWriter, r *http.Request) {

	// swagger:operation POST /cvts cvts CreateCVT
	//
	//新增模板
	//
	// Create a new cv template
	//
	// ---
	// Consumes:
	// - application/json
	// produces:
	// - application/json
	// parameters:
	// - name: body
	//   in: body
	//   description: request by
	//   required: true
	//   schema:
	//    "$ref": "#/definitions/CreateCVTParam"
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

	Logger.Info("CreateCVT method is invorked....")
	var cvt models.CVTemplate
	unmarshal2Object(w, r, &cvt)
	err := cvtsRepository.CreateNewCVTemplate(cvt)
	if err != nil {
		panic(common.ErrTrace(err))
	}
	//response
	rm := map[string]interface{}{
		"rsp_msg": "ok",
	}

	jsonResponseOK(w, rm)
}

//修改模板
func UpdateCVT(w http.ResponseWriter, r *http.Request) {

	// swagger:operation PUT /cvts/{id} cvts UpdateCVT
	//
	//修改模板
	//
	// Update the cv template
	//
	// ---
	// Consumes:
	// - application/json
	// produces:
	// - application/json
	// parameters:
	// - name: id
	//   in: path
	//   description: cv template id
	//   required: true
	// - name: body
	//   in: body
	//   description: request by
	//   required: true
	//   schema:
	//    "$ref": "#/definitions/CreateCVTParam"
	// responses:
	//   '200':
	//     description: "{\"rsp_msg\":\"ok\"}"
	//   '400':
	//     description: "{\"rsp_msg\":errro msg} - Bad Request Error"
	//   '401':
	//     description: "{\"rsp_msg\":errro msg} - Unauthorized Error"
	//   '403':
	//     description: "{\"rsp_msg\":errro msg} - Forbidden Error"
	//   '500':
	//     description: "{\"rsp_msg\":errro msg} - Internal Server Error"

	Logger.Info("UpdateCVT method is invorked....")
	vars := mux.Vars(r)
	cvtId := vars["id"]
	if cvtId == "" {
		panic(common.ErrBadRequest("the cvt id is required"))
	}
	var cvt models.CVTemplate
	cvt.CVTId = cvtId
	unmarshal2Object(w, r, &cvt)
	if err := cvtsRepository.UpdateCVTemplate(cvt); err != nil {
		panic(common.ErrTrace(err))
	}
	//response
	rm := map[string]interface{}{
		"rsp_msg": "ok",
	}
	jsonResponseOK(w, rm)
}


/*
* 返回所有模板-分页
 */
 func GetAllCVTS(w http.ResponseWriter, r *http.Request) {
	// swagger:operation GET /cvts cvts GetAllCVTS
	//
	//返回所有模板-分页
	//
	// Return All Users
	//
	// ---
	// Consumes:
	// - application/json
	// produces:
	// - application/json
	// parameters:
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
	//     description: "返回模板分页列表"
	//     schema:
	//       "$ref": "#/definitions/Pageable"
	//   '400':
	//     description: "{\"rsp_msg\":errro msg} - Bad Request Error"
	//   '401':
	//     description: "{\"rsp_msg\":errro msg} - Unauthorized Error"
	//   '500':
	//     description: "{\"rsp_msg\":errro msg} - Internal Server Error"

	// vars := mux.SetURLVars(r)
	// vars := r.URL.Query()
	ls := ParseQueryGet(r,"limit")
	pgs := ParseQueryGet(r,"page")
	Logger.Infof("ls:[%s];pgs:[%s]", ls,pgs)

	
	l, _ := strconv.ParseInt(ls, 10, 64)
	pg, _ := strconv.ParseInt(pgs, 10, 64)

	Logger.Info("==================", l, pg)

	p, err := common.NewPageable(l, pg)

	if err != nil {
		panic(common.ErrTrace(err))
	}

	pr, err := cvtsRepository.GetAllCVTS(p)

	if err != nil {
		panic(common.ErrTrace(err))
	}

	jsonResponseOK(w, &pr)
}