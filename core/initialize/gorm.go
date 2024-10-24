package initialize

import (
	"gorm.io/gorm"
	"yxgin/core/global"
)

func Gorm() *gorm.DB {
	switch global.CONFIG.System.DbType {
	case "mysql":
		return GormMysql()
	default:
		return GormMysql()
	}

}
