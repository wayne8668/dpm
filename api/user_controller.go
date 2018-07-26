/*
* User 控制器
 */
package api

import (
	"dpm/common"
	"dpm/middleware"
	"dpm/models"
	"dpm/repositories"
)

var (
	usersRepository = repositories.NewUsersRepository()
)

//用户登录
func Loggin(req *common.ApiRequest) (rsp common.ApiRsponse) {
	var u models.User
	if err := req.BindJSON(&u); err != nil {
		return rsp.Error(common.ErrTrace(err))
	}
	if udb, err := usersRepository.GetUserForAuth(u); err != nil {
		return rsp.Error(common.ErrTrace(err))
	} else {
		if udb.UId == "" {
			return rsp.Error(common.ErrForbidden("user name or pwd err."))
		}

		ut := &middleware.UserToken{
			Name: udb.Name,
			Pwd:  udb.Pwd,
			Id:   udb.UId,
		}
		if token, err := middleware.CreateToken(ut); err != nil {
			return rsp.Error(common.ErrTrace(err))
		} else {
			return rsp.AddAttribute(middleware.TOKEN_KEY, token).AddAttribute("auth_user", udb)
		}
	}
}

//新增用户
func CreateUser(req *common.ApiRequest) (rsp common.ApiRsponse) {
	var u models.User
	if err := req.BindJSON(&u); err != nil {
		return rsp.Error(common.ErrTrace(err))
	}
	if err := usersRepository.CreateUser(u); err != nil {
		return rsp.Error(common.ErrTrace(err))
	}
	return
}

//用户注册
func RegisterUser(req *common.ApiRequest) (rsp common.ApiRsponse) {
	return CreateUser(req)
}

/*
* 返回用户列表
 */
func GetAllUsers(req *common.ApiRequest) (rsp common.ApiRsponse) {
	var pr PageableRequest
	if err := req.BindStruct(&pr); err != nil {
		return rsp.Error(common.ErrTrace(err))
	}
	p, err := common.NewPageable(pr.Limit, pr.Page)
	if err != nil {
		return rsp.Error(common.ErrTrace(err))
	}
	p, err = usersRepository.GetAllUsers(p)
	return rsp.Error(err).AddObject(p)
}
