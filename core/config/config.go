package config

type Server struct {
	YXGIN   YxGinInfo `mapstructure:"yxGin" json:"yxGin" yaml:"yxGin"`
	Zap     Zap       `mapstructure:"zap" json:"zap" yaml:"zap"`
	System  System    `mapstructure:"system" json:"system" yaml:"system"`
	Mysql   Mysql     `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Captcha Captcha   `mapstructure:"captcha" json:"captcha" yaml:"captcha"`
	JWT     JWT       `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
}
