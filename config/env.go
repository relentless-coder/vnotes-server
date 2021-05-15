package config

import (
	"log"

	"github.com/spf13/viper"
)

type DBConfig struct {
	Dbname   string `mapstructure:"dbname"`
	Password string `mapstructure:"password"`
	User     string `mapstructure:"user"`
}

func SetConfig() {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	log.Printf("hey yooo")
	if err != nil {
		log.Panic(err)
	}
}

func GetDBVars() (config DBConfig, err error) {
	err = viper.Unmarshal(&config)
	return
}
