package routers

import (
	_userData "alta/air-bnb/features/users/data"
	_userHandler "alta/air-bnb/features/users/handler"
	_userService "alta/air-bnb/features/users/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRouters(db *gorm.DB, e *echo.Echo) {
	UserData := _userData.New(db)
	UserService := _userService.New(UserData)
	UserHandler := _userHandler.New(UserService)

	e.POST("/users", UserHandler.PostUserHandler)
	e.GET("/users", UserHandler.GetAllUsersHandler)
	e.GET("/users/:id", UserHandler.GetUserByIdHandler)
	e.PUT("/users/:id", UserHandler.UpdateUserByIdHandler)
	e.DELETE("/users/:id", UserHandler.DeleteUserByIdHandler)
}