package handlers

import "github.com/gofiber/fiber/v2"

type ReviewerSubmissionHandler struct{}

func (h *ReviewerSubmissionHandler) Queue(c *fiber.Ctx) error { return ok(c, "review queue", []string{}) }
func (h *ReviewerSubmissionHandler) Detail(c *fiber.Ctx) error { return ok(c, "submission review detail", fiber.Map{"id": c.Params("submissionId")}) }
func (h *ReviewerSubmissionHandler) Score(c *fiber.Ctx) error { return ok(c, "submission scored", nil) }
