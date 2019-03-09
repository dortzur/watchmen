package config

import (
	"github.com/spf13/viper"
	"watchmen/services"
)

var config *viper.Viper

// Init is an exported method that takes the environment starts the viper
// (external lib) and returns the configuration struct.
func Init(env string) {
	var err error
	config = viper.New()
	config.SetConfigType("yaml")
	config.SetConfigName(env)
	config.AddConfigPath("../config/")
	config.AddConfigPath("config/")
	err = config.ReadInConfig()
	if err != nil {
		services.Logger.Fatal("error on parsing configuration file")
	}
}

func GetConfig() *viper.Viper {
	return config
}
