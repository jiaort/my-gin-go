package config

type System struct {
	Env     string `mapstructure:"env" json:"env" yaml:"env"`
	Addr    int    `mapstructure:"addr" json:"addr" yaml:"addr"`
	OssType string `mapstructure:"oss-type" json:"ossType" yaml:"oss-type"`
	Timeout int64  `mapstructure:"timeout" json:"timeout" yaml:"timeout"`
}
