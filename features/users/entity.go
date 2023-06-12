package users

type Core struct {
	ID 				uint 		`json:"id" form:"id"`	
	FullName	string		`json:"fullname" form:"fullname"`
	Email			string		`json:"email" form:"email"`
	Phone			string		`json:"phone" form:"phone"`
	Birth			string		`json:"birth" form:"birth"`
	Address		string		`json:"address" form:"address"`
}

type CoreGetAllResponse struct {
	ID 				uint 		`json:"id" form:"id"`
	FullName	string		`json:"fullname" form:"fullname"`
	Birth			string		`json:"birth" form:"birth"`
	Address		string		`json:"address" form:"address"`
}

type CoreUserRequest struct {
	FullName	string		`json:"fullname" form:"fullname" validate:"required"`
	Email			string		`json:"email" form:"email" validate:"required,email"`
	Phone			string		`json:"phone" form:"phone"`
	Password	string 		`json:"password" form:"password" validate:"required,min=8"`
	Birth			string		`json:"birth" form:"birth"`
	Address		string		`json:"address" form:"address"`
}

type UserDataInterface interface {
	Insert(userData CoreUserRequest) (uint, error)
	Update(userId uint, userData CoreUserRequest) error
	Select(userId uint) (Core, error)
	SelectAll() ([]CoreGetAllResponse, error)
	Delete(userId uint) error
}

type UserServiceInterface interface {
	RegisterUser(userData CoreUserRequest) (uint, error)
	EditUserById(userId uint, userData CoreUserRequest) error
	GetUserById(userId uint) (Core, error)
	GetAllUsers() ([]CoreGetAllResponse, error)
	DeleteUserById(userId uint) error
}