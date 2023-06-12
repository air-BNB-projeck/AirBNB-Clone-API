package handler

import (
	"alta/air-bnb/app/helper"
	"alta/air-bnb/features/users"
	"strconv"
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

func (handler *UserHandler) GetAllUsersHandler(c echo.Context) error {
	users, err := handler.service.GetAllUsers()
	if err != nil {
		return helper.StatusInternalServerError(c, err.Error())
	}
	return helper.StatusOKWithData(c, "Berhasil mendapatkan data pengguna", map[string]any{
		"users": users,
	})
}

func (handler *UserHandler) GetUserByIdHandler(c echo.Context) error {
	userId, errParam := strconv.Atoi(c.Param("id"))
	if userId == 0 && errParam != nil {
		return helper.StatusBadRequestResponse(c, "invalid param request: " + errParam.Error())
	} else {
		user, err := handler.service.GetUserById(uint(userId))
		if err != nil {
			return helper.StatusInternalServerError(c, err.Error())
		}
		return helper.StatusOKWithData(c, "Berhasil mendapatkan data pengguna", map[string]any{
			"user": user,
		})
	}
}

func (handler *UserHandler) UpdateUserByIdHandler(c echo.Context) error {
	userId, errParam := strconv.Atoi(c.Param("id"))
	if userId == 0 && errParam != nil {
		return helper.StatusBadRequestResponse(c, "invalid param request: " + errParam.Error())
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
	userId, errParam := strconv.Atoi(c.Param("id"))
	if userId == 0 && errParam != nil {
		return helper.StatusBadRequestResponse(c, "invalid param request: " + errParam.Error())
	} else {
		if err := handler.service.DeleteUserById(uint(userId)); err != nil {
			return helper.StatusInternalServerError(c, err.Error())
		}
		return helper.StatusOK(c, "Berhasil menghapus data pengguna")
	}
}

func New(service users.UserServiceInterface) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

