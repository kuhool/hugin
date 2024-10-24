package request

import "time"

type ReqTurnTaskPageInfo struct {
	Page     int    `json:"page" form:"page"`         // 页码
	PageSize int    `json:"pageSize" form:"pageSize"` // 每页大小
	Keyword  string `json:"keyword" form:"keyword"`   //关键字
}

type ReqTurnTaskPageInfoByDocId struct {
	Page     int    `json:"page" form:"page"`         // 页码
	PageSize int    `json:"pageSize" form:"pageSize"` // 每页大小
	Keyword  string `json:"keyword" form:"keyword"`   //关键字
	DocID    int    `json:"doc_id" form:"doc_id"`     //关键字
	DocType  uint   `json:"doc_type" form:"doc_type"`
}

type ReqTurnStatusByDocId struct {
	DocID   int  `json:"doc_id" form:"doc_id"` //关键字
	DocType uint `json:"doc_type" form:"doc_type"`
}

type ReqTurnTaskInfo struct {
	ID      uint `json:"id"`
	Confirm uint `json:"confirm"`
}

type ReqTurnTask struct {
	DocId      uint   `json:"doc_id"`
	DocType    uint   `json:"doc_type"`
	Identifier string `json:"identifier"`

	Title string `json:"title"`

	ContentAbstract string `json:"content_abstract"`

	DocumentNumber uint `json:"document_number"`
	SenderId       uint `json:"sender_id"`

	SenderCompanyId uint `json:"sender_company_id"`

	SendDate time.Time `json:"send_date"`

	ReceiverId uint `json:"receiver_id"`
}

type ResponseTurnTaskList struct {
	List  []TurnTask `json:"list"`
	Total int64      `json:"total"`
}

type ResponseTurnStatus struct {
	Total int64 `json:"total"`
}

type ResponseTurnTaskDetail struct {
	Detail TurnTask `json:"detail"`
}

type TurnTask struct {
	ID uint `json:"ID"`

	DocId           uint   `json:"doc_id"`
	DocType         uint   `json:"doc_type"`
	Identifier      string `json:"identifier"`
	Title           string `json:"title"`
	ContentAbstract string `json:"content_abstract"`

	SenderId        uint      `json:"sender_id"`
	SenderName      string    `json:"sender_name"`
	SenderCompanyId uint      `json:"sender_company_id" gorm:"type:bigint(20);default:0;comment:发起者公司ID"`
	SendDate        time.Time `json:"send_date"`

	ReceiverId          uint      `json:"receiver_id"`
	ReceiverName        string    `json:"receiver_name"`
	ReceiverCompanyId   uint      `json:"receiver_company_id"`
	ReceiverConfirm     uint      `json:"receiver_confirm"`
	ReceiverConfirmDate time.Time `json:"receiver_confirm_date"`
}
