package restapi

import (
	"educational_program_creator/internal/models"
	"github.com/gofiber/fiber/v2"
	"log"
	"net/http"
	"strconv"
)

func (h *CrudHandler) GetAllCycles(c *fiber.Ctx) error {
	pageParam := c.Query("page", "1")
	perPageParam := c.Query("perPage", "10")

	page, err := strconv.Atoi(pageParam)
	if err != nil {
		log.Println(err)
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}
	perPage, err := strconv.Atoi(perPageParam)
	if err != nil {
		log.Println(err)
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	cycles, err := h.Service.GetAllCycles(c.Context(), page, perPage)
	if err != nil {
		log.Println(err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve cycles"})
	}
	return c.Status(http.StatusOK).JSON(cycles)
}

func (h *CrudHandler) GetCycleByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		log.Println(err)
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	cycle, err := h.Service.GetCycleByID(c.Context(), id)
	if err != nil {
		log.Println(err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve cycle"})
	}

	if cycle == nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "Cycle not found"})
	}

	return c.Status(http.StatusOK).JSON(cycle)
}

func (h *CrudHandler) CreateCycle(c *fiber.Ctx) error {
	var req models.CreateCycleRequest
	if err := c.BodyParser(&req); err != nil {
		log.Println(err)
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	cycle := models.Cycle{
		ModuleID: req.ModuleID,
		CodeKz:   req.CodeKz,
		CodeRu:   req.CodeRu,
		CodeEn:   req.CodeEn,
	}

	id, err := h.Service.CreateCycle(c.Context(), &cycle)
	if err != nil {
		log.Println(err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create cycle"})
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"message": "Cycle created successfully",
		"id":      id,
	})
}

func (h *CrudHandler) UpdateCycle(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		log.Println(err)
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	var req models.UpdateCycleRequest
	if err := c.BodyParser(&req); err != nil {
		log.Println(err)
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	cycle := models.Cycle{
		ID:       id,
		ModuleID: req.ModuleID,
		CodeKz:   req.CodeKz,
		CodeRu:   req.CodeRu,
		CodeEn:   req.CodeEn,
	}

	if err := h.Service.UpdateCycle(c.Context(), &cycle); err != nil {
		log.Println(err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update cycle"})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "Cycle updated successfully"})
}

func (h *CrudHandler) DeleteCycle(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		log.Println(err)
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	if err := h.Service.DeleteCycle(c.Context(), id); err != nil {
		log.Println(err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete cycle"})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "Cycle deleted successfully"})
}
