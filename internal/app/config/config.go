package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	App struct {
		Addr string `mapstructure:"addr"`
		Port string `mapstructure:"port"`
	}
}

var _Config Config

// GetConfig ...
func GetConfig() *Config {
	return &_Config
}

func Init() {
	viper.SetEnvPrefix("APP")
	viper.SetConfigType("yml")
	viper.AddConfigPath("cmd/conf")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&_Config)
	if err != nil {
		panic(err)
	}
}
