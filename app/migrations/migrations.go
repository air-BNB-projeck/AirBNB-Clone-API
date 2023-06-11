package migrations

import (
	_userData "alta/air-bnb/features/users/data"

	"gorm.io/gorm"
)

func InitMigrate(db *gorm.DB) error {
	err := db.AutoMigrate(&_userData.Users{})
	if err != nil {
		return err
	}

	return nil
}