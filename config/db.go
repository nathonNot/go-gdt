package config

type DbConfig struct {
	DbUrl string `json:"db_url,omitempty"`
}

type RedisConfig struct {
	RedisAddr string `json:"redis_addr"`
	RedisPass string `json:"redis_pass"`
	RedisUser string `json:"redis_user"`
}
