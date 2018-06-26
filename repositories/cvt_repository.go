package repositories

import (
	"dpm/common"
	"dpm/models"
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
		CREATE (p:cv_template {
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
	if err := common.ErrInternalServer(err); err != nil {
		return numResult, err
	}
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
	if err := common.ErrInternalServer(err); err != nil {
		return numResult, err
	}
	numResult, err = result.RowsAffected()
	if err := common.ErrInternalServer(err); err != nil {
		return numResult, err
	}
	return numResult, err
}

//修改简历模板
func (this *cvtRepository) UpdateCVTemplate(md models.CVTemplate) error {
	conn := GetConn()
	defer conn.Close()

	sqlStr := `
		MATCH (n:cv_template) 
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

	if err := common.ErrInternalServer(err); err != nil {
		return err
	}
	_, err = stmt.ExecNeo(params)

	if err := common.ErrInternalServer(err); err != nil {
		return err
	}
	return err
}

func (this *cvtRepository) GetAllCVTS(p common.Pageable) (common.Pageable, error) {
	conn := GetConn()
	defer conn.Close()
	sqlStrCount := `MATCH (n:cv_template) RETURN count(*)`

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

	sqlStr := `MATCH (n:cv_template) 
		RETURN 
			n.cvt_id,
			n.cvt_no,
			n.cvt_name,
			n.cvt_fmt,
			n.cvt_size,
			n.cvt_language,
			n.cvt_color,
			n.cvt_imgpath,
			n.cvt_csspath,
			n.cvt_createtime,
			n.cvt_updatetime
		SKIP {offset} 
		LIMIT {limit}`

	params := make(map[string]interface{})
	params["limit"] = p.PageSize
	params["offset"] = p.GetOffSet()

	data, _, _, err := conn.QueryNeoAll(sqlStr, params)

	if err := common.ErrInternalServer(err); err != nil {
		return p, err
	}

	// results := make([]*models.User, len(data))

	for _, row := range data {
		c, _ := common.UMStr2JSONTime(row[9].(string))
		u, _ := common.UMStr2JSONTime(row[10].(string))
		m := &models.CVTemplate{
			CVTId:         row[0].(string),
			CVTNo:         row[1].(string),
			CVTName:       row[2].(string),
			CVTFmt:        row[3].(string),
			CVTSize:       row[4].(string),
			CVTLanguage:   row[5].(string),
			CVTColor:      row[6].(string),
			CVTImgPath:    row[7].(string),
			CVTCssPath:    row[8].(string),
			CVTCreateTime: c,
			CVTUpdateTime: u,
		}
		p.AddContent(m)
	}

	return p, err
}
