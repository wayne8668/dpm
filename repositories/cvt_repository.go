package repositories

import (
	"dpm/common"
	"dpm/models"
	"github.com/juju/errors"
)

type cvtRepository struct{}

var (
	cvtr = &cvtRepository{}
)

//单例构造函数
func NewCVTRepository() *cvtRepository {
	return cvtr
}

//新增简历模板
func (this *cvtRepository) CreateNewCVTemplate(md models.CVTemplate) (numResult int64, err error) {
	conn := GetConn()
	defer conn.Close()
	sqlStr := `
		CREATE (p:CVTEMPLATE {
			cvt_id:{cvt_id}, 
			cvt_no:{cvt_no}, 
			cvt_name:{cvt_name},
			cvt_fmt:{cvt_fmt}, 
			cvt_size:{cvt_size}, 
			cvt_language:{cvt_language}, 
			cvt_color:{cvt_color}, 
			cvt_imgpath:{cvt_imgpath}, 
			cvt_csspath:{cvt_csspath},
			cvt_createtime:{cvt_createtime},
			cvt_updatetime:{cvt_updatetime}})`

	stmt, err := conn.PrepareNeo(sqlStr)
	defer stmt.Close()

	now := common.NowStringFormat()

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

	result, err := stmt.ExecNeo(params)
	if err != nil {
		Logger.Errorf("error:%s", err.Error())
		return 0, err
	}
	numResult, err = result.RowsAffected()
	return numResult, err
}

//修改简历模板
func (this *cvtRepository) UpdateCVTemplate(md models.CVTemplate) error {
	conn := GetConn()
	defer conn.Close()
	sqlStr := `
		MATCH (n:CVTEMPLATE) 
		where n.cvt_id={cvt_id} set n.cvt_no = {cvt_no}, 
		n.cvt_name = {cvt_name},
		n.cvt_fmt = {cvt_fmt}, 
		n.cvt_size = {cvt_size}, 
		n.cvt_language = {cvt_language}, 
		n.cvt_color = {cvt_color}, 
		n.cvt_imgpath = {cvt_imgpath}, 
		n.cvt_csspath = {cvt_csspath},
		n.cvt_updatetime = {cvt_updatetime}`

	now := common.NowStringFormat()

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

	stmt, err := conn.PrepareNeo(sqlStr)
	defer stmt.Close()

	if err != nil {
		err = errors.Trace(err)
		return err
	}
	// result, err := conn.ExecNeo(sqlStr, params)

	if stmt.ExecNeo(params); err != nil {
		err = errors.Trace(err)
		return err
	}

	return err
}
