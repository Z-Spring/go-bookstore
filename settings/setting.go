package settings

import (
	"bookstore/global"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"time"
)

func GetJwtSettings() error {
	viper.SetConfigFile("configs/config.yml")
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	if err := viper.UnmarshalKey("JWT", &global.JwtSetting); err != nil {
		return err
	}
	global.JwtSetting.Expire *= time.Second

	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Printf("%s was updated", in.Name)
		if err := viper.ReadInConfig(); err != nil {
			fmt.Println("ReadInConfigError: ", err)
		}
		if err := viper.UnmarshalKey("JWT", &global.JwtSetting); err != nil {
			fmt.Println("UnmarshalKeyError: ", err)
		}
		global.JwtSetting.Expire *= time.Second
	})
	return nil
}

func GetServerSettings() error {
	viper.SetConfigFile("configs/config.yml")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("ReadInConfigError %s", err)
		return err
	}
	if err := viper.UnmarshalKey("Server", &global.ServerSetting); err != nil {
		fmt.Printf("UnmarshalKeyError %s", err)
		return err
	}
	return nil
}
