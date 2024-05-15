package main

import (
	"github.com/JimmyTanapon/Go-Clean-Architecture/adepters"
	"github.com/JimmyTanapon/Go-Clean-Architecture/entities"
	"github.com/JimmyTanapon/Go-Clean-Architecture/usecase"

	"github.com/glebarez/sqlite"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func main() {
	app := fiber.New()
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		panic("fail to connect database!")
	}

	orderRepo := adepters.NewGormOrderRepository(db)
	orderService := usecase.NewOrderService(orderRepo)
	orderHandler := adepters.NewHttpOrderHander(orderService)
	app.Post("/order", orderHandler.CreatOrder)
	db.AutoMigrate(&entities.Order{})
	app.Listen(":8000")
}
