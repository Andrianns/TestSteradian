package routes

import (
	"Stera/service"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type OrderRoute struct {
	service service.OrderService
}

func NewOrderRoute(app *fiber.App, service service.OrderService) {
	route := &OrderRoute{service: service}
	app.Get("/cars", route.GetAll)
}

func (h *OrderRoute) GetAll(c *fiber.Ctx) error {
	orders, err := h.service.GetAll()
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error:": err.Error()})
	}

	return c.JSON(orders)
}
