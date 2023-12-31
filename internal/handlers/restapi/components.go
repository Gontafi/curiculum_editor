package restapi

import (
	"educational_program_creator/internal/models"
	"github.com/gofiber/fiber/v2"
	"log"
	"net/http"
	"strconv"
)

func (h *CrudHandler) GetAllComponents(c *fiber.Ctx) error {
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

	components, err := h.Service.GetAllComponents(page, perPage)
	if err != nil {
		log.Println(err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve components"})
	}
	return c.Status(http.StatusOK).JSON(components)
}

func (h *CrudHandler) GetComponentByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		log.Println(err)
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	component, err := h.Service.GetComponentByID(id)
	if err != nil {
		log.Println(err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve component"})
	}

	if component == nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "Component not found"})
	}

	return c.Status(http.StatusOK).JSON(component)
}

func (h *CrudHandler) CreateComponent(c *fiber.Ctx) error {
	var inputComponent models.CreateComponentRequest
	err := c.BodyParser(&inputComponent)
	if err != nil {
		log.Println(err)
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	component := models.Component{
		ProfID:        inputComponent.ProfID,
		CodeKz:        inputComponent.CodeKz,
		CodeRu:        inputComponent.CodeRu,
		CodeEn:        inputComponent.CodeEn,
		DescriptionKz: inputComponent.DescriptionKz,
		DescriptionRu: inputComponent.DescriptionRu,
		DescriptionEn: inputComponent.DescriptionEn,
		Order:         inputComponent.Order,
	}
	id, err := h.Service.CreateComponent(&component)
	if err != nil {
		log.Println(err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create component"})
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"message": "Component created successfully",
		"id":      id,
	})
}

func (h *CrudHandler) UpdateComponent(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		log.Println(err)
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}
	var inputComponent models.UpdateComponentRequest
	err = c.BodyParser(&inputComponent)
	if err != nil {
		log.Println(err)
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	component := models.Component{
		ID:            id,
		ProfID:        inputComponent.ProfID,
		CodeKz:        inputComponent.CodeKz,
		CodeRu:        inputComponent.CodeRu,
		CodeEn:        inputComponent.CodeEn,
		DescriptionKz: inputComponent.DescriptionKz,
		DescriptionRu: inputComponent.DescriptionRu,
		DescriptionEn: inputComponent.DescriptionEn,
		Order:         inputComponent.Order,
	}
	component.ID = id
	err = h.Service.UpdateComponent(&component)
	if err != nil {
		log.Println(err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update component"})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "Component updated successfully"})
}

func (h *CrudHandler) DeleteComponent(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		log.Println(err)
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	err = h.Service.DeleteComponent(id)
	if err != nil {
		log.Println(err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete component"})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "Component deleted successfully"})
}
