package migrations

import (
	_stayData "alta/air-bnb/features/stays/data"
	_userData "alta/air-bnb/features/users/data"

	"gorm.io/gorm"
)

func InitMigrate(db *gorm.DB) error {
	err := db.AutoMigrate(&_userData.Users{}, &_stayData.Stays{}, &_stayData.StayImages{})
	if err != nil {
		return err
	}

	return nil
}