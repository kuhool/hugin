package core

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"yxgin/core/global"
)

var configFile = "./core/config.yaml"

func Viper(configPath ...string) *viper.Viper {
	config := configFile
	if len(configPath) > 0 {
		config = configPath[0]
	}
	v := viper.New()
	v.SetConfigFile(config)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err = v.Unmarshal(&global.CONFIG); err != nil {
			fmt.Println(err)
		}
		fmt.Println(global.CONFIG.YXGIN)
	})
	if err = v.Unmarshal(&global.CONFIG); err != nil {
		panic(err)
	}
	return v
}
