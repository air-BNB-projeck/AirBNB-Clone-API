package routers

import (
	"alta/air-bnb/app/middlewares"
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
	e.GET("/users/profile", UserHandler.GetUserByIdHandler, middlewares.JWTMiddleware())
	e.PUT("/users/profile", UserHandler.UpdateUserByIdHandler, middlewares.JWTMiddleware())
	e.DELETE("/users/profile", UserHandler.DeleteUserByIdHandler, middlewares.JWTMiddleware())
	e.POST("/login", UserHandler.LoginUserHandler)
}