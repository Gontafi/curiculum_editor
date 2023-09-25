package restapi

import (
	"educational_program_creator/internal/models"
	"github.com/gofiber/fiber/v2"
	"log"
	"net/http"
	"strconv"
)

func (h *CrudHandler) GetAllDepartments(c *fiber.Ctx) error {
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

	departments, err := h.Service.GetAllDepartments(c.Context(), page, perPage)
	if err != nil {
		log.Println(err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve departments"})
	}
	return c.Status(http.StatusOK).JSON(departments)
}

func (h *CrudHandler) GetDepartmentByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		log.Println(err)
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	department, err := h.Service.GetDepartmentByID(c.Context(), id)
	if err != nil {
		log.Println(err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve department"})
	}

	if department == nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "Department not found"})
	}

	return c.Status(http.StatusOK).JSON(department)
}

func (h *CrudHandler) CreateDepartment(c *fiber.Ctx) error {
	var req models.CreateDepartmentRequest
	if err := c.BodyParser(&req); err != nil {
		log.Println(err)
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	department := models.Department{
		DescriptionKz: req.DescriptionKz,
		DescriptionRu: req.DescriptionRu,
		DescriptionEn: req.DescriptionEn,
	}

	if err := h.Service.CreateDepartment(c.Context(), &department); err != nil {
		log.Println(err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create department"})
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{"message": "Department created successfully"})
}

func (h *CrudHandler) UpdateDepartment(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		log.Println(err)
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	var req models.UpdateDepartmentRequest
	if err := c.BodyParser(&req); err != nil {
		log.Println(err)
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	department := models.Department{
		ID:            id,
		DescriptionKz: req.DescriptionKz,
		DescriptionRu: req.DescriptionRu,
		DescriptionEn: req.DescriptionEn,
	}

	if err := h.Service.UpdateDepartment(c.Context(), &department); err != nil {
		log.Println(err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update department"})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "Department updated successfully"})
}

func (h *CrudHandler) DeleteDepartment(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		log.Println(err)
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	if err := h.Service.DeleteDepartment(c.Context(), id); err != nil {
		log.Println(err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete department"})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "Department deleted successfully"})
}
