package models

import (
	"dpm/common"
)

// CVTemplate
//
// cv template struct
//
// swagger:model CVTemplate
type CVTemplate struct {
	
	// 模板id
    //
    //swagger:ignore
	CVTId         string          `json:"cvt_id"`
	//模板编号
	CVTNo         string          `json:"cvt_no"`
	//模板名称
	CVTName       string          `json:"cvt_name"`
	//支持格式
	CVTFmt        string          `json:"cvt_fmt"`
	//模板尺寸
	CVTSize       string          `json:"cvt_size"`
	//模板语言
	CVTLanguage   string          `json:"cvt_language"`
	//模板颜色
	CVTColor      string          `json:"cvt_color"`
	//模板图片路径
	CVTImgPath    string          `json:"cvt_imgpath"`
	//模板css路径
	CVTCssPath    string          `json:"cvt_csspath"`
	//创建时间
	//swagger:ignore
	CVTCreateTime common.JSONTime `json:"cvt_createtime"`
	//修改时间
	//swagger:ignore
	CVTUpdateTime common.JSONTime `json:"cvt_updatetime"`
}