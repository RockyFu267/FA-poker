package config

type Wechat struct {
	Appid  string `mapstructure:"appid" json:"appid" yaml:"appid"`    // 小程序的appid
	Secret string `mapstructure:"secret" json:"secret" yaml:"secret"` // 小程序的secret
	Url    string `mapstructure:"url" json:"url" yaml:"url"`          // 请求微信服务器的url
}
