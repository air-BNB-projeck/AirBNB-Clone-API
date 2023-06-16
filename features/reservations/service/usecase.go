package service

import (
	"alta/air-bnb/features/reservations"
	"errors"

	"github.com/go-playground/validator/v10"
)

type ReservationsService struct {
	reservationData reservations.ReservationsDataInterface
	validator       validator.Validate
}

// CheckReservationAvailable implements reservations.ReservationServiceInterface
func (service *ReservationsService) CheckReservationAvailable(reservationData reservations.CoreReservationCheckRequest) (isAvailable bool, err error) {
	if errValidate := service.validator.Struct(reservationData); errValidate != nil {
		return false, errors.New("error validation: " + errValidate.Error())
	}
	isStayAvailable, errGetReservation := service.reservationData.SelectReservationAvailable(reservationData)
	if errGetReservation != nil {
		return false, errors.New("error get reservation data: " + errGetReservation.Error())
	}
	return isStayAvailable, nil
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

func New(reservationData reservations.ReservationsDataInterface) reservations.ReservationServiceInterface {
	return &ReservationsService{
		reservationData: reservationData,
		validator:       *validator.New(),
	}
}
