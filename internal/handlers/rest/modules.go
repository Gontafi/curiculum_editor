package rest

import (
	"educational_program_creator/internal/models"
	"github.com/gofiber/fiber/v2"
	"log"
	"net/http"
	"strconv"
)

func (h *CrudHandler) GetAllModules(c *fiber.Ctx) error {
	modules, err := h.Service.GetAllModules()
	if err != nil {
		log.Println(err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve modules"})
	}
	return c.Status(http.StatusOK).JSON(modules)
}

func (h *CrudHandler) GetModuleByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		log.Println(err)
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	module, err := h.Service.GetModuleByID(id)
	if err != nil {
		log.Println(err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve module"})
	}

	if module == nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "Module not found"})
	}

	return c.Status(http.StatusOK).JSON(module)
}

func (h *CrudHandler) CreateModule(c *fiber.Ctx) error {
	var req models.CreateModuleRequest
	if err := c.BodyParser(&req); err != nil {
		log.Println(err)
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	module := models.Module{
		Code:   req.Code,
		NameKz: req.NameKz,
		NameRu: req.NameRu,
		NameEn: req.NameEn,
	}

	if err := h.Service.CreateModule(&module); err != nil {
		log.Println(err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create module"})
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{"message": "Module created successfully"})
}

func (h *CrudHandler) UpdateModule(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		log.Println(err)
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	var req models.UpdateModuleRequest
	if err := c.BodyParser(&req); err != nil {
		log.Println(err)
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	module := models.Module{
		ID:     id,
		Code:   req.Code,
		NameKz: req.NameKz,
		NameRu: req.NameRu,
		NameEn: req.NameEn,
	}

	if err := h.Service.UpdateModule(&module); err != nil {
		log.Println(err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update module"})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "Module updated successfully"})
}

func (h *CrudHandler) DeleteModule(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		log.Println(err)
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	if err := h.Service.DeleteModule(id); err != nil {
		log.Println(err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete module"})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "Module deleted successfully"})
}
