package app

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/nizsimsek/go-fiber-mongodb/models"
	"github.com/nizsimsek/go-fiber-mongodb/services"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

	if err != nil || !result.Status {
		return err
	}

	return c.Status(http.StatusCreated).JSON(result)
}

func (h TodoHandler) GetAllTodo(c *fiber.Ctx) error {
	result, err := h.Service.TodoGetAll()

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(http.StatusOK).JSON(result)
}

func (h TodoHandler) DeleteTodo(c *fiber.Ctx) error {
	q := c.Params("id")
	cnv, _ := primitive.ObjectIDFromHex(q)

	result, err := h.Service.TodoDelete(cnv)

	if err != nil || !result {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"State": false})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"State": true})
}
