package handler

import (
	"alta/air-bnb/app/helper"
	"alta/air-bnb/app/middlewares"
	"alta/air-bnb/features/users"
	"strings"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	service 	users.UserServiceInterface	
}

func (handler *UserHandler) PostUserHandler(c echo.Context) error {
	var payload users.CoreUserRequest
	if errBind := c.Bind(&payload); errBind != nil {
		return helper.StatusBadRequestResponse(c, "error bind payload: " + errBind.Error())
	}
	userId, errRegister := handler.service.RegisterUser(payload) 
	if errRegister != nil {
		if strings.Contains(errRegister.Error(), "validation") {
			return helper.StatusBadRequestResponse(c, errRegister.Error())
		} else {
			return helper.StatusInternalServerError(c, errRegister.Error())
		}
	}
	return helper.StatusCreated(c, "Berhasil menambahkan pengguna", map[string]any{
		"userId": userId,
	})
}

func (handler *UserHandler) GetUserByIdHandler(c echo.Context) error {
	userId := middlewares.ExtractTokenUserId(c)
	if userId == 0 {
		return helper.StatusAuthorizationErrorResponse(c, "Anda harus login untuk mengakses resource ini")
	} else {
		user, err := handler.service.GetUserById(userId)
		if err != nil {
			return helper.StatusInternalServerError(c, err.Error())
		}
		return helper.StatusOKWithData(c, "Berhasil mendapatkan data pengguna", map[string]any{
			"user": user,
		})
	}
}

func (handler *UserHandler) UpdateUserByIdHandler(c echo.Context) error {
	userId := middlewares.ExtractTokenUserId(c)
	if userId == 0 {
		return helper.StatusAuthorizationErrorResponse(c, "Anda harus login untuk merubah resource ini")
	} else {
		var payload users.CoreUserRequest
		if errBind := c.Bind(&payload); errBind != nil {
			return helper.StatusBadRequestResponse(c, "error bind payload: " + errBind.Error())
		}
		if err := handler.service.EditUserById(uint(userId), payload); err != nil {
			if strings.Contains(err.Error(), "validation") {
				return helper.StatusBadRequestResponse(c, err.Error())
			} else {
				return helper.StatusInternalServerError(c, err.Error())
			}
		}
		return helper.StatusOK(c, "Berhasil memperbarui data pengguna")
	}
}

func (handler *UserHandler) DeleteUserByIdHandler(c echo.Context) error {
	userId := middlewares.ExtractTokenUserId(c)
	if userId == 0 {
		return helper.StatusAuthorizationErrorResponse(c, "Anda harus login untuk menghapus resource ini")
	} else {
		if err := handler.service.DeleteUserById(uint(userId)); err != nil {
			return helper.StatusInternalServerError(c, err.Error())
		}
		return helper.StatusOK(c, "Berhasil menghapus data pengguna")
	}
}

func (handler *UserHandler) LoginUserHandler(c echo.Context) error {
	var payload users.CoreLoginUserRequest
	if errBind := c.Bind(&payload); errBind != nil {
		return helper.StatusBadRequestResponse(c, "error bind payload: " + errBind.Error())
	}
	userId, errLogin := handler.service.LoginUser(payload)
	if errLogin != nil {
		return helper.StatusBadRequestResponse(c, errLogin.Error())
	}
	accessToken, errCreateAccessToken := middlewares.CreateAccessToken(userId)
	if errCreateAccessToken != nil {
		return helper.StatusInternalServerError(c, "error create access token: " + errCreateAccessToken.Error())
	}
	return helper.StatusOKWithData(c, "Login berhasil", map[string]any{
		"accessToken": accessToken,
	})
}

func New(service users.UserServiceInterface) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

