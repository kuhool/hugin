package request

import (
	"gitlab.yx.cc/front/admin/model/common/request"
	"gitlab.yx.cc/front/admin/model/system"
)

type SysOperationRecordSearch struct {
	system.SysOperationRecord
	request.PageInfo
}
