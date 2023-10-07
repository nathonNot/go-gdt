package config

type IServerConfig interface {
	DefaultConfig()
}

type DefaultConfig struct {
	Db    DbConfig    `json:"db"`
	Redis RedisConfig `json:"redis"`
}

var localConfig IServerConfig

func (c DefaultConfig) DefaultConfig() {
	localConfig = DefaultConfig{}
}

func GetGlobalConfig() IServerConfig {
	return localConfig
}

func SetConfig(config IServerConfig) {
	localConfig = config
}

type SerializationConfig func() IServerConfig

func InitServerConfig(configFrom SerializationConfig) {
	c := configFrom()
	SetConfig(c)
}
