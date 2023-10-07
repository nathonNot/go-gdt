package config

import (
	"encoding/json"
	"os"
)

func DefaultJsonConfig() IServerConfig {
	// 先判断是否有config/config.json文件
	configPath := "config/config.json"
	_, err := os.Stat(configPath)
	if os.IsNotExist(err) {
		_, err = os.Stat("config.json")
		if !os.IsNotExist(err) {
			configPath = "config.json"
		}
		_, err = os.Stat("/app/config/config.json")
		if !os.IsNotExist(err) {
			configPath = "/app/config/config.json"
		}
	}
	file, err := os.Open(configPath)
	defer file.Close()
	serverConfig := DefaultConfig{}
	serverConfig.DefaultConfig()
	if err != nil {
		filePtr, _ := os.Create(configPath)
		defer filePtr.Close()
		encoder := json.NewEncoder(filePtr)
		err = encoder.Encode(serverConfig)
		return serverConfig
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&serverConfig)
	if err != nil {
		panic(err)
	}
	return serverConfig
}
