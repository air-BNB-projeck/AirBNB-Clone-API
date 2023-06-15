package routers

import (
	"alta/air-bnb/app/middlewares"
	_userData "alta/air-bnb/features/users/data"
	_userHandler "alta/air-bnb/features/users/handler"
	_userService "alta/air-bnb/features/users/service"

	_stayData "alta/air-bnb/features/stays/data"
	_stayHandler "alta/air-bnb/features/stays/handler"
	_stayService "alta/air-bnb/features/stays/service"

	_reservationData "alta/air-bnb/features/reservations/data"
	_reservationHandler "alta/air-bnb/features/reservations/handler"
	_reservationService "alta/air-bnb/features/reservations/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRouters(db *gorm.DB, e *echo.Echo) {
	UserData := _userData.New(db)
	UserService := _userService.New(UserData)
	UserHandler := _userHandler.New(UserService)

	StayData := _stayData.New(db)
	StayService := _stayService.New(StayData)
	StayHandler := _stayHandler.New(StayService)

	ReservationData := _reservationData.New(db)
	ReservationService := _reservationService.New(ReservationData)
	ReservationHandler := _reservationHandler.New(ReservationService)

	e.POST("/users", UserHandler.PostUserHandler)
	e.GET("/users/profile", UserHandler.GetUserByIdHandler, middlewares.JWTMiddleware())
	e.PUT("/users/profile", UserHandler.UpdateUserByIdHandler, middlewares.JWTMiddleware())
	e.DELETE("/users/profile", UserHandler.DeleteUserByIdHandler, middlewares.JWTMiddleware())
	e.POST("/login", UserHandler.LoginUserHandler)

	e.POST("/stays", StayHandler.PostStayHandler, middlewares.JWTMiddleware())
	e.GET("/stays", StayHandler.GetAllStaysHandler)
	e.GET("/stays/:id", StayHandler.GetStayHandler)
	e.PUT("/stays/:id", StayHandler.PutStayHandler, middlewares.JWTMiddleware())
	e.DELETE("/stays/:id", StayHandler.DeleteStayHandler, middlewares.JWTMiddleware())
	e.POST("/stays/:id/images", StayHandler.PostStayImageHandler, middlewares.JWTMiddleware())

	e.POST("/reservations", ReservationHandler.PostReservationHandler, middlewares.JWTMiddleware())
}