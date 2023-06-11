package config

import (
	"log"
	"os"
	"strconv"

	"github.com/spf13/viper"
)

type AppConfig struct {
	DB_USERNAME 			string
	DB_PASS						string
	DB_HOSTNAME				string
	DB_PORT						int
	DB_NAME						string
	JWT_ACCESS_TOKEN	string
}

func ReadEnv() *AppConfig {
	appConfig := AppConfig{}
	
	isRead := false
	if val, found := os.LookupEnv("DB_USERNAME"); found {
		appConfig.DB_USERNAME = val
		isRead = true
	} 
	if val, found := os.LookupEnv("DB_PASS"); found {
		appConfig.DB_PASS = val
		isRead = true
	}
	if val, found := os.LookupEnv("DB_HOSTNAME"); found {
		appConfig.DB_HOSTNAME = val
		isRead = true
	} 
	if val, found := os.LookupEnv("DB_PORT"); found {
		appConfig.DB_PORT, _ = strconv.Atoi(val)
		isRead = true
	}
	if val, found := os.LookupEnv("DB_NAME"); found {
		appConfig.DB_NAME = val
		isRead = true
	}
	if val, found := os.LookupEnv("JWT_ACCESS_TOKEN"); found {
		appConfig.JWT_ACCESS_TOKEN = val
		isRead = true
	}

	if !isRead {
		viper.AddConfigPath(".")
		viper.SetConfigName(".local")
		viper.SetConfigType("env")

		if err := viper.ReadInConfig(); err != nil {
			log.Println("error load config: ", err.Error())
			return nil
		}

		appConfig.DB_USERNAME = viper.Get("DB_USERNAME").(string)
		appConfig.DB_PASS = viper.Get("DB_PASS").(string)
		appConfig.DB_HOSTNAME = viper.Get("DB_HOSTNAME").(string)
		appConfig.DB_PORT, _ = strconv.Atoi(viper.Get("DB_PORT").(string))
		appConfig.DB_NAME = viper.Get("DB_NAME").(string)
		appConfig.JWT_ACCESS_TOKEN = viper.Get("JWT_ACCESS_TOKEN").(string)
	}

	return &appConfig
}