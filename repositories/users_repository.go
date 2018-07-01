package repositories

import (
	"dpm/common"
	"dpm/models"
	"time"
	// "github.com/goinggo/mapstructure"
)

type UsersRepository struct{}

var (
	ur = &UsersRepository{}
)

//单例构造函数
func NewUsersRepository() *UsersRepository {
	return ur
}

func checkErr(err error) error {
	if err != nil {

	}
	return nil
}

//返回用户信息
func (this *UsersRepository) GetUserForAuth(u models.User) (udb *models.User, err error) {

	params := make(map[string]interface{})
	params["name"] = u.Name
	params["pwd"] = u.Pwd

	c := NewCypher().Match("(n:user) where n.name={name} and n.pwd={pwd}").Return("n.uid,n.name,n.pwd").Params(params)

	callback := func(row []interface{}) {
		udb = &models.User{
			UId:  common.NilParseString(row[0]),
			Name: common.NilParseString(row[1]),
			Pwd:  common.NilParseString(row[2]),
		}
	}

	err = QueryNeo(callback, c)

	if err := common.ErrInternalServer(err); err != nil {
		return nil, err
	}
	return udb, err
}

//用户是否存在
func (this *UsersRepository) IsExist(u models.User) (bool, error) {

	params := make(map[string]interface{})
	params["name"] = u.Name

	c := NewCypher().Match("(n:user)").Where("n.name={name}").Return("n").Params(params)

	var rowNum int64

	callback := func(row []interface{}) {
		rowNum++
	}

	err := QueryNeo(callback, c)

	if err := common.ErrInternalServer(err); err != nil {
		return false, err
	}

	if rowNum == 0 {
		return false, nil
	}

	return true, nil
}

//新增用户
func (this *UsersRepository) CreateUser(u models.User) error {

	isExist, err := this.IsExist(u)

	if err := common.ErrInternalServer(err); err != nil {
		return err
	}

	if isExist {
		err = common.ErrForbiddenf("user name [%s] is exist..", u.Name)
		return err
	}

	m := map[string]interface{}{
		"uid":         NewUUID(),
		"name":        u.Name,
		"pwd":         u.Pwd,
		"create_time": time.Now().UnixNano(),
	}

	c := NewCypher().Create("(n:user {uid:{uid}, name:{name},pwd:{pwd},create_time:{create_time}})").Params(m)

	err = ExecNeo(c)

	if err := common.ErrInternalServer(err); err != nil {
		return err
	}

	return nil
}

//返回所有用户
func (this *UsersRepository) GetAllUsers(p common.Pageable) (common.Pageable, error) {

	c := NewCypher().Match("(n:user)").Return("count(*)")

	var count int64

	callback := func(row []interface{}) {
		count = common.NilParseInt64(row[0])
	}

	err := QueryNeo(callback, c)

	if err := common.ErrInternalServer(err); err != nil {
		return p, err
	}

	Logger.Info("row count is:", count)
	if count == 0 {
		return p, nil
	}
	p.SetTotalElements(count)

	params := make(map[string]interface{})
	params["limit"] = p.PageSize
	params["offset"] = p.GetOffSet()

	c = NewCypher().Match("(n:user)").Return("n.uid,n.name,n.pwd,n.create_time").Skip("{offset}").Limit("{limit}").Params(params)

	callback = func(row []interface{}) {
		m := &models.User{
			UId:        common.NilParseString(row[0]),
			Name:       common.NilParseString(row[1]),
			Pwd:        common.NilParseString(row[2]),
			CreateTime: common.NilParseJSONTime(row[3]),
		}
		p.AddContent(m)
	}

	err = QueryNeo(callback, c)

	if err := common.ErrInternalServer(err); err != nil {
		return p, err
	}

	return p, nil
}
