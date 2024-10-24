package request

import "encoding/json"

// PageInfo Paging common input parameter structure
// swagger:model PageInfo
type PageInfo struct {
	Page     int    `json:"page" form:"page"`         // 页码
	PageSize int    `json:"pageSize" form:"pageSize"` // 每页大小
	Keyword  string `json:"keyword" form:"keyword"`   //关键字
}

// 获取正确的page信息
func (i PageInfo) GetPage() int {
	if i.Page <= 0 {
		return 1
	}
	return i.Page
}

// 过滤后的page size
func (i PageInfo) GetSize() int {
	if i.PageSize <= 0 {
		return 10
	}
	return i.PageSize
}

// 计算数据库使用的Offset
func (i PageInfo) Offset() int {
	return (i.GetPage() - 1) * i.GetSize()
}

type UserList struct {
	Page     int    `json:"page" form:"page"`         // 页码
	PageSize int    `json:"pageSize" form:"pageSize"` // 每页大小
	Keyword  string `json:"keyword" form:"keyword"`   //关键字
}

// GetById Find by id structure
type GetById struct {
	ID int `json:"id" form:"id"` // 主键ID
}

func (r *GetById) Uint() uint {
	return uint(r.ID)
}

type IdsReq struct {
	Ids []int `json:"ids" form:"ids"`
}

// GetAuthorityId Get role by id structure
type GetAuthorityId struct {
	AuthorityId uint `json:"authorityId" form:"authorityId"` // 角色ID
}

type GetAid struct {
	AuthorityId uint `json:"authorityId" form:"authorityId"` // 角色ID
}

type Empty struct{}

func NewStrListFromJSON(s string) (sl StrList) {
	json.Unmarshal([]byte(s), &sl)
	return
}

type StrList []string

func (sl StrList) ToJSON() (str []byte) {
	str, err := json.Marshal(sl)
	if err != nil {
		return str
	}
	return str
}

func (sl StrList) IsEmpty() bool {
	return len(sl) == 0
}

type Attachment struct {
	Url   string `json:"url"`
	Title string `json:"title"`
	Size  uint   `json:"size,omitempty"`
}

type AttachmentList []Attachment

func (al *AttachmentList) Scan(b []byte) error {
	return json.Unmarshal(b, &al)
}

func (al *AttachmentList) String() string {
	b, err := json.Marshal(al)
	if err != nil {
		return ""
	}
	return string(b)
}

func (al *AttachmentList) IsEmpty() bool {
	return len(*al) == 0
}
