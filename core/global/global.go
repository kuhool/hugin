package global

import (
	"github.com/songzhibin97/gkit/cache/local_cache"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"yxgin/core/config"
)

var (
	YXGIN  *config.YxGinInfo
	VP     *viper.Viper
	CONFIG *config.Server
	LOG    *zap.Logger
	BD     *gorm.DB

	BlackCache local_cache.Cache
)
