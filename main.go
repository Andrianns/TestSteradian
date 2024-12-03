package main

import (
	"Stera/config"
	"Stera/repository"
	"Stera/routes"
	"Stera/service"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	db, err := config.ConnectDB()
	if err != nil {
		log.Fatalf("error connect db: %v", err)
	}

	carRepository := repository.NewCarRepository(db)
	carService := service.NewCarService(carRepository)
	routes.NewCarRoute(app, carService)

	orderRepository := repository.NewOrderRepository(db)
	orderService := service.NewOrderService(orderRepository)
	routes.NewOrderRoute(app, orderService)

	if err := app.Listen(":3000"); err != nil {
		log.Fatalf("Failed start server: %v", err)
	}
}
