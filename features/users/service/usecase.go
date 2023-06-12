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
	if passwordValidate := helper.ValidatePassword(userData.Password); !passwordValidate {
		return 0, errors.New("error input password: Password harus alphanumeric dan paling tidak berisi 1 simbol dan 1 huruf besar")
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

func (service *UserService) LoginUser(loginPayload users.CoreLoginUserRequest) (uint, error) {
	if errValidate := service.validate.Struct(loginPayload); errValidate != nil {
		return 0, errors.New("error validation: " + errValidate.Error())
	}
	user, errVerifyEmail := service.data.VerifyEmailUser(loginPayload.Email)
	if errVerifyEmail != nil {
		return 0, errors.New("email tidak terdaftar")
	}
	if match := helper.CompareHashedPassword(user.Password, loginPayload.Password); !match {
		return 0, errors.New("kredensial tidak valid")
	} 
	return user.ID, nil
}

func New(data users.UserDataInterface) users.UserServiceInterface {
	return &UserService{
		data: data,
		validate: validator.New(),
	}
}
