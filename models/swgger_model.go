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
