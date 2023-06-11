package users

import "alta/air-bnb/features/users/data"

type Core struct {
	ID 				string 		`json:"id" form:"id"`
	FullName	string		`json:"fullname" form:"fullname"`
	Email			string		`json:"email" form:"email"`
	Phone			string		`json:"phone" form:"phone"`
	Birth			string		`json:"birth" form:"birth"`
	Address		string		`json:"address" form:"address"`
}

type CoreGetAllResponse struct {
	ID 				string 		`json:"id" form:"id"`
	FullName	string		`json:"fullname" form:"fullname"`
	Birth			string		`json:"birth" form:"birth"`
	Address		string		`json:"address" form:"address"`
}

func ModelToCore(model data.Users) Core {
	return Core{
		ID: model.ID,
		FullName: model.FullName,
		Email: model.Email,
		Phone: model.Phone,
		Birth: model.Birth,
		Address: model.Address,
	}
}

func CoreToModel(user Core) data.Users {
	return data.Users{
		ID: user.ID,
		FullName: user.FullName,
		Email: user.Email,
		Phone: user.Phone,
		Birth: user.Birth,
		Address: user.Address,
	}
}

type UserDataInterface interface {
	Insert(userData Core) error
	Update(userId string, userData Core) error
	Select(userId string) (Core error)
	SelectAll() ([]CoreGetAllResponse, error)
	Delete(userId string) error
}

type UserServiceInterface interface {
	RegisterUser(userData Core) error
	EditUserById(userId string, userData Core) error
	GetUserById(userId string) (Core error)
	GetAllUsers() ([]CoreGetAllResponse, error)
	DeleteUserById(userId string) error
}