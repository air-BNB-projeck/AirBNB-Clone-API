package data

import (
	"alta/air-bnb/app/helper"
	"alta/air-bnb/features/reservations"
	"time"
)

type Reservations struct {
	ID						string 				`gorm:"type:varchar(100);primaryKey"`
	UserID				uint					`gorm:"type:uint;notNull"`
	StayID				string				`gorm:"type:varchar(100);notNull"`
	StartDate			time.Time			
	EndDate				time.Time			
	TransactionID	string				`gorm:"type:varchar(100)"`
	OrderID				string				`gorm:"type:varchar(100)"`
	TransactionStatus		string				`gorm:"type:varchar(50)"`
	PaymentType		string				`gorm:"type:varchar(50)"`
	GrossAmount		string				`gorm:"type:decimal(10,2)"`
	VANumbers	string				`gorm:"type:varchar(50)"`
	User					Users					`gorm:"foreignKey:UserID"`
	Stay					Stays					`gorm:"foreignKey:StayID"`
}

type Users struct {
	ID				string			`gorm:"type:uint;primaryKey"`
	FullName	string 			`gorm:"type:varchar(100);notNull"`
	Email			string 			`gorm:"type:varchar(50);unique:notNull"`
}

type Stays struct {
	ID						string			`gorm:"type:varchar(50);primaryKey"`
	Name					string			`gorm:"type:varchar(100)"`
	Description		string			`gorm:"type:text"`
	Latitude			float64			`gorm:"type:double"`
	Longitude			float64			`gorm:"type:double"`
	City					string			`gorm:"type:varchar(50)"`
	Rating				float64			`gorm:"type:double"`
}

func CoreRequestToModel(reservationData reservations.CoreReservationRequest) Reservations {
	startDate, _ := helper.ParseDate(reservationData.StartDate)
	endDate, _ := helper.ParseDate(reservationData.EndDate)
	return Reservations{
		UserID: reservationData.UserID,
		StayID: reservationData.StayID,
		StartDate: startDate,
		EndDate: endDate,
	}
}

func ModelToCoreReservation(reservation Reservations) reservations.CoreReservation {
	return reservations.CoreReservation{
		ID: reservation.ID,
		StartDate: reservation.StartDate,
		EndDate: reservation.EndDate,
	}
}