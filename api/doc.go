package api

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

///////////////////////////////////////////////////////////////////////////////

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

///////////////////////////////////////////////////////////////////////////////

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

///////////////////////////////////////////////////////////////////////////////

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

///////////////////////////////////////////////////////////////////////////////

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

///////////////////////////////////////////////////////////////////////////////

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

///////////////////////////////////////////////////////////////////////////////

// swagger:operation POST /users/login users Loggin
//
//用户登录
//
// User Login
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
//    "$ref": "#/definitions/RegisterUserParam"
// responses:
//   '200':
//     description: "{\"Authorization\":token string}"
//   '403':
//     description: "{\"rsp_msg\":errro msg} - Forbidden Error"
//   '500':
//     description: "{\"rsp_msg\":errro msg} - Internal Server Error"

///////////////////////////////////////////////////////////////////////////////

// swagger:operation POST /users users CreateUser
//
//新增用户
//
// Create a new user
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
//    "$ref": "#/definitions/CreateUserParam"
// responses:
//   '200':
//     description: "{\"rsp_msg\":ok}"
//   '401':
//     description: "{\"rsp_msg\":errro msg} - Unauthorized Error"
//   '403':
//     description: "{\"rsp_msg\":errro msg} - Forbidden Error"
//   '500':
//     description: "{\"rsp_msg\":errro msg} - Internal Server Error"

///////////////////////////////////////////////////////////////////////////////

// swagger:operation POST /users/register users RegisterUser
//
////用户注册
//
// Create a new user
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
//    "$ref": "#/definitions/RegisterUserParam"
// responses:
//   '200':
//     description: "{\"rsp_msg\":\"ok\"}"
//   '403':
//     description: "{\"rsp_msg\":errro msg} - Forbidden Error"
//   '500':
//     description: "{\"rsp_msg\":errro msg} - Internal Server Error"

///////////////////////////////////////////////////////////////////////////////

// swagger:operation GET /users users GetAllUsers
//
//返回用户列表
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
//     description: "返回用户分页列表"
//     schema:
//       "$ref": "#/definitions/Pageable"
//   '400':
//     description: "{\"rsp_msg\":errro msg} - Bad Request Error"
//   '401':
//     description: "{\"rsp_msg\":errro msg} - Unauthorized Error"
//   '500':
//     description: "{\"rsp_msg\":errro msg} - Internal Server Error"
