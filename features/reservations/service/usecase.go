package service

import (
	"alta/air-bnb/features/reservations"

	"github.com/go-playground/validator/v10"
)

type ReservationsService struct {
	reservationData reservations.ReservationsDataInterface
	validator       validator.Validate
}

// AddReservation implements reservations.ReservationServiceInterface
func (service *ReservationsService) AddReservation(reservationData reservations.CoreReservationRequest) (reservationId string, err error) {
	if errValidate := service.validator.Struct(reservationData); errValidate != nil {
		return "", errValidate
	}
	id, errInsert := service.reservationData.InsertReservation(reservationData)
	if errInsert != nil {
		return "", err
	}
	return id, nil
}

// CheckReservationAvailable implements reservations.ReservationServiceInterface
func (*ReservationsService) CheckReservationAvailable(reservationId string) (isAvailable bool) {
	panic("unimplemented")
}

func New(reservationData reservations.ReservationsDataInterface) reservations.ReservationServiceInterface {
	return &ReservationsService{
		reservationData: reservationData,
		validator:       *validator.New(),
	}
}
