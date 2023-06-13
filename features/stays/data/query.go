package data

import (
	"alta/air-bnb/app/helper"
	"alta/air-bnb/features/stays"
	"fmt"

	"gorm.io/gorm"
)

type StayData struct {
	db *gorm.DB
}

// Insert implements stays.StayDataInterface
func (repo *StayData) Insert(stayData stays.CoreStayRequest) (stayId string, err error) {
	var id = helper.GenerateNewId()
	var stayCoreMap = CoreStayToModel(stayData)
	stayCoreMap.ID = id
	if tx := repo.db.Create(&stayCoreMap); tx.Error != nil {
		return "", tx.Error
	}
	return stayCoreMap.ID, nil
}

// SelectAll implements stays.StayDataInterface
func (repo *StayData) SelectAll() (allStays []stays.Core, err error) {
	var staysData []Stays
	if tx := repo.db.Preload("User").Find(&staysData); tx.Error != nil {
		return nil, tx.Error
	}
	fmt.Println(staysData)
	var staysCoreMap []stays.Core
	for _, stay := range staysData {
		stayMap := ModelStayToCore(stay)
		stayMap.User = stays.Users{
			ID: stay.User.ID,
			FullName: stay.User.FullName,
			Email: stay.User.Email,
		}
		staysCoreMap = append(staysCoreMap, stayMap)
	}
	return staysCoreMap, nil
}

// Select implements stays.StayDataInterface
func (repo *StayData) Select(stayId string) (stay stays.Core, err error) {
	var stayData Stays
	if tx := repo.db.Where("id = ?", stayId).Preload("User").First(&stayData); tx.Error != nil {
		return stays.Core{}, tx.Error
	}
	var stayDataMap = ModelStayToCore(stayData)
	stayDataMap.User = stays.Users{
		ID: stayData.User.ID,
		FullName: stayData.User.FullName,
		Email: stayData.User.Email,
	}
	return stayDataMap, nil
}

// Update implements stays.StayDataInterface
func (repo *StayData) Update(stayId string, stayData stays.CoreStayRequest) error {
	var stay Stays
	if tx := repo.db.Where("id = ?", stayId).First(&stay); tx.Error != nil {
		return tx.Error
	}
	var stayDataMap = CoreStayToModel(stayData)
	if tx := repo.db.Model(&stay).Updates(stayDataMap); tx.Error != nil {
		return tx.Error
	}
	return nil
}

// Delete implements stays.StayDataInterface
func (repo *StayData) Delete(stayId string) error {
	if tx := repo.db.Where("id = ?", stayId).Delete(&Stays{}); tx.Error != nil {
		return tx.Error
	}
	return nil
}

func New(db *gorm.DB) stays.StayDataInterface {
	return &StayData{
		db: db,
	}
}
