package handlers

import "github.com/gofiber/fiber/v2"

type QuestionHandler struct{}

func (h *QuestionHandler) Create(c *fiber.Ctx) error { return ok(c, "question created", nil) }
func (h *QuestionHandler) List(c *fiber.Ctx) error { return ok(c, "question list", []string{}) }
func (h *QuestionHandler) Detail(c *fiber.Ctx) error { return ok(c, "question detail", fiber.Map{"id": c.Params("questionId")}) }
func (h *QuestionHandler) Update(c *fiber.Ctx) error { return ok(c, "question updated", nil) }
func (h *QuestionHandler) Delete(c *fiber.Ctx) error { return ok(c, "question deleted", nil) }
func (h *QuestionHandler) Publish(c *fiber.Ctx) error { return ok(c, "question published", nil) }
