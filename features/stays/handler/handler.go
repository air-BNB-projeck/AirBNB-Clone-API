package handler

import (
	"alta/air-bnb/app/helper"
	"alta/air-bnb/app/middlewares"
	"alta/air-bnb/features/stays"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type StayHandler struct {
	service stays.StayServiceInterface
}

func (handler *StayHandler) PostStayReviewHandler(c echo.Context) error {
	stayId := c.Param("id")
	var payload stays.CoreStayReviewRequest
	if errBind := c.Bind(&payload); errBind != nil {
		return helper.StatusBadRequestResponse(c, "error bind payload: " + errBind.Error())
	}
	userId := middlewares.ExtractTokenUserId(c)
	payload.UserID = userId
	payload.StayID = stayId
	if err := handler.service.AddStayReview(payload); err != nil {
		if strings.Contains(err.Error(), "validation") {
			return helper.StatusBadRequestResponse(c, err.Error())
		} else {
			return helper.StatusInternalServerError(c, err.Error())
		}
	}
	return helper.StatusCreated(c, "Berhasil menambahkan stay review", map[string]any{
		"stayId": stayId,
	})
}

func (handler *StayHandler) DeleteStayImageHandler(c echo.Context) error {
	stayId := c.Param("stayId")
	imageId, _ := strconv.Atoi(c.Param("imageId"))
	if err := handler.service.DeleteStayImage(stayId, uint(imageId)); err != nil {
		return helper.StatusInternalServerError(c, err.Error())
	}
	return helper.StatusOK(c, "Berhasil menghapus stay image")
}

func (handler *StayHandler) PostStayImageHandler(c echo.Context) error {
	stayId := c.Param("id")
	var payload stays.CoreStayImageRequest
	if errBind := c.Bind(&payload); errBind != nil {
		return helper.StatusBadRequestResponse(c, "error bind payload: " + errBind.Error())
	}
	file, err := c.FormFile("image");
	if err != nil {
		return helper.StatusBadRequestResponse(c, "error get file image: " + err.Error())
	}
	payload.Image = file
	errAddImage := handler.service.AddStayImage(stayId, payload); 
	if errAddImage != nil {
		if strings.Contains(errAddImage.Error(), "validation") {
			return helper.StatusBadRequestResponse(c, errAddImage.Error())
		} else {
			return helper.StatusInternalServerError(c, errAddImage.Error())
		}
	}
	return helper.StatusCreated(c, "Berhasil menambahkan stay image", map[string]any{
		"stayId": stayId,
	})
}

func (handler *StayHandler) PostStayHandler(c echo.Context) error {
	var payload stays.CoreStayRequest
	if errBind := c.Bind(&payload); errBind != nil {
		return helper.StatusBadRequestResponse(c, "error bind payload: " + errBind.Error())
	}
	file, err := c.FormFile("image");
	if err != nil {
		return helper.StatusBadRequestResponse(c, "error get file image: " + err.Error())
	}
	userId := middlewares.ExtractTokenUserId(c)
	payload.Image = file 
	payload.UserID = userId
	stayId, err := handler.service.AddStay(payload);
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return helper.StatusBadRequestResponse(c, err.Error())
		} else {
			return helper.StatusInternalServerError(c, err.Error())
		}
	}
	return helper.StatusCreated(c, "Berhasil menambahkan stay", map[string]any{
		"stayId": stayId,
	})
}

func (handler *StayHandler) GetAllStaysHandler(c echo.Context) error {
	allStays, err := handler.service.GetAllStays()
	if err != nil {
		return helper.StatusInternalServerError(c, err.Error())
	}
	return helper.StatusOKWithData(c, "Berhasil mendapatkan semua data stays", map[string]any{
		"stays": allStays,
	})
}

func (handler *StayHandler) GetStayHandler(c echo.Context) error {
	stayId := c.Param("id")
	stay, err := handler.service.GetStay(stayId)
	if err != nil {
		return helper.StatusInternalServerError(c, err.Error())
	}
	return helper.StatusOKWithData(c, "Berhasil mendapatkan data stay", map[string]any{
		"stay": stay,
	})
}

func (handler *StayHandler) PutStayHandler(c echo.Context) error {
	stayId := c.Param("id")
	var payload stays.CoreStayRequest
	if errBind := c.Bind(&payload); errBind != nil {
		return helper.StatusBadRequestResponse(c, "error bind payload: " + errBind.Error())
	}
	userId := middlewares.ExtractTokenUserId(c)
	payload.UserID = userId
	if err := handler.service.EditStay(stayId, payload); err != nil {
		return helper.StatusInternalServerError(c, err.Error())
	}
	return helper.StatusOK(c, "Berhasil memperbarui data stay")
}

func (handler *StayHandler) DeleteStayHandler(c echo.Context) error {
	stayId := c.Param("id")
	if err := handler.service.DeleteStay(stayId); err != nil {
		return helper.StatusInternalServerError(c, err.Error())
	}
	return helper.StatusOK(c, "Berhasil menghapus data stay")
}

func New(service stays.StayServiceInterface) *StayHandler {
	return &StayHandler{
		service: service,
	}
}