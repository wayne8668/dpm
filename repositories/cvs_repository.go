package repositories

import (
	"dpm/common"
	"dpm/models"
	"time"
)

type cvsRepository struct{}

var (
	cvr = &cvsRepository{}
)

//单例构造函数
func NewCVSRepository() *cvsRepository {
	return cvr
}

//设置简历模板
func (this *cvsRepository) ReSetCVTemp(uid string, cvid string, cvtid string) (numResult int, err error) {
	conn := GetConn()
	defer conn.Close()
	sqlStr := `match (u:USER) - [:has_cv] ->(cv:curriculum_vitae) - [r1:include_cvt] -> (),(cvt:cv_template) 
			where u.id = {uid} and cv.cv_id = {cvid} and cvt.cvt_id = {cvtid}
			create (cv) - [:include_cvt] -> (cvt) 
			delete r1`

	params := make(map[string]interface{})
	params["uid"] = uid
	params["cvtid"] = cvtid
	params["cvid"] = cvid

	_, err = conn.ExecNeo(sqlStr, params)

	if err := common.ErrInternalServer(err); err != nil {
		return numResult, err
	}

	return numResult, err
}

//新增简历
func (this *cvsRepository) CreateCVWithTemp(uid string, cvtid string) (numResult int, err error) {
	conn := GetConn()
	defer conn.Close()
	sqlStr := `match (u:user),(cvt:cv_template) 
			where u.uid = {uid} and cvt.cvt_id = {cvtid}
			create (u) - [:has_cv] ->(cv:curriculum_vitae{
				cv_id:{cv_id},
				cv_name:{cv_name},
				cv_createtime:{cv_createtime},
				cv_updatetime:{cv_updatetime}}) - [:include_cvt] -> (cvt) `

	now :=time.Now()
	nowNano := time.Now().UnixNano()
	nowFmt := now.Format("060102")

	params := make(map[string]interface{})
	params["uid"] = uid
	params["cvtid"] = cvtid
	params["cv_id"] = NewUUID()
	params["cv_name"] = "我的简历-" + nowFmt
	params["cv_createtime"] = nowNano
	params["cv_updatetime"] = nowNano

	_, err = conn.ExecNeo(sqlStr, params)

	if err := common.ErrInternalServer(err); err != nil {
		return numResult, err
	}

	return 1, err
}

//返回用户的所有简历
func (this *cvsRepository) GetUsersCVS(uid string) (r []models.CurriculumVitae, err error) {
	conn := GetConn()
	defer conn.Close()

	sqlStr := `MATCH (u:user) - [:has_cv] ->(cv:curriculum_vitae)
		where u.uid = {uid} 
		RETURN cv.cv_id,cv.cv_name,cv.cview_pwd,cv.custom_domainname,cv.cvisibili_type,cv.cv_createtime,cv.cv_updatetime
		order by cv.cv_updatetime desc`

	params := make(map[string]interface{})
	params["uid"] = uid

	data, _, _, err := conn.QueryNeoAll(sqlStr, params)

	if err := common.ErrInternalServer(err); err != nil {
		return r, err
	}

	r = make([]models.CurriculumVitae, len(data))

	for idx, row := range data {
		r[idx] = models.CurriculumVitae{
			CVId:             common.NilParseString(row[0]),
			CVName:           common.NilParseString(row[1]),
			CViewPwd:         common.NilParseString(row[2]),
			CustomDomainName: common.NilParseString(row[3]),
			CVisibiliType:    common.NilParseInt(row[4]),
			CVCreateTime:     common.NilParseJSONTime(row[5]),
			CVUpdateTime:     common.NilParseJSONTime(row[6]),
		}
	}
	Logger.Info("GetUsersCVS method return")
	return r, err
}
