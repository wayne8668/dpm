package models

// RegisterUserParam
//
// RegisterUser Param JSON
//
// swagger:model RegisterUserParam
type RegisterUserParam struct {
	Name string `json:"name"`
	Pwd  string `json:"pwd"`
}

// CreateUserParam
//
// CreateUserParam Param JSON
//
// swagger:model CreateUserParam
type CreateUserParam struct {
	Name string `json:"name"`
	Pwd  string `json:"pwd"`
}

// CreateUsersCVSParam
//
// CreateUsersCVS Param JSON
//
// swagger:model CreateUsersCVSParam
type CreateUsersCVSParam struct {
	//简历名称
	CVName string `json:"cv_name"`
	//查看密码
	CViewPwd string `json:"cview_pwd"`
	//自定义域名
	CustomDomainName string `json:"custom_domainname"`
	//可见类型
	CVisibiliType int `json:"cvisibili_type"`
	//简历模板
	CVTemplate `json:"cv_template"`
}

// CreateCVTParam
//
// CreateCVT Param JSON
//
// swagger:model CreateCVTParam
type CreateCVTParam struct {
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
}
