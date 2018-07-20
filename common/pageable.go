package common

import (
	"math"
)

// Pageable
//
// Pageable JSON Struct
//
// swagger:model Pageable
type Pageable struct {
	//查询结果列表
	Content []interface{} `json:"content"`
	//是否为最后一页
	IsLastPage bool `json:"is_last_page"`
	//总页数
	TotolPages int64 `json:"totol_pages"`
	//总记录数
	TotalElements int64 `json:"total_elements"`
	//每页记录数
	PageSize int64 `json:"page_size"`
	//当前为第几页
	NumberOfPage int64 `json:"number_of_page"`
	//是否为第一页
	IsFirstPage bool `json:"is_first_page"`
	//当前页记录数
	NumberOfElements int64 `json:"number_of_elements"`
}

//实例化Pageable
func NewPageable(pageSize int64, pageNumber int64) (p Pageable, err error) {

	if pageSize == 0 {
		return p, ErrBadRequestf("Bad Request Args:[pageSize],the value is:[%d]", pageSize)
	}

	p = Pageable{
		PageSize:     pageSize,
		NumberOfPage: pageNumber,
	}

	if pageNumber <= 0 {
		p.NumberOfPage = 1
	}

	if p.NumberOfPage == 1 {
		p.IsFirstPage = true
	}

	return p, err
}

func (this *Pageable) SetTotalElements(t int64) {

	this.TotalElements = t

	this.TotolPages = int64(math.Ceil(float64(t) / float64(this.PageSize)))

	if this.TotolPages == 0 {
		this.IsFirstPage = true
		this.IsLastPage = true
	}

	if this.NumberOfPage > this.TotolPages {
		this.NumberOfPage = this.TotolPages
	}

	if this.NumberOfPage == this.TotolPages {
		this.IsLastPage = true
	}

	if this.NumberOfPage == 1 {
		this.IsFirstPage = true
	}

}

func (this *Pageable) GetOffSet() int64 {
	return this.PageSize * (this.NumberOfPage - 1)
}

func (this *Pageable) AddContent(c interface{}) {
	this.Content = append(this.Content, c)
	this.NumberOfElements = int64(len(this.Content))
}
