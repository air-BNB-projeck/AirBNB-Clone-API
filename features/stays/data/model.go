package data

import (
	"alta/air-bnb/features/stays"
	"time"

	"gorm.io/gorm"
)

type Stays struct {
	ID 						string 			`gorm:"type:varchar(50);primaryKey"`
	CreatedAt 		time.Time
	UpdatedAt 		time.Time	
	DeletedAt 		gorm.DeletedAt 	`gorm:"index"`
	Name 					string			`gorm:"type:varchar(100);notNull"`
	Price					float64			`gorm:"type:decimal(10,2);notNull"`
	Description		string			`gorm:"type:text"`
	Latitude			float64			`gorm:"type:double"`
	Longitude			float64			`gorm:"type:double"`
	City					string			`gorm:"type:varchar(50)"`
	Bedrooms			int					`gorm:"type:int"`
	Bathrooms			int					`gorm:"type:int"`
	TV						int					`gorm:"type:int"`
	Wifi					int					`gorm:"type:int"`
	Pool					int					`gorm:"type:int"`
	Rating				float64			`gorm:"type:double"`
	UserID				uint				`gorm:"type:uint"`
	User					Users				`gorm:"foreignKey:UserID"`
	StaysImages		[]StayImages `gorm:"foreignKey:StayID"`
}

type StayImages struct {
	gorm.Model
	ImageUrl			string 		`gorm:"type:varchar(150)"`
	StayID				string		`gorm:"type:varchar(50)"`
}

type Users struct {
	ID				string			`gorm:"type:uint;primaryKey"`
	FullName	string 			`gorm:"type:varchar(100);notNull"`
	Email			string 			`gorm:"type:varchar(50);unique:notNull"`
}

func CoreStayToModel(stayCore stays.CoreStayRequest) Stays {
	return Stays{
		ID: "",
		Name: stayCore.Name,
		Price: stayCore.Price,
		Description: stayCore.Description,
		Latitude: stayCore.Latitude,
		Longitude: stayCore.Longitude,
		City: stayCore.City,
		Bedrooms: stayCore.Bedrooms,
		Bathrooms: stayCore.Bathrooms,
		TV: stayCore.TV,
		Wifi: stayCore.Wifi,
		Pool: stayCore.Pool,
		Rating: stayCore.Rating,
		UserID: stayCore.UserID,
	}
}

func CoreRequestToModel(stayRequest stays.CoreStayRequest) Stays {
	return Stays{
		ID: "",
		Name: stayRequest.Name,
		Price: stayRequest.Price,
		Description: stayRequest.Description,
		Latitude: stayRequest.Latitude,
		Longitude: stayRequest.Longitude,
		City: stayRequest.City,
		Bedrooms: stayRequest.Bedrooms,
		Bathrooms: stayRequest.Bathrooms,
		TV: stayRequest.TV,
		Wifi: stayRequest.Wifi,
		Pool: stayRequest.Pool,
		Rating: stayRequest.Rating,
		UserID: stayRequest.UserID,
	}
}

func ModelStayToCore(stayModel Stays) stays.Core {
	return stays.Core{
		ID: stayModel.ID,
		Name: stayModel.Name,
		Price: stayModel.Price,
		Description: stayModel.Description,
		Latitude: stayModel.Latitude,
		Longitude: stayModel.Longitude,
		City: stayModel.City,
		Bedrooms: stayModel.Bedrooms,
		Bathrooms: stayModel.Bathrooms,
		TV: stayModel.TV,
		Wifi: stayModel.Wifi,
		Pool: stayModel.Pool,
		Rating: stayModel.Rating,
		User: stays.Users{},
		StayImages: []string{},
	}
}

