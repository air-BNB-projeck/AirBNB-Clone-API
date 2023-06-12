package data

import (
	"alta/air-bnb/features/users"

	"gorm.io/gorm"
)

type UserData struct {
	db *gorm.DB
}

// Insert implements users.UserDataInterface
func (repo *UserData) Insert(userData users.CoreUserRequest) (uint, error) {
	var user = CoreRequestToModel(userData)
	if tx := repo.db.Create(&user); tx.Error != nil {
		return 0, tx.Error
	}
	return user.ID, nil 
}

// Select implements users.UserDataInterface
func (repo *UserData) Select(userId uint) (users.Core, error) {
	var user Users
	if tx := repo.db.First(&user, userId); tx.Error != nil {
		return users.Core{}, tx.Error
	}

	var mapUser = ModelToCore(user)
	return mapUser, nil
}

// Update implements users.UserDataInterface
func (repo *UserData) Update(userId uint, userData users.CoreUserRequest) error {
	var user Users
	if tx := repo.db.First(&user, userId); tx.Error != nil {
		return tx.Error
	}
	var mapUser = CoreRequestToModel(userData)
	if tx := repo.db.Model(&user).Updates(mapUser); tx.Error != nil {
		return tx.Error
	}
	return nil
}

// Delete implements users.UserDataInterface
func (repo *UserData) Delete(userId uint) error {
	if tx := repo.db.Delete(&Users{}, userId); tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (repo *UserData) VerifyEmailUser(email string) (users.CoreLoginUserData, error) {
	var user Users
	if tx := repo.db.Where("email = ?", email).First(&user); tx.Error != nil {
		return users.CoreLoginUserData{}, tx.Error
	}
	var userMap = users.CoreLoginUserData{
		ID: user.ID,
		Password: user.Password,
	}
	return userMap, nil
}

func New(db *gorm.DB) users.UserDataInterface {
	return &UserData{
		db: db,
	}
}
