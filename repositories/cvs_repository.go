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
func (this *cvsRepository) ReSetCVTemp(uid string, cvid string, cvtid string) error {

	params := make(map[string]interface{})
	params["uid"] = uid
	params["cvtid"] = cvtid
	params["cvid"] = cvid

	c := NewCypher().Match("(u:user) - [:has_cv] ->(cv:curriculum_vitae) - [r:include_cvt] -> (),(cvt:cv_template)").
		Where("u.id = {uid} and cv.cv_id = {cvid} and cvt.cvt_id = {cvtid}").
		Create("(cv) - [:include_cvt] -> (cvt) ").
		Delete("r").Params(params)

	err := ExecNeo(c)

	if err := common.ErrInternalServer(err); err != nil {
		return err
	}
	return nil
}

//新增简历
func (this *cvsRepository) CreateCVWithTemp(uid string, cvtid string) error {

	now := time.Now()
	nowNano := time.Now().UnixNano()
	nowFmt := now.Format("060102")

	params := make(map[string]interface{})
	params["uid"] = uid
	params["cvtid"] = cvtid
	params["cv_id"] = NewUUID()
	params["cv_name"] = "我的简历-" + nowFmt
	params["cv_createtime"] = nowNano
	params["cv_updatetime"] = nowNano

	c := NewCypher().
		Match("(u:user),(cvt:cv_template)").
		Where("u.uid = {uid} and cvt.cvt_id = {cvtid}").
		Create(`(u) - [:has_cv] ->(cv:curriculum_vitae{
					cv_id:{cv_id},
					cv_name:{cv_name},
					cv_createtime:{cv_createtime},
					cv_updatetime:{cv_updatetime}}) - [:include_cvt] -> (cvt)`).Params(params)

	err := ExecNeo(c)
	if err := common.ErrInternalServer(err); err != nil {
		return err
	}
	return nil
}

//返回用户的所有简历
func (this *cvsRepository) GetUsersCVS(p common.Pageable, uid string) (common.Pageable, error) {

	params := make(map[string]interface{})
	params["uid"] = uid

	var count int64

	c := NewCypher().
		Match("(u:user) - [:has_cv] -> (cv:curriculum_vitae)").
		Where("u.uid = {uid}").
		Return("count(*)").Params(params)

	err := QueryNeo(func(row []interface{}) {
		count = common.NilParseInt64(row[0])
	}, c)

	if err := common.ErrInternalServer(err); err != nil {
		return p, err
	}

	Logger.Info("row count is:", count)

	if count == 0 {
		return p, nil
	}
	p.SetTotalElements(count)

	params["limit"] = p.PageSize
	params["offset"] = p.GetOffSet()

	c = NewCypher().
		Match("(u:user) - [:has_cv] ->(cv:curriculum_vitae)").
		Where("u.uid = {uid}").
		Return("cv.cv_id,cv.cv_name,cv.cview_pwd,cv.custom_domainname,cv.cvisibili_type,cv.cv_createtime,cv.cv_updatetime").
		OrderBy("cv.cv_updatetime desc").
		Skip("{offset}").
		Limit("{limit}").Params(params)

	QueryNeo(func(row []interface{}) {
		m := &models.CurriculumVitae{
			CVId:             common.NilParseString(row[0]),
			CVName:           common.NilParseString(row[1]),
			CViewPwd:         common.NilParseString(row[2]),
			CustomDomainName: common.NilParseString(row[3]),
			CVisibiliType:    common.NilParseInt(row[4]),
			CVCreateTime:     common.NilParseJSONTime(row[5]),
			CVUpdateTime:     common.NilParseJSONTime(row[6]),
		}
		p.AddContent(m)
	}, c)

	Logger.Info("GetUsersCVS method return")

	return p, err
}
