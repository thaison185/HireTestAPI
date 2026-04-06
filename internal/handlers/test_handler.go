package handlers

import "github.com/gofiber/fiber/v2"

type TestHandler struct{}

func (h *TestHandler) Create(c *fiber.Ctx) error { return ok(c, "test created", nil) }
func (h *TestHandler) List(c *fiber.Ctx) error { return ok(c, "test list", []string{}) }
func (h *TestHandler) Detail(c *fiber.Ctx) error { return ok(c, "test detail", fiber.Map{"id": c.Params("testId")}) }
func (h *TestHandler) Update(c *fiber.Ctx) error { return ok(c, "test updated", nil) }
func (h *TestHandler) Delete(c *fiber.Ctx) error { return ok(c, "test deleted", nil) }
func (h *TestHandler) AddQuestion(c *fiber.Ctx) error { return ok(c, "question added to test", nil) }
func (h *TestHandler) Publish(c *fiber.Ctx) error { return ok(c, "test published", nil) }
