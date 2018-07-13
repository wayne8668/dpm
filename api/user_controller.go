/*
* User 控制器
 */
package api

import (
	"dpm/common"
	"dpm/models"
	"dpm/repositories"
)

var (
	usersRepository = repositories.NewUsersRepository()
)

//用户登录
func Loggin(u models.User) map[string]interface{} {

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

	if udb, err := usersRepository.GetUserForAuth(u); err != nil {
		panic(common.ErrTrace(err))
	} else {
		if udb == nil {
			panic(common.ErrForbidden("user name or pwd err."))
		}

		ut := &common.UserToken{
			Name: u.Name,
			Pwd:  u.Pwd,
			Id:   udb.UId,
		}
		if token, err := common.CreateToken(ut); err != nil {
			panic(common.ErrTrace(err))
		} else {
			m := make(map[string]interface{})
			m[common.TOKEN_KEY] = token
			m["auth_user"] = udb
			return m
		}
	}
}

//新增用户
func CreateUser(u models.User) error {

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

	return usersRepository.CreateUser(u)
}

//用户注册
func RegisterUser(u models.User) {
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
	//     description: "{\"numResult\":numResult}"
	//   '403':
	//     description: "{\"rsp_msg\":errro msg} - Forbidden Error"
	//   '500':
	//     description: "{\"rsp_msg\":errro msg} - Internal Server Error"

	CreateUser(u)
}

/*
* 返回用户列表
 */
func GetAllUsers(req PageableRequest) (common.Pageable,error) {
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


	p, err := common.NewPageable(req.Limit, req.Page)

	if err != nil {
		panic(common.ErrTrace(err))
	}

	return usersRepository.GetAllUsers(p)
}
