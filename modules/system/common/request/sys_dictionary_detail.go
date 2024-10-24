package request

import (
	"gitlab.yx.cc/front/admin/model/common/request"
	"gitlab.yx.cc/front/admin/model/system"
)

type SysDictionaryDetailSearch struct {
	system.SysDictionaryDetail
	request.PageInfo
}
