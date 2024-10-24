package request

import (
	"gitlab.yx.cc/front/admin/model/system"
)

// Register User register structure
type Register struct {
	Username     string `json:"userName" example:"用户名"`
	Password     string `json:"passWord" example:"密码"`
	NickName     string `json:"nickName" example:"昵称"`
	HeaderImg    string `json:"headerImg" example:"头像链接"`
	AuthorityId  uint   `json:"authorityId" swaggertype:"string" example:"int 角色id"`
	Enable       int    `json:"enable" swaggertype:"string" example:"int 是否启用"`
	AuthorityIds []uint `json:"authorityIds" swaggertype:"string" example:"[]uint 角色id"`
	Phone        string `json:"phone" example:"电话号码"`
	Email        string `json:"email" example:"电子邮箱"`
	Type         uint   `json:"type"  example:"uint 用户类型"`
}

// User login structure
type Login struct {
	Username  string `json:"username"`  // 用户名
	Password  string `json:"password"`  // 密码
	Captcha   string `json:"captcha"`   // 验证码
	CaptchaId string `json:"captchaId"` // 验证码ID
}

// Modify password structure
type ChangePasswordReq struct {
	ID          uint   `json:"-"`           // 从 JWT 中提取 user id，避免越权
	Password    string `json:"password"`    // 密码
	NewPassword string `json:"newPassword"` // 新密码
}

// Modify  user's auth structure
type SetUserAuth struct {
	AuthorityId uint `json:"authorityId"` // 角色ID
}

type GetAuthorityId struct {
	AuthorityId uint `json:"authorityId"` // 角色ID
}

// Modify  user's auth structure
type SetUserAuthorities struct {
	ID           uint
	AuthorityIds []uint `json:"authorityIds"` // 角色ID
}

type ChangeUserInfo struct {
	ID           uint                  `gorm:"primarykey"`                                                                                                                    // 主键ID
	NickName     string                `json:"nickName" gorm:"default:系统用户;comment:用户昵称"`                                                                                     // 用户昵称
	Phone        string                `json:"phone"  gorm:"comment:用户手机号"`                                                                                                   // 用户手机号
	AuthorityIds []uint                `json:"authorityIds" gorm:"-"`                                                                                                         // 角色ID
	Email        string                `json:"email"  gorm:"comment:用户邮箱"`                                                                                                    // 用户邮箱
	HeaderImg    string                `json:"headerImg" gorm:"default:/storage/admin/a75b1ab4/073576fd/b8d7e3e8/c0c18c2b/a75b1ab4073576fdb8d7e3e8c0c18c2b.png;comment:用户头像"` // 用户头像
	SideMode     string                `json:"sideMode"  gorm:"comment:用户侧边主题"`                                                                                               // 用户侧边主题
	Enable       int                   `json:"enable" gorm:"comment:冻结用户"`                                                                                                    //冻结用户
	Authorities  []system.SysAuthority `json:"-" gorm:"many2many:sys_user_authority;"`
	Type         uint                  `json:"type"  example:"uint 用户类型"`
}

type SearchUserPageInfo struct {
	Page     int    `json:"page" form:"page"`         // 页码
	PageSize int    `json:"pageSize" form:"pageSize"` // 每页大小
	Keyword  string `json:"keyword" form:"keyword"`   //关键字
	TypeFlag int    `json:"typeFlag" form:"typeFlag"` //关键字
}
