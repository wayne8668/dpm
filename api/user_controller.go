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
func Loggin(u models.User) (m RspModel, err error) {

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

	if u, err := usersRepository.GetUserForAuth(u); err != nil {
		return m, common.ErrTrace(err)
	} else {
		if u.UId == "" {
			return m, common.ErrForbidden("user name or pwd err.")
		}

		ut := &common.UserToken{
			Name: u.Name,
			Pwd:  u.Pwd,
			Id:   u.UId,
		}
		if token, err := common.CreateToken(ut); err != nil {
			return m, common.ErrTrace(err)
		} else {
			return m.AddAttribute(common.TOKEN_KEY, token).
				AddAttribute("auth_user", u), nil
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
func RegisterUser(u models.User) error {
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

	return CreateUser(u)
}

/*
* 返回用户列表
 */
func GetAllUsers(req PageableRequest) (p common.Pageable, err error) {
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

	if p, err = common.NewPageable(req.Limit, req.Page); err != nil {
		return p, common.ErrTrace(err)
	}

	return usersRepository.GetAllUsers(p)
}
