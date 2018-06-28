package models

import (
	"dpm/common"
)

//简历模板
type CVTemplate struct {

	// 模板id
	CVTId string `json:"cvt_id"`
	//模板编号
	CVTNo string `json:"cvt_no"`
	//模板名称
	CVTName string `json:"cvt_name"`
	//支持格式
	CVTFmt string `json:"cvt_fmt"`
	//模板尺寸
	CVTSize string `json:"cvt_size"`
	//模板语言
	CVTLanguage string `json:"cvt_language"`
	//模板颜色
	CVTColor string `json:"cvt_color"`
	//模板图片路径
	CVTImgPath string `json:"cvt_imgpath"`
	//模板css路径
	CVTCssPath string `json:"cvt_csspath"`
	//是否上架
	OnLine bool `json:"on_line"`
	//创建时间
	CVTCreateTime common.JSONTime `json:"cvt_createtime"`
	//修改时间
	CVTUpdateTime common.JSONTime `json:"cvt_updatetime"`
}
