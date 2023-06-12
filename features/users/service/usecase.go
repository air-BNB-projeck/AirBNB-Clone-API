package service

import (
	"alta/air-bnb/app/helper"
	"alta/air-bnb/features/users"
	"errors"

	"github.com/go-playground/validator/v10"
)

type UserService struct {
	data users.UserDataInterface
	validate *validator.Validate
}

// RegisterUser implements users.UserServiceInterface
func (service *UserService) RegisterUser(userData users.CoreUserRequest) (uint, error) {
	if errValidate := service.validate.Struct(userData); errValidate != nil {
		return 0, errors.New("error validation: " + errValidate.Error())
	}
	hashedPassword, err := helper.HashPassword(userData.Password)
	if err != nil {
		return 0, errors.New("error hash password: " + err.Error())
	}
	userData.Password = string(hashedPassword)
	userId, err := service.data.Insert(userData) 
	if err != nil {
		return 0, err
	}
	return userId, nil
}

// EditUserById implements users.UserServiceInterface
func (service *UserService) EditUserById(userId uint, userData users.CoreUserRequest) error {
	if errValidate := service.validate.Struct(userData); errValidate != nil {
		return errors.New("error validation: " + errValidate.Error())
	}
	hashedPassword, err := helper.HashPassword(userData.Password)
	if err != nil {
		return errors.New("error hash password: " + err.Error())
	}
	userData.Password = string(hashedPassword)
	if err := service.data.Update(userId, userData); err != nil {
		return err
	}
	return nil
}

// GetAllUsers implements users.UserServiceInterface
func (service *UserService) GetAllUsers() ([]users.CoreGetAllResponse, error) {
	users, err := service.data.SelectAll();
	if err != nil {
		return nil, err
	}
	return users, nil 
}

// GetUserById implements users.UserServiceInterface
func (service *UserService) GetUserById(userId uint) (users.Core, error) {
	user, err := service.data.Select(userId);
	if err != nil {
		return users.Core{}, err
	}
	return user, nil
}

// DeleteUserById implements users.UserServiceInterface
func (service *UserService) DeleteUserById(userId uint) error {
	if err := service.data.Delete(userId); err != nil {
		return err
	}
	return nil
}

func New(data users.UserDataInterface) users.UserServiceInterface {
	return &UserService{
		data: data,
		validate: validator.New(),
	}
}
