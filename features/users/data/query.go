package data

import (
	"alta/air-bnb/features/users"

	"gorm.io/gorm"
)

type UserData struct {
	db *gorm.DB
}

// Insert implements users.UserDataInterface
func (repo *UserData) Insert(userData users.Core) error {
	var user = CoreToModel(userData)
	if tx := repo.db.Create(&user); tx.Error != nil {
		return tx.Error
	}
	return nil 
}

// Select implements users.UserDataInterface
func (repo *UserData) Select(userId string) (users.Core, error) {
	var user Users
	if tx := repo.db.First(&user, userId); tx.Error != nil {
		return users.Core{}, tx.Error
	}

	var mapUser = ModelToCore(user)
	return mapUser, nil
}

// SelectAll implements users.UserDataInterface
func (repo *UserData) SelectAll() ([]users.CoreGetAllResponse, error) {
	var usersData []Users
	if tx := repo.db.Find(&usersData); tx.Error != nil {
		return nil, tx.Error
	}

	var mapUsers []users.CoreGetAllResponse 
	for _, user := range usersData {
		var userData = ModelToCoreGetAll(user)
		mapUsers = append(mapUsers, userData)
	}

	return mapUsers, nil
}

// Update implements users.UserDataInterface
func (repo *UserData) Update(userId string, userData users.Core) error {
	var user Users
	if tx := repo.db.First(&user, userId); tx.Error != nil {
		return tx.Error
	}
	var mapUser = CoreToModel(userData)
	if tx := repo.db.Model(&user).Updates(mapUser); tx.Error != nil {
		return tx.Error
	}
	return nil
}

// Delete implements users.UserDataInterface
func (repo *UserData) Delete(userId string) error {
	if tx := repo.db.Delete(&Users{}, userId); tx.Error != nil {
		return tx.Error
	}
	return nil
}

func New(db *gorm.DB) users.UserDataInterface {
	return &UserData{
		db: db,
	}
}
