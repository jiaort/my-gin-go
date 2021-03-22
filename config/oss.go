package config

type Local struct {
	Path   string `mapstructure:"path" json:"path" yaml:"path" `
	Result string `mapstructure:"result" json:"result" yaml:"result" `
}
