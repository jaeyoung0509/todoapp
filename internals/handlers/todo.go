package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jaeyoung0509/todo/internals/core/service"
)

type TodoHandler struct {
	Service *service.TodoService
}

func NewTodoHandler(s *service.TodoService) *TodoHandler {
	return &TodoHandler{Service: s}
}

func (h *TodoHandler) GetAll(c *fiber.Ctx) error {
	_, err := h.Service.GetAll(c.Context(), "a")
	if err != nil {
		return err
	}
	return nil
}
