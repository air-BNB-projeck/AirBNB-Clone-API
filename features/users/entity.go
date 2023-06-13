package users

type Core struct {
	ID 				uint 		`json:"id" form:"id"`	
	FullName	string		`json:"fullname" form:"fullname"`
	Email			string		`json:"email" form:"email"`
	Phone			string		`json:"phone" form:"phone"`
	Birth			string		`json:"birth" form:"birth"`
	Gender		string		`json:"gender" form:"gender"`
}

type CoreGetAllResponse struct {
	ID 				uint 		`json:"id" form:"id"`
	FullName	string		`json:"fullname" form:"fullname"`
	Birth			string		`json:"birth" form:"birth"`
	Gender		string		`json:"gender" form:"gender"`
}

type CoreUserRequest struct {
	FullName	string		`json:"fullname" form:"fullname" validate:"required"`
	Email			string		`json:"email" form:"email" validate:"required,email"`
	Phone			string		`json:"phone" form:"phone"`
	Password	string 		`json:"password" form:"password" validate:"required,min=8"`
	Birth			string		`json:"birth" form:"birth"`
	Gender		string		`json:"gender" form:"gender"`
}

type CoreLoginUserRequest struct {
	Email 		string 			`json:"email" form:"email" validate:"required,email"`
	Password	string 			`json:"password" form:"email" validate:"required,min=8"`
}

type CoreLoginUserData struct {
	ID			 uint 
	Password string
}

type UserDataInterface interface {
	Insert(userData CoreUserRequest) (uint, error)
	Update(userId uint, userData CoreUserRequest) error
	Select(userId uint) (Core, error)
	Delete(userId uint) error
	VerifyEmailUser(email string) (CoreLoginUserData, error)
}

type UserServiceInterface interface {
	RegisterUser(userData CoreUserRequest) (uint, error)
	EditUserById(userId uint, userData CoreUserRequest) error
	GetUserById(userId uint) (Core, error)
	DeleteUserById(userId uint) error
	LoginUser(loginPayload CoreLoginUserRequest) (uint, error)
}