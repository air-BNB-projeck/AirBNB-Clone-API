package reservations

type CoreReservation struct {
	ID						string		`json:"id" form:"id"`
	User					Users			`json:"users" form:"users"`
	Stay					Stays			`json:"stays" form:"stays"`
	StartDate			string		`json:"startDate" form:"startDate"`
	EndDate				string		`json:"endDate" form:"endDate"`
	TransactionID	string		`json:"transactionId" form:"transactionId"`
	Status				string		`json:"status" form:"status"`
	PaymentType		string		`json:"paymentType" form:"paymentType"`
	GrossAmount		float64		`json:"grossAmount" form:"grossAmount"`
}

type Users struct {
	ID						string			`json:"id" form:"id"`
	FullName			string			`json:"fullname" form:"fullname"`
	Email					string			`json:"email" form:"email"`
}

type Stays struct {
	ID						string			`json:"id" form:"id"`
	Name					string			`json:"name" form:"name"`
	Description		string			`json:"description" form:"description"`
	Latitude			float64			`json:"latitude" form:"latitude"`
	Longitude			float64			`json:"longitude" form:"longitude"`
	City					string			`json:"city" form:"city"`
	Rating				float64			`json:"rating" form:"string"`
}

type CoreReservationRequest struct {
	UserID				uint 			`json:"userId" form:"userId"`
	StayID				string			`json:"stayId" form:"stayId" validate:"required"`
	StartDate			string			`json:"startDate" form:"startDate" validate:"required"`
	EndDate				string 			`json:"endDate" form:"endDate" validate:"required"`
}

type ReservationsDataInterface interface {
	InsertReservation(reservationData CoreReservationRequest) (reservationId string, err error)
	SelectReservationAvailable(reservationId string) (isAvailable bool)
}

type ReservationServiceInterface interface {
	AddReservation(reservationData CoreReservationRequest) (reservationId string, err error)
	CheckReservationAvailable(reservationId string) (isAvailable bool)
}