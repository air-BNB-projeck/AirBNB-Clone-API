package database

import (
	"alta/air-bnb/app/config"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(cfg *config.AppConfig) *gorm.DB {
	connection := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", cfg.DB_USERNAME, cfg.DB_PASS, cfg.DB_HOSTNAME, cfg.DB_PORT, cfg.DB_NAME)
	db, err := gorm.Open(mysql.Open(connection), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}