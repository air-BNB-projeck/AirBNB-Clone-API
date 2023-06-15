package data

import (
	"alta/air-bnb/app/helper"
	"alta/air-bnb/features/reservations"
	_stayData "alta/air-bnb/features/stays/data"
	_userData "alta/air-bnb/features/users/data"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
	"gorm.io/gorm"
)

type ReservationData struct {
	db *gorm.DB
}

// InsertReservation implements reservations.ReservationsDataInterface
func (repo *ReservationData) InsertReservation(reservationData reservations.CoreReservationRequest) (reservationId string, err error) {
	reservationDataMap := CoreRequestToModel(reservationData)
	reserveId := helper.GenerateNewId()
	differenceDay, errParseDate := helper.SubstractDate(reservationData.StartDate, reservationData.EndDate)
	if errParseDate != nil {
		return "", errParseDate
	}
	var stay _stayData.Stays 
	if tx := repo.db.Where("id = ?", reservationDataMap.StayID).First(&stay); tx.Error != nil {
		return "", tx.Error
	}
	var user _userData.Users
	if tx := repo.db.Where("id = ?", reservationDataMap.UserID).First(&user); tx.Error != nil {
		return "", tx.Error
	}
	grossAmount := int64(stay.Price) * int64(differenceDay)
	items := []midtrans.ItemDetails{
		{
			ID: reserveId,
			Name: "Reservation: " + stay.Name,
			Price: int64(stay.Price),
			Qty: int32(differenceDay),
		},
	}
	initMidtrans := helper.InitMidtrans()
	transactionRequest := &coreapi.ChargeReq{
		PaymentType: coreapi.PaymentTypeBankTransfer,
		TransactionDetails: midtrans.TransactionDetails{
			OrderID: reserveId,
			GrossAmt: grossAmount,
		},
		BankTransfer: &coreapi.BankTransferDetails{
			Bank: midtrans.BankBri,
			VaNumber: "024",
		},
		CustomerDetails: &midtrans.CustomerDetails{
			FName: user.FullName,
			Email: user.Email,
			Phone: user.Phone,
		},
		Items: &items, 
	}
	charge, errCharge := initMidtrans.ChargeTransaction(transactionRequest)
	if errCharge != nil {
		return "", errCharge
	}
	reservationDataMap.TransactionID = charge.TransactionID
	reservationDataMap.OrderID = charge.OrderID
	reservationDataMap.TransactionStatus = charge.TransactionStatus
	reservationDataMap.PaymentType = charge.PaymentType
	reservationDataMap.GrossAmount = charge.GrossAmount;
	reservationDataMap.VANumbers = charge.VaNumbers[0].VANumber
	if tx := repo.db.Create(reservationDataMap); tx.Error != nil {
		return "", tx.Error
	}
	return reserveId, nil
}

// SelectReservationAvailable implements reservations.ReservationsDataInterface
func (*ReservationData) SelectReservationAvailable(reservationId string) (isAvailable bool) {
	panic("unimplemented")
}

func New(db *gorm.DB) reservations.ReservationsDataInterface {
	return &ReservationData{
		db: db,
	}
}
