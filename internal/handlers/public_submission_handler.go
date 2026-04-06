package handlers

import "github.com/gofiber/fiber/v2"

type PublicSubmissionHandler struct{}

func (h *PublicSubmissionHandler) Start(c *fiber.Ctx) error { return ok(c, "submission started", nil) }
func (h *PublicSubmissionHandler) Autosave(c *fiber.Ctx) error { return ok(c, "submission autosaved", nil) }
func (h *PublicSubmissionHandler) FinalSubmit(c *fiber.Ctx) error { return ok(c, "submission final submitted", nil) }
