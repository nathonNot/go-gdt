package config

import "time"

type IServerConfig interface {
	DefaultConfig()
	GetGameFrameTime() time.Duration
}

type DefaultConfig struct {
	Db          DbConfig     `json:"db"`
	Redis       RedisConfig  `json:"redis"`
	GameRunTime SGameRunTime `json:"gameRunTime"`
}

var localConfig IServerConfig

func (c DefaultConfig) DefaultConfig() {
	localConfig = DefaultConfig{}
}

func (c DefaultConfig) GetGameFrameTime() time.Duration {
	return c.GameRunTime.TimeSpentPerFrame
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
