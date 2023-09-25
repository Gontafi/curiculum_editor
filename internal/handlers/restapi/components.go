package restapi

import (
	"educational_program_creator/internal/models"
	"github.com/gofiber/fiber/v2"
	"log"
	"net/http"
	"strconv"
)

func (h *CrudHandler) GetAllComponents(c *fiber.Ctx) error {
	pageParam := c.Params("page")
	perPageParam := c.Params("perPage")

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

	components, err := h.Service.GetAllComponents(c.Context(), page, perPage)
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

	component, err := h.Service.GetComponentByID(c.Context(), id)
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
	err = h.Service.CreateComponent(c.Context(), &component)
	if err != nil {
		log.Println(err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create component"})
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{"message": "Component created successfully"})
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
	err = h.Service.UpdateComponent(c.Context(), &component)
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

	err = h.Service.DeleteComponent(c.Context(), id)
	if err != nil {
		log.Println(err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete component"})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "Component deleted successfully"})
}
