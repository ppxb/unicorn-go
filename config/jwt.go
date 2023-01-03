package config

type Jwt struct {
	SecretKey string `mapstructure:"secret-key" json:"secret-key"`
	ExpireAt  string `mapstructure:"expire-at" json:"expire-at"`
}
