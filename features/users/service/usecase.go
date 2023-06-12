package service

import "alta/air-bnb/features/users"

type UserService struct {
	data users.UserDataInterface
}

// RegisterUser implements users.UserServiceInterface
func (*UserService) RegisterUser(userData users.Core) error {
	panic("unimplemented")
}

// DeleteUserById implements users.UserServiceInterface
func (*UserService) DeleteUserById(userId string) error {
	panic("unimplemented")
}

// EditUserById implements users.UserServiceInterface
func (*UserService) EditUserById(userId string, userData users.Core) error {
	panic("unimplemented")
}

// GetAllUsers implements users.UserServiceInterface
func (*UserService) GetAllUsers() ([]users.CoreGetAllResponse, error) {
	panic("unimplemented")
}

// GetUserById implements users.UserServiceInterface
func (*UserService) GetUserById(userId string) (users.Core, error) {
	panic("unimplemented")
}

func New(data users.UserDataInterface) users.UserServiceInterface {
	return &UserService{
		data: data,
	}
}
