package request

import (
	"time"

	"gitlab.yx.cc/front/admin/model/common/request"
	"gitlab.yx.cc/front/admin/model/system"
)

type SysExportTemplateSearch struct {
	system.SysExportTemplate
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	request.PageInfo
}
