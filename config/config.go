package config

type Server struct {
	Mysql   Mysql   `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Redis   Redis   `mapstructure:"redis" json:"redis" yaml:"redis"`
	System  System  `mapstructure:"system" json:"system" yaml:"system"`
	Zap     Zap     `mapstructure:"zap" json:"zap" yaml:"zap"`
	Captcha Captcha `mapstructure:"captcha" json:"captcha" yaml:"captcha"`
	Local   Local   `mapstructure:"local" json:"local" yaml:"local"`
}
