package repositories

import (
	"dpm/common"
	"dpm/models"
	"time"
)

type cvtRepository struct{}

var (
	cvtr = &cvtRepository{}
)

//单例构造函数
func NewCVTRepository() *cvtRepository {
	return cvtr
}

//判断模板是否存在
func (this *cvtRepository) IsExist(cvtid string) (bool, error) {

	params := make(map[string]interface{})
	params["cvt_id"] = cvtid

	c := NewCypher("cvt_reposity.is_exist.cypher").Params(params)

	var count int64

	err := QueryNeo(func(row []interface{}) {
		count = common.NilParseInt64(row[0])
	}, c)

	if err := common.ErrInternalServer(err); err != nil {
		return false, err
	}

	if count == 0 {
		return false, nil
	}

	return true, nil
}

//新增简历模板
func (this *cvtRepository) CreateNewCVTemplate(md models.CVTemplate) error {

	now := time.Now().UnixNano()

	params := make(map[string]interface{})
	params["cvt_id"] = NewUUID()
	params["cvt_no"] = md.CVTNo
	params["cvt_name"] = md.CVTName
	params["cvt_fmt"] = md.CVTFmt
	params["cvt_size"] = md.CVTSize
	params["cvt_language"] = md.CVTLanguage
	params["cvt_color"] = md.CVTColor
	params["cvt_imgpath"] = md.CVTImgPath
	params["cvt_csspath"] = md.CVTCssPath
	params["cvt_createtime"] = now
	params["cvt_updatetime"] = now

	c := NewCypher("cvt_reposity.create_new_cvt.cypher").Params(params)

	err := ExecNeo(c)

	if err := common.ErrInternalServer(err); err != nil {
		return err
	}

	return nil
}

//修改简历模板
func (this *cvtRepository) UpdateCVTemplate(md models.CVTemplate) error {

	now := time.Now().UnixNano()

	params := make(map[string]interface{})
	params["cvt_id"] = md.CVTId
	params["cvt_no"] = md.CVTNo
	params["cvt_name"] = md.CVTName
	params["cvt_fmt"] = md.CVTFmt
	params["cvt_size"] = md.CVTSize
	params["cvt_language"] = md.CVTLanguage
	params["cvt_color"] = md.CVTColor
	params["cvt_imgpath"] = md.CVTImgPath
	params["cvt_csspath"] = md.CVTCssPath
	params["cvt_updatetime"] = now

	c := NewCypher("cvt_reposity.update_new_cvt.cypher").Params(params)

	err := ExecNeo(c)

	if err := common.ErrInternalServer(err); err != nil {
		return err
	}
	return nil
}

func (this *cvtRepository) GetAllCVTS(p common.Pageable) (common.Pageable, error) {

	c := NewCypher("cvt_reposity.get_all_cvts.cypher_count")

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

	c = NewCypher("cvt_reposity.get_all_cvts.cypher_pageable").Params(params)

	callback = func(row []interface{}) {
		m := &models.CVTemplate{
			CVTId:         common.NilParseString(row[0]),
			CVTNo:         common.NilParseString(row[1]),
			CVTName:       common.NilParseString(row[2]),
			CVTFmt:        common.NilParseString(row[3]),
			CVTSize:       common.NilParseString(row[4]),
			CVTLanguage:   common.NilParseString(row[5]),
			CVTColor:      common.NilParseString(row[6]),
			CVTImgPath:    common.NilParseString(row[7]),
			CVTCssPath:    common.NilParseString(row[8]),
			CVTCreateTime: common.NilParseJSONTime(row[9]),
			CVTUpdateTime: common.NilParseJSONTime(row[10]),
		}
		p.AddContent(m)
	}

	err = QueryNeo(callback, c)

	if err := common.ErrInternalServer(err); err != nil {
		return p, err
	}

	return p, nil
}
