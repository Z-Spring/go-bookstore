package settings

import (
	"bookstore/global"
	"github.com/spf13/viper"
	"time"
)

func GetSettings() error {
	viper.SetConfigFile("configs/config.yml")
	viper.WatchConfig()
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	if err := viper.UnmarshalKey("JWT", &global.JwtSetting); err != nil {
		return err
	}
	global.JwtSetting.Expire *= time.Second
	return nil
}
