package response

// swagger:model PageResult
type PageResult struct {
	List     interface{} `json:"list"`
	Total    int64       `json:"total"`
	Page     int         `json:"page"`
	PageSize int         `json:"pageSize"`
}
type ResultList struct {
	List []GetAuthorityInfoResult `json:"list"`
}

type GetAuthorityInfoResult struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Type  string `json:"Type"`
	IdInt uint64 `json:"id_int,omitempty"`
}

type ResultPositionList struct {
	List []GetPositionList `json:"list"`
	Desc string            `json:"desc"`
}

type GetPositionList struct {
	Id       uint   `json:"id"`
	Name     string `json:"name"`
	Type     uint   `json:"type"`
	Sort     uint   `json:"sort"`
	Value    uint   `json:"value"`
	Describe string `json:"describe"`
	ParentId uint   `json:"parent_id"`
}

type Department struct {
	// 部门ID
	ID uint64 `json:"id"`
	// 部门名称
	Name string `json:"name"`
	// 是否是主部门
	IsMain bool `json:"is_main"`
}
