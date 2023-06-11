package main

import (
	"alta/air-bnb/app/config"
	"alta/air-bnb/app/database"
	"alta/air-bnb/app/migrations"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)


func main() {
	config := config.ReadEnv()
	database := database.InitDB(config)
	if errMigrate := migrations.InitMigrate(database); errMigrate != nil {
		log.Fatal(errMigrate.Error())
	}

	echo := echo.New()
	echo.Pre(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc3339}] ${status} ${method} ${host}${path} ${latency_human}` + "\n",
	}))
	echo.Pre(middleware.RemoveTrailingSlash())
	echo.Use(middleware.CORS())

	echo.Logger.Fatal(echo.Start(":8080"))
}