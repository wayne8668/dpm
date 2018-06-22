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
	Name       string          `json:"name"`
	Pwd        string          `json:"pwd"`
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