package repositories

import (
	"dpm/common"
	"dpm/models"
	"encoding/json"
	"fmt"
)

type cvsRepository struct{}

var (
	cvr = &cvsRepository{}
)

//单例构造函数
func NewCVSRepository() *cvsRepository {
	return cvr
}

//返回用户简历
func (this *cvsRepository) GetUsersCVS(uid string) (r []models.CurriculumVitae, err error) {
	conn := GetConn()
	defer conn.Close()

	sqlStr := `MATCH (n:curriculum_vitae) 
		where n.uid = {uid} 
		RETURN n.cv_id,
			n.cv_name,
			n.custom_domain_name,
			n.cvisibili_type,
			n.cv_createtime,
			n.cv_updatetime`

	params := make(map[string]interface{})
	params["uid"] = uid

	data, _, _, err := conn.QueryNeoAll(sqlStr, params)

	if err := common.ErrInternalServer(err); err != nil {
		return r, err
	}

	for idx, row := range data {
		ct, err := common.UMStr2JSONTime(row[4].(string))
		if err := common.ErrInternalServer(err); err != nil {
			return r, err
		}
		ut, err := common.UMStr2JSONTime(row[5].(string))
		if err := common.ErrInternalServer(err); err != nil {
			return r, err
		}
		r[idx] = models.CurriculumVitae{
			CVId:             row[0].(string),
			CVName:           row[1].(string),
			CustomDomainName: row[2].(string),
			CVisibiliType:    row[3].(int),
			CVCreateTime:     ct,
			CVUpdateTime:     ut,
		}
	}
	Logger.Info("GetUsersCVS method return")
	return r, err
}

func (this *cvsRepository) CreateCV(md models.CurriculumVitae) (numResult int, err error) {

	conn := GetConn()

	tx, err := conn.Begin()
	///////////////////////
	q1, m1 := createCVRoot(md)

	stmt, err := conn.PreparePipeline(q1)

	pipelineResults, err := stmt.ExecPipeline(m1)

	fmt.Println(pipelineResults)

	tx.Commit()
	return 0, nil
}

func createCVRoot(m models.CurriculumVitae) (string, map[string]interface{}) {
	sqlCreateCVRoot := `
		create (cv:curriculum_vitae {//新增简历
			cv_id:{cvid}, //简历Id
			cv_name:{cvname}, //简历名称
			cview_pwd:{desc},//查看密码
			custom_domain_name:{customdomainname},//自定义域名
			cvisibili_type:{cvisibilitype},//可见类型
			cv_createtime:{cvcreatetime},//创建时间
			cv_updatetime:{cvupdatetime}//更新时间
			})`
	return sqlCreateCVRoot, nil
}

func Print() {
	cv := models.NewCurriculumVitae()
	b, _ := json.Marshal(cv)
	fmt.Printf("%s\n", b)
}
