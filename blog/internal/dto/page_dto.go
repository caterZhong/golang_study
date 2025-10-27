// pkg/pagination/pagination.go
package dto

type Pagination struct {
	PageNo   int `form:"pageNo" json:"page_no" binding:"omitempty,min=1"`
	PageSize int `form:"pageSize" json:"page_size" binding:"omitempty,min=1,max=100"`
}

type PageResponse struct {
	Data  interface{} `json:"data"`
	Total int64       `json:"total"`
	Page  int         `json:"page"`
	Size  int         `json:"size"`
	Pages int         `json:"pages"`
}

func (p *Pagination) GetOffset() int {
	if p.PageNo <= 0 {
		p.PageNo = 1
	}
	if p.PageSize <= 0 {
		p.PageSize = 10
	}
	return (p.PageNo - 1) * p.PageSize
}

func (p *Pagination) GetLimit() int {
	if p.PageSize <= 0 {
		return 10
	}
	return p.PageSize
}

func NewPageResponse(data interface{}, total int64, pagination Pagination) PageResponse {
	pages := int(total) / pagination.PageSize
	if int(total)%pagination.PageSize > 0 {
		pages++
	}

	return PageResponse{
		Data:  data,
		Total: total,
		Page:  pagination.PageNo,
		Size:  pagination.PageSize,
		Pages: pages,
	}
}
