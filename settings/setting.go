package settings

import (
	"bookstore/global"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
	"time"
)

func ViperConfig() *viper.Viper {
	v := viper.New()
	//v.SetConfigFile("configs/config.yml")
	v.SetConfigFile("C:\\Users\\Murphy\\GolandProjects\\go-bookstore\\configs\\config.yml")

	if err := v.ReadInConfig(); err != nil {
		log.Println("读取配置文件失败 ", err)
	}

	return v
}

func GetJwtSettings() error {
	v := ViperConfig()
	if err := v.UnmarshalKey("JWT", &global.JwtSetting); err != nil {
		return err
	}
	global.JwtSetting.Expire *= time.Second

	v.WatchConfig()
	v.OnConfigChange(func(in fsnotify.Event) {
		log.Printf("%s was updated", in.Name)
		if err := v.ReadInConfig(); err != nil {
			log.Println("ReadInConfigError: ", err)
		}
		if err := v.UnmarshalKey("JWT", &global.JwtSetting); err != nil {
			log.Println("UnmarshalKeyError: ", err)
		}
		global.JwtSetting.Expire *= time.Second
	})
	return nil
}

func GetServerSettings() error {
	//v := ViperConfig()
	v := ViperConfig()

	if err := v.UnmarshalKey("Server", &global.ServerSetting); err != nil {
		log.Printf("UnmarshalKeyError %s", err)
		return err
	}
	v.WatchConfig()
	v.OnConfigChange(func(in fsnotify.Event) {
		log.Printf("%s was updated", in.Name)
		if err := v.ReadInConfig(); err != nil {
			log.Println("ReadInConfigError: ", err)
		}
		if err := v.UnmarshalKey("Server", &global.ServerSetting); err != nil {
			log.Println("UnmarshalKeyError: ", err)
		}
		global.JwtSetting.Expire *= time.Second
		log.Println("runMode: ", global.ServerSetting.RunMode)
	})

	return nil
}

func GetRateLimiterSettings() error {
	v := ViperConfig()
	if err := v.UnmarshalKey("RateLimiter", &global.RateLimiterSetting); err != nil {
		return err
	}
	// todo: 程序运行时，修改配置文件无法生效，因为程序一直在监听端口
	v.WatchConfig()
	v.OnConfigChange(func(in fsnotify.Event) {
		log.Printf("%s was updated", in.Name)
	})
	return nil
}

func GetContextTimeout() error {
	v := ViperConfig()
	if err := v.UnmarshalKey("ContextTimeout", &global.TimeOutSetting); err != nil {
		return err
	}
	global.TimeOutSetting.TimeOut *= time.Second
	return nil
}
