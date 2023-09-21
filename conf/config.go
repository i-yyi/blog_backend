package config

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var Conf *Config

func InitConfig() {
	viper.AutomaticEnv()
	viper.SetConfigType("yml")
	viper.AddConfigPath("./conf/")
	viper.SetConfigName("config.bj4.online")
	log.Info("Init Config")
	Conf = &Config{}
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := viper.Unmarshal(Conf); err != nil {
		panic(err)
	}
}
func GetConf() *Config {
	if Conf == nil {
		InitConfig()
	}
	return Conf
}
