package page

type PageRequest struct {
	PageNo   int `json:"pageNo"`
	PageSize int `json:"pageSize"`
}

func (r *PageRequest) GetPageNo() int {
	if r.PageNo <= 0 {
		return 1
	}
	return r.PageNo
}

func (r *PageRequest) GetPageSize() int {
	if r.PageSize <= 0 {
		return 10
	}
	return r.PageSize
}

func (r *PageRequest) GetOffset() int {
	return (r.GetPageNo() - 1) * r.GetPageSize()
}

type PageData[T any] struct {
	List     []T `json:"list"`
	Total    int `json:"total"`
	PageNo   int `json:"pageNo"`
	PageSize int `json:"pageSize"`
}

// OfData create a new PageData
func OfData[T any](list []T, total, pageNo, pageSize int) *PageData[T] {
	return &PageData[T]{
		List:     list,
		Total:    total,
		PageNo:   pageNo,
		PageSize: pageSize,
	}
}

// OfRequest create a new PageRequest
func OfRequest(pageNo, pageSize int) *PageRequest {
	return &PageRequest{
		PageNo:   pageNo,
		PageSize: pageSize,
	}
}
