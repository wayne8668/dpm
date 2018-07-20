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
func (this *cvsRepository) ReSetCVTemp(uid string, cvid string, cvtid string) (err error) {

	params := make(map[string]interface{})
	params["uid"] = uid
	params["cvtid"] = cvtid
	params["cvid"] = cvid

	c := NewCypher("cvs_reposity.reset_cvtemp.cypher").Params(params)

	err = ExecNeo(c)

	if err := common.ErrInternalServer(err); err != nil {
		return err
	}
	return nil
}

//新增简历
func (this *cvsRepository) CreateCVWithTemp(uid string, cvtid string) (cvid string, err error) {

	now := time.Now()
	nowNano := time.Now().UnixNano()
	nowFmt := now.Format("060102")

	cvid = NewUUID()

	params := make(map[string]interface{})
	params["uid"] = uid
	params["cvtid"] = cvtid
	params["cv_id"] = cvid
	params["cv_name"] = "我的简历-" + nowFmt
	params["cv_createtime"] = nowNano
	params["cv_updatetime"] = nowNano

	c := NewCypher("cvs_reposity.createcv_with_temp.cypher").Params(params)

	err = ExecNeo(c)
	if err := common.ErrInternalServer(err); err != nil {
		return "", err
	}
	return cvid, nil
}

//返回用户的所有简历
func (this *cvsRepository) GetUsersCVS(p common.Pageable, uid string) (_ common.Pageable, err error) {

	params := make(map[string]interface{})
	params["uid"] = uid

	var count int64

	c := NewCypher("cvs_reposity.get_users_cvs.cypher_count").Params(params)

	err = QueryNeo(func(row []interface{}) {
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

	c = NewCypher("cvs_reposity.get_users_cvs.cypher_pageable").Params(params)

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
