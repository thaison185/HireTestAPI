package handlers

import "github.com/gofiber/fiber/v2"

type CandidateHandler struct{}

func (h *CandidateHandler) Create(c *fiber.Ctx) error { return ok(c, "candidate created", nil) }
func (h *CandidateHandler) List(c *fiber.Ctx) error { return ok(c, "candidate list", []string{}) }
func (h *CandidateHandler) Detail(c *fiber.Ctx) error { return ok(c, "candidate detail", fiber.Map{"id": c.Params("candidateId")}) }
func (h *CandidateHandler) Update(c *fiber.Ctx) error { return ok(c, "candidate updated", nil) }
