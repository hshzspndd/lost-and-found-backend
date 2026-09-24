package config

import (
	"log"

	"github.com/spf13/viper"
)

var Config = viper.New()

// 加载配置文件
func LoadConfig() {
	Config.SetConfigName("config")
	Config.SetConfigType("yaml")
	Config.AddConfigPath(".")
	Config.AddConfigPath("./")
	Config.WatchConfig() //支持热更新

	err := Config.ReadInConfig()
	if err != nil {
		log.Fatalf("读取配置文件失败")
	}
}
