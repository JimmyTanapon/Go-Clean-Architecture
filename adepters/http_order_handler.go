package adepters

import (
	"github.com/JimmyTanapon/Go-Clean-Architecture/entities"
	"github.com/JimmyTanapon/Go-Clean-Architecture/usecase"
	"github.com/gofiber/fiber/v2"
)

type HttpOrderHandler struct {
	Orderusecase usecase.OrderUseCase
}

func NewHttpOrderHander(useCase usecase.OrderUseCase) *HttpOrderHandler {
	return &HttpOrderHandler{Orderusecase: useCase}
}

func (h *HttpOrderHandler) CreatOrder(c *fiber.Ctx) error {
	var order entities.Order
	if err := c.BodyParser(&order); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})

	}
	if err := h.Orderusecase.CreatOrder(order); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(order)
}
