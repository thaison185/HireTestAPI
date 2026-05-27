package handlers

import (
	"hiretest-api/internal/common/utils"
	"hiretest-api/internal/requests"
	"hiretest-api/internal/services"

	"github.com/gofiber/fiber/v2"
)

type QuestionHandler struct {
	Service *services.QuestionService
}

func NewQuestionHandler(service *services.QuestionService) *QuestionHandler {
	return &QuestionHandler{Service: service}
}

func (h *QuestionHandler) Create(c *fiber.Ctx) error {
	return utils.Success(c, fiber.StatusOK, "question created", nil)
}
func (h *QuestionHandler) List(c *fiber.Ctx) error {
	page := c.QueryInt("page", 1)
	pageSize := c.QueryInt("page_size", 10)

	query := requests.ListQuestionRequest{
		Keyword:  c.Query("keyword"),
		Type:     c.Query("type"),
		Level:    c.Query("level"),
		Category: c.Query("category"),
		Page:     page,
		PageSize: pageSize,
	}

	results, err := h.Service.List(query)
	if err != nil {
		return utils.Fail(c, fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "questions retrieved successfully",
		"data":    results.Items,
		"meta":    results.Meta,
	})
}
func (h *QuestionHandler) Detail(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return utils.Fail(c, fiber.StatusBadRequest, "question ID is required")
	}

	question, err := h.Service.GetByID(id)
	if err != nil {
		return utils.Fail(c, fiber.StatusNotFound, err.Error())
	}

	return utils.Success(c, fiber.StatusOK, "question detail", question)
}
func (h *QuestionHandler) Update(c *fiber.Ctx) error {
	return utils.Success(c, fiber.StatusOK, "question updated", nil)
}
func (h *QuestionHandler) Delete(c *fiber.Ctx) error {
	return utils.Success(c, fiber.StatusOK, "question deleted", nil)
}
func (h *QuestionHandler) Publish(c *fiber.Ctx) error {
	return utils.Success(c, fiber.StatusOK, "question published", nil)
}
