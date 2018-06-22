package api

import (
	"github.com/gorilla/mux"
	"github.com/juju/errors"
	// "fmt"
	"dpm/models"
	"dpm/repositories"
	"net/http"
)

var (
	cvtRepositories = repositories.NewCVTRepository()
)

//新增简历模板
func CreateCVT(w http.ResponseWriter, r *http.Request) {

	// swagger:operation POST /cvts cvts CreateCVT
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
	//    "$ref": "#/definitions/CVTemplate"
	//
	//
	// responses:
	//   '200':
	//     description: "{\"numResult\":numResult}"
	//   default:
	//     description: unexpected error
	//     schema:
	//       "$ref": "#/responses/ResponseMsg"

	Logger.Info("CreateCVT method is invorked....")
	var cvt models.CVTemplate
	unmarshal2Object(w, r, &cvt)
	numResult, err := cvtRepositories.CreateNewCVTemplate(cvt)
	if err != nil {
		panic(errors.Trace(err))
	}
	//response
	rm := map[string]interface{}{
		"numResult": numResult,
	}

	jsonResponse(w, http.StatusOK, rm)
}

//修改简历模板
func UpdateCVT(w http.ResponseWriter, r *http.Request) {

	// swagger:operation PUT /cvts/{id} cvts UpdateCVT
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
	//    "$ref": "#/definitions/CVTemplate"
	// responses:
	//   '200':
	//     description: "ok"
	//   default:
	//     description: unexpected error

	Logger.Info("UpdateCVT method is invorked....")
	vars := mux.Vars(r)
	cvtId := vars["id"]
	if cvtId == "" {
		panic(errors.NewErr("the cvt id is required"))
	}
	var cvt models.CVTemplate
	cvt.CVTId = cvtId
	unmarshal2Object(w, r, &cvt)
	if err := cvtRepositories.UpdateCVTemplate(cvt); err != nil {
		panic(errors.Trace(err))
	}
	jsonResponse(w, http.StatusOK, "ok")
}
