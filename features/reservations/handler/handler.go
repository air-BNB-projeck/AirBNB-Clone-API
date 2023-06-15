package handler

import (
	"alta/air-bnb/app/helper"
	"alta/air-bnb/app/middlewares"
	"alta/air-bnb/features/reservations"

	"github.com/labstack/echo/v4"
)

type ReservationHandler struct {
	service reservations.ReservationServiceInterface
}

func (handler *ReservationHandler) PostReservationHandler(c echo.Context) error {
	var payload reservations.CoreReservationRequest
	if errBind := c.Bind(&payload); errBind != nil {
		return helper.StatusBadRequestResponse(c, "echo error bind: " + errBind.Error())
	}
	userId := middlewares.ExtractTokenUserId(c)
	payload.UserID = userId
	reservationId, err := handler.service.AddReservation(payload)
	if err != nil {
		return helper.StatusInternalServerError(c, err.Error())
	}
	return helper.StatusCreated(c, "Berhasil menambahkan transaksi reservasi", map[string]any{
		"reservationId": reservationId,
	})
}

func New(service reservations.ReservationServiceInterface) *ReservationHandler {
	return &ReservationHandler{
		service: service,
	}
}