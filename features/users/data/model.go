package data

import "alta/air-bnb/features/users"

type Users struct {
	ID 				string  `gorm:"type:varchar(50);primaryKey"`
	FullName	string 	`gorm:"type:varchar(100);notNull"`
	Email			string 	`gorm:"type:varchar(50);unique:notNull"`
	Password 	string 	`gorm:"type:varchar(50);notNull"`
	Phone			string	`gorm:"type:varchar(50);unique"`
	Birth			string	`gorm:"type:date"`
	Address		string	`gorm:"type:text"`
}

func ModelToCore(model Users) users.Core {
	return users.Core{
		ID: model.ID,
		FullName: model.FullName,
		Email: model.Email,
		Phone: model.Phone,
		Birth: model.Birth,
		Address: model.Address,
	}
}

func CoreToModel(user users.Core) Users {
	return Users{
		ID: user.ID,
		FullName: user.FullName,
		Email: user.Email,
		Phone: user.Phone,
		Birth: user.Birth,
		Address: user.Address,
	}
}

func ModelToCoreGetAll(model Users) users.CoreGetAllResponse {
	return users.CoreGetAllResponse{
		ID: model.ID,
		FullName: model.FullName,
		Birth: model.Birth,
		Address: model.Address,
	}
}