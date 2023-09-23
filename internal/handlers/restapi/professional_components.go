package restapi

import (
	"educational_program_creator/internal/models"
	"github.com/gofiber/fiber/v2"
	"log"
	"net/http"
	"strconv"
)

func (h *CrudHandler) GetAllProfessionalComponents(c *fiber.Ctx) error {
	components, err := h.Service.GetAllProfessionalComponents()
	if err != nil {
		log.Println(err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve professional components"})
	}
	return c.Status(http.StatusOK).JSON(components)
}

func (h *CrudHandler) GetProfessionalComponentByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		log.Println(err)
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	component, err := h.Service.GetProfessionalComponentByID(id)
	if err != nil {
		log.Println(err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve professional component"})
	}

	if component == nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "Professional component not found"})
	}

	return c.Status(http.StatusOK).JSON(component)
}

func (h *CrudHandler) CreateProfessionalComponent(c *fiber.Ctx) error {
	var req models.CreateProfessionalComponentRequest
	if err := c.BodyParser(&req); err != nil {
		log.Println(err)
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	component := models.ProfessionalComponent{
		CodeKz:        req.CodeKz,
		CodeRu:        req.CodeRu,
		CodeEn:        req.CodeEn,
		DescriptionKz: req.DescriptionKz,
		DescriptionRu: req.DescriptionRu,
		DescriptionEn: req.DescriptionEn,
		Order:         req.Order,
	}

	if err := h.Service.CreateProfessionalComponent(&component); err != nil {
		log.Println(err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create professional component"})
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{"message": "Professional component created successfully"})
}

func (h *CrudHandler) UpdateProfessionalComponent(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		log.Println(err)
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	var req models.UpdateProfessionalComponentRequest
	if err := c.BodyParser(&req); err != nil {
		log.Println(err)
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	component := models.ProfessionalComponent{
		ID:            id,
		CodeKz:        req.CodeKz,
		CodeRu:        req.CodeRu,
		CodeEn:        req.CodeEn,
		DescriptionKz: req.DescriptionKz,
		DescriptionRu: req.DescriptionRu,
		DescriptionEn: req.DescriptionEn,
		Order:         req.Order,
	}

	if err := h.Service.UpdateProfessionalComponent(&component); err != nil {
		log.Println(err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update professional component"})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "Professional component updated successfully"})
}

func (h *CrudHandler) DeleteProfessionalComponent(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		log.Println(err)
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	if err := h.Service.DeleteProfessionalComponent(id); err != nil {
		log.Println(err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete professional component"})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "Professional component deleted successfully"})
}
