package repositories

import (
	"dpm/common"
	"dpm/models"
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
	conn := GetConn()
	defer conn.Close()
	sqlStr := `match (n:USER) where n.name={name} and n.pwd={pwd} return n.id,n.name,n.pwd`

	params := make(map[string]interface{})
	params["name"] = u.Name
	params["pwd"] = u.Pwd

	data, _, _, err := conn.QueryNeoAll(sqlStr, params)

	if err := common.ErrInternalServer(err); err != nil {
		return udb, err
	}

	if len(data) == 0 {
		return udb, err
	}

	results := make([]*models.User, len(data))
	for idx, row := range data {
		results[idx] = &models.User{
			Id:   row[0].(string),
			Name: row[1].(string),
			Pwd:  row[2].(string),
		}
	}
	udb = results[0]
	return udb, err
}

//用户是否存在
func (this *UsersRepository) IsExist(u models.User) (isExist bool, err error) {

	conn := GetConn()
	defer conn.Close()
	sqlStr := `match (n:USER) where n.name={name} return n`

	params := make(map[string]interface{})
	params["name"] = u.Name

	data, _, _, err := conn.QueryNeoAll(sqlStr, params)

	if err := common.ErrInternalServer(err); err != nil {
		return isExist, err
	}

	if len(data) == 0 {
		return isExist, err
	}
	return true, err
}

//新增用户
func (this *UsersRepository) CreateUser(u models.User) (numResult int64, err error) {

	isExist, err := this.IsExist(u)

	if err := common.ErrInternalServer(err); err != nil {
		return numResult, err
	}

	if isExist {
		err = common.ErrForbiddenf("user name [%s] is exist..", u.Name)
		return numResult, err
	}

	connStr := "CREATE (n:USER {id:{id}, name:{name},pwd:{pwd},create_time:{create_time}})"

	m := map[string]interface{}{
		"id":          NewUUID(),
		"name":        u.Name,
		"pwd":         u.Pwd,
		"create_time": common.NowStringFormat(),
	}

	conn := GetConn()
	defer conn.Close()

	result, err := conn.ExecNeo(connStr, m)

	if err := common.ErrInternalServer(err); err != nil {
		return numResult, err
	}

	numResult, err = result.RowsAffected()

	if err := common.ErrInternalServer(err); err != nil {
		return numResult, err
	}

	return numResult, err
}

//返回所有用户
func (this *UsersRepository) GetAllUsers(p common.Pageable) (common.Pageable, error) {
	conn := GetConn()
	defer conn.Close()
	sqlStrCount := `MATCH (n:USER) RETURN count(*)`

	rows, err := conn.QueryNeo(sqlStrCount, nil)
	defer rows.Close()

	if err := common.ErrInternalServer(err); err != nil {
		return p, err
	}

	nextDate, _, err := rows.NextNeo()
	rows.Close()

	if err := common.ErrInternalServer(err); err != nil {
		return p, err
	}

	count := nextDate[0].(int64)

	if err := common.ErrInternalServer(err); err != nil {
		return p, err
	}

	Logger.Info("row count is:", count)
	p.SetTotalElements(count)

	sqlStr := `MATCH (n:USER) RETURN n.id,n.name,n.pwd SKIP {offset} LIMIT {limit}`

	params := make(map[string]interface{})
	params["limit"] = p.PageSize
	params["offset"] = p.GetOffSet()

	data, _, _, err := conn.QueryNeoAll(sqlStr, params)

	if err := common.ErrInternalServer(err); err != nil {
		return p, err
	}

	// results := make([]*models.User, len(data))

	for _, row := range data {
		m := &models.User{
			Id:   row[0].(string),
			Name: row[1].(string),
			Pwd:  row[2].(string),
		}
		p.AddContent(m)
	}

	return p, err
}
