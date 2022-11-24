package main

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Port       string
	Connection string
}

var appConf *Config

func loadConf() {
	log.Println("Loading server config ...")
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	err := viper.ReadInConfig()

	if err != nil {
		log.Fatal(err)
	}
	err = viper.Unmarshal(&appConf)

	if err != nil {
		log.Fatal(err)
	} else {
		log.Println(appConf)
	}

}
