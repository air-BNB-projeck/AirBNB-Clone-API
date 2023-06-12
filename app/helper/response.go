package helper

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func StatusOK(c echo.Context, message string) error {
	return c.JSON(http.StatusOK, map[string]any{
		"status": "success",
		"message": message,
	})
}


func StatusOKWithData(c echo.Context, message string, data any) error {
	return c.JSON(http.StatusOK, map[string]any{
		"status": "success",
		"message": message,
		"data": data,
	})
}

func StatusCreated(c echo.Context, message string, data any) error {
	return c.JSON(http.StatusCreated, map[string]any{
	"status": "success",
		"message": message,
		"data": data,
	})
}

func StatusBadRequestResponse(c echo.Context, message string) error {
	return c.JSON(http.StatusBadRequest, map[string]any{
		"status": "fail",
		"message": message,
	})
}

func StatusNotFoundResponse(c echo.Context, message string) error {
	return c.JSON(http.StatusNotFound, map[string]any{
		"status": "fail",
		"message": message,
	})
}

func StatusAuthorizationErrorResponse(c echo.Context, message string) error {
	return c.JSON(http.StatusUnauthorized, map[string]any{
		"status": "fail",
		"message": message,
	})
}

func StatusForbiddenResponse(c echo.Context, message string) error {
	return c.JSON(http.StatusForbidden, map[string]any{
		"status": "fail",
		"message": message,
	})
}

func StatusInternalServerError(c echo.Context ,message string) error {
	return c.JSON(http.StatusInternalServerError, map[string]any{
		"status": "fail",
		"message": "Terjadi kesalahan di server kami: " + message,
	})
}

