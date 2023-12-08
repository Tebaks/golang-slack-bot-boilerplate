package configs

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
)

const (
	configFileName = "configs"
	secretFileName = "secrets"
)

var (
	AppConfig      = &Configs{}
	AppCredentials = &Credentials{}
)

func Init() {
	viper.AutomaticEnv()
	vConfig := ReadWithViper(configFileName, AppConfig)
	vConfig.WatchConfig()
	vConfig.OnConfigChange(func(in fsnotify.Event) {
		ReadWithViper(configFileName, AppConfig)
		log.Println("Application configs are changed")
	})

	vSecret := ReadWithViper(secretFileName, AppCredentials)
	vSecret.WatchConfig()
	vSecret.OnConfigChange(func(in fsnotify.Event) {
		ReadWithViper(secretFileName, AppCredentials)
		log.Println("Application secrets are changed")
	})
}

func ReadWithViper(name string, object interface{}) *viper.Viper {
	v := viper.New()
	v.AddConfigPath("./configs")
	v.SetConfigName(name)
	v.SetConfigType("yml")

	err := v.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = v.Unmarshal(object)
	if err != nil {
		panic(err)
	}

	return v
}
