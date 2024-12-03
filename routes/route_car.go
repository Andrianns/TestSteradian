package routes

import (
	"Stera/models"
	"Stera/service"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type CarRoute struct {
	service service.CarService
}

func NewCarRoute(app *fiber.App, service service.CarService) {
	route := &CarRoute{service: service}
	app.Get("/cars", route.GetAll)
	app.Get("/cars/:id", route.GetById)
	app.Post("/cars", route.Create)
	app.Put("/cars/1", route.Update)
	app.Delete("/cars/:id", route.Delete)
}

func (h *CarRoute) GetAll(c *fiber.Ctx) error {
	cars, err := h.service.GetAll()
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error:": err.Error()})
	}

	return c.JSON(cars)
}

func (h *CarRoute) GetById(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params(("id")))
	car, err := h.service.GetById(id)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error:": err.Error()})
	}
	return c.JSON(car)
}

func (h *CarRoute) Create(c *fiber.Ctx) error {
	var car models.Car
	if err := c.BodyParser(&car); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error:": err.Error()})
	}

	createCar, err := h.service.Create(car)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error:": err.Error()})
	}

	return c.Status(http.StatusAccepted).JSON(fiber.Map{
		"message": "success create car",
		"data":    createCar,
	})
}

func (h *CarRoute) Update(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params(("id")))
	var car models.Car
	if err := c.BodyParser(&car); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error:": err.Error()})
	}
	car.CarId = int(id)
	updatedCar, err := h.service.Update(car)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error:": err.Error()})
	}

	return c.Status(http.StatusAccepted).JSON(updatedCar)
}

func (h *CarRoute) Delete(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params(("id")))

	err := h.service.Delete(int(id))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error:": err.Error()})
	}

	return c.Status(http.StatusAccepted).JSON("success delete car")
}
