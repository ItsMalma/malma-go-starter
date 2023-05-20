package main

import (
	"fmt"

	"github.com/ItsMalma/malma-go-starter/controller"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	config, err := GetConfig()
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(postgres.Open(config.Database))
	if err != nil {
		panic(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	sqlDB.SetMaxIdleConns(100)
	sqlDB.SetMaxOpenConns(100)

	errorHandler := controller.NewErrorHandler()

	app := fiber.New(fiber.Config{
		StrictRouting: true,
		CaseSensitive: true,
		ErrorHandler:  errorHandler,
	})
	app.Static("/", "./statics")

	if err := app.Listen(fmt.Sprintf("%v:%v", config.Server.Host, config.Server.Port)); err != nil {
		panic(err)
	}
}
