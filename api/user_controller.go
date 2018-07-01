/*
* User 控制器
 */
package api

import (
	"dpm/common"
	"dpm/models"
	"dpm/repositories"
	"fmt"
	"net/http"
	"strconv"
)

var (
	usersRepository = repositories.NewUsersRepository()
)

//用户登录
func Loggin(w http.ResponseWriter, r *http.Request) {

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

	var u models.User
	unmarshal2Object(w, r, &u)
	m := make(map[string]interface{})
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
			m[common.TOKEN_KEY] = token
			jsonResponseOK(w, m)
		}
	}
}

//新增用户
func CreateUser(w http.ResponseWriter, r *http.Request) {

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

	var u models.User
	unmarshal2Object(w, r, &u)
	m := make(map[string]interface{})
	if err := usersRepository.CreateUser(u); err != nil {
		panic(common.ErrTrace(err))
	} else {
		m["rsp_msg"] = "ok"
		jsonResponseOK(w, m)
	}
}

//用户注册
func RegisterUser(w http.ResponseWriter, r *http.Request) {
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

	CreateUser(w, r)
}

/*
* 返回用户列表
 */
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
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

	// vars := mux.SetURLVars(r)
	// vars := r.URL.Query()
	ls := ParseQueryGet(r,"limit")
	pgs := ParseQueryGet(r,"page")
	fmt.Printf("ls:[%s];pgs:[%s]", ls,pgs)

	
	l, _ := strconv.ParseInt(ls, 10, 64)
	pg, _ := strconv.ParseInt(pgs, 10, 64)

	fmt.Println("==================", l, pg)

	p, err := common.NewPageable(l, pg)

	if err != nil {
		panic(common.ErrTrace(err))
	}

	pr, err := usersRepository.GetAllUsers(p)

	if err != nil {
		panic(common.ErrTrace(err))
	}

	jsonResponseOK(w, &pr)
}
