package handlers

import (
	"hiretest-api/internal/common/constants"
	code_errors "hiretest-api/internal/common/errors"
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
	var req requests.CreateQuestionRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.Fail(c, fiber.StatusBadRequest, "invalid request body")
	}

	if err := utils.Validate.Struct(req); err != nil {
		return utils.Fail(c, fiber.StatusBadRequest, utils.ParseValidationError(err))
	}

	userID, err := utils.CurrentUserID(c)
	if err != nil {
		return utils.Fail(c, fiber.StatusUnauthorized, err.Error())
	}

	question, err := h.Service.Create(req, userID)
	if err != nil {
		return utils.Fail(c, fiber.StatusInternalServerError, err.Error())
	}

	return utils.Success(c, fiber.StatusCreated, "question created successfully", question)
}

func (h *QuestionHandler) List(c *fiber.Ctx) error {
	page := c.QueryInt("page", 1)
	pageSize := c.QueryInt("page_size", 10)

	query := requests.ListQuestionRequest{
		Keyword:  c.Query("keyword"),
		Type:     c.Query("type"),
		Level:    c.Query("level"),
		Category: c.Query("category"),
		Status:   c.Query("status", "active"),
		Sort:     c.Query("sort", "created_at.desc"),
		Page:     page,
		PageSize: pageSize,
	}

	role, err := utils.CurrentUserRole(c)
	if err != nil {
		return utils.Fail(c, fiber.StatusUnauthorized, err.Error())
	}

	results, err := h.Service.List(query, role)
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
	id := c.Params("id")
	if id == "" {
		return utils.Fail(c, fiber.StatusBadRequest, "question ID is required")
	}

	var req requests.UpdateQuestionRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.Fail(c, fiber.StatusBadRequest, "invalid request body")
	}

	if err := utils.Validate.Struct(req); err != nil {
		return utils.Fail(c, fiber.StatusBadRequest, utils.ParseValidationError(err))
	}

	userID, err := utils.CurrentUserID(c)
	if err != nil {
		return utils.Fail(c, fiber.StatusUnauthorized, err.Error())
	}

	role, err := utils.CurrentUserRole(c)
	if err != nil {
		return utils.Fail(c, fiber.StatusUnauthorized, err.Error())
	}

	updatedQuestion, err := h.Service.Update(id, req, userID, role)
	if err != nil {
		switch err.Error() {
		case code_errors.CodeQuestionNotFound, code_errors.CodeQuestionNotFoundOrNotOwn:
			return utils.Fail(c, fiber.StatusNotFound, err.Error())
		case code_errors.CodeUnauthorized:
			return utils.Fail(c, fiber.StatusUnauthorized, err.Error())
		case code_errors.CodeNoFieldsToUpdate, code_errors.CodeQuestionIDRequired:
			return utils.Fail(c, fiber.StatusBadRequest, err.Error())
		default:
			return utils.Fail(c, fiber.StatusInternalServerError, err.Error())
		}
	}
	return utils.Success(c, fiber.StatusOK, "question updated", updatedQuestion)
}

func (h *QuestionHandler) Delete(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return utils.Fail(c, fiber.StatusBadRequest, "question ID is required")
	}

	userID, err := utils.CurrentUserID(c)
	if err != nil {
		return utils.Fail(c, fiber.StatusUnauthorized, err.Error())
	}

	role, err := utils.CurrentUserRole(c)
	if err != nil {
		return utils.Fail(c, fiber.StatusUnauthorized, err.Error())
	}

	err = h.Service.Delete(id, userID, role)
	if err != nil {
		switch err.Error() {
		case code_errors.CodeQuestionNotFound, code_errors.CodeQuestionNotFoundOrNotOwn:
			return utils.Fail(c, fiber.StatusNotFound, err.Error())
		case code_errors.CodeUnauthorized:
			return utils.Fail(c, fiber.StatusUnauthorized, err.Error())
		default:
			return utils.Fail(c, fiber.StatusInternalServerError, err.Error())
		}
	}
	return utils.Success(c, fiber.StatusOK, "question deleted", fiber.Map{
		"id":        id,
		"is_active": false,
	})
}

func (h *QuestionHandler) Restore(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return utils.Fail(c, fiber.StatusBadRequest, code_errors.CodeQuestionIDRequired)
	}

	userID, err := utils.CurrentUserID(c)
	if err != nil {
		return utils.Fail(c, fiber.StatusUnauthorized, err.Error())
	}

	role, err := utils.CurrentUserRole(c)
	if err != nil {
		return utils.Fail(c, fiber.StatusUnauthorized, err.Error())
	}

	if role != constants.RoleAdmin {
		return utils.Fail(c, fiber.StatusForbidden, code_errors.CodeForbidden)
	}

	updatedQuestion, err := h.Service.Restore(id, userID, role)
	if err != nil {
		switch err.Error() {
		case code_errors.CodeQuestionNotFound:
			return utils.Fail(c, fiber.StatusNotFound, err.Error())
		case code_errors.CodeQuestionAlreadyActive, code_errors.CodeQuestionIDRequired:
			return utils.Fail(c, fiber.StatusBadRequest, err.Error())
		case code_errors.CodeUnauthorized:
			return utils.Fail(c, fiber.StatusUnauthorized, err.Error())
		case code_errors.CodeForbidden:
			return utils.Fail(c, fiber.StatusForbidden, err.Error())
		default:
			return utils.Fail(c, fiber.StatusInternalServerError, err.Error())
		}
	}
	return utils.Success(c, fiber.StatusOK, "question restored", updatedQuestion)
}

func (h *QuestionHandler) Publish(c *fiber.Ctx) error {
	return utils.Success(c, fiber.StatusOK, "question published", nil)
}
