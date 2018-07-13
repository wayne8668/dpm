package api

import (
	"dpm/common"
	"dpm/models"
	"dpm/repositories"
)

var (
	cvtsRepository = repositories.NewCVTRepository()
)

type CreateCVTRequest struct{
	CVT models.CVTemplate	`qval:"inbody"`
}

//新增模板
func CreateCVT(req CreateCVTRequest) error {

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
	return cvtsRepository.CreateNewCVTemplate(req.CVT)
}

type UpdateCVTRequest struct{
	Id string	`qval:"id,inpath"`
	CVT models.CVTemplate	`qval:"inbody"`
}

//修改模板
func UpdateCVT(req UpdateCVTRequest) error {

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


	req.CVT.CVTId = req.Id
	return cvtsRepository.UpdateCVTemplate(req.CVT)
}

/*
* 返回所有模板-分页
 */
 func GetAllCVTS(req PageableRequest) common.Pageable {
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


	p, err := common.NewPageable(10, req.Page)

	if err != nil {
		panic(common.ErrTrace(err))
	}

	pr, err := cvtsRepository.GetAllCVTS(p)

	if err != nil {
		panic(common.ErrTrace(err))
	}

	return pr
}