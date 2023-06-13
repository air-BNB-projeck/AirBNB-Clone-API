package data

import (
	"alta/air-bnb/features/stays/data"
	"alta/air-bnb/features/users"

	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	FullName	string 			`gorm:"type:varchar(100);notNull"`
	Email			string 			`gorm:"type:varchar(50);unique:notNull"`
	Password 	string 			`gorm:"type:varchar(100);notNull"`
	Phone			string			`gorm:"type:varchar(50);unique"`
	Birth			string			`gorm:"type:varchar(50)"`
	Gender		string			`gorm:"type:text"`
	Stays 		[]data.Stays `gorm:"foreignKey:UserID"`	
}

func ModelToCore(model Users) users.Core {
	return users.Core{
		ID: model.ID,
		FullName: model.FullName,
		Email: model.Email,
		Phone: model.Phone,
		Birth: model.Birth,
		Gender: model.Gender,
	}
}

func CoreToModel(user users.Core) Users {
	return Users{
		FullName: user.FullName,
		Email: user.Email,
		Phone: user.Phone,
		Birth: user.Birth,
		Gender: user.Gender,
	}
}

func ModelToCoreGetAll(model Users) users.CoreGetAllResponse {
	return users.CoreGetAllResponse{
		ID: model.ID,
		FullName: model.FullName,
		Birth: model.Birth,
		Gender: model.Gender,
	}
}

func CoreRequestToModel(request users.CoreUserRequest) Users {
	return Users{
		FullName: request.FullName,
		Email: request.Email,
		Password: request.Password,
		Phone: request.Phone,
		Birth: request.Birth,
		Gender: request.Gender, 
	}
}