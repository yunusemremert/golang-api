package app

import (
	"github.com/gofiber/fiber/v2"
	"golang-api/models"
	"golang-api/services"
	"net/http"
)

type TodoHandler struct {
	Service services.TodoService
}

func (h TodoHandler) CreateTodo(c *fiber.Ctx) error {
	var todo models.Todo

	if err := c.BodyParser(&todo); err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	result, err := h.Service.TodoInsert(todo)
	if err != nil || result.Status == false {
		return c.Status(http.StatusBadGateway).JSON(err.Error())
	}

	return c.Status(http.StatusCreated).JSON(result)
}
