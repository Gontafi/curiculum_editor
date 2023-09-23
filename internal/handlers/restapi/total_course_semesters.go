package restapi

import (
	"educational_program_creator/internal/models"
	"github.com/gofiber/fiber/v2"
	"log"
	"net/http"
	"strconv"
)

func (h *CrudHandler) GetAllTotalCourseSemesters(c *fiber.Ctx) error {
	totalCourseSemesters, err := h.Service.GetAllTotalCourseSemesters()
	if err != nil {
		log.Println(err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve total course semesters"})
	}
	return c.Status(http.StatusOK).JSON(totalCourseSemesters)
}

func (h *CrudHandler) GetTotalCourseSemesterByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		log.Println(err)
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	totalCourseSemester, err := h.Service.GetTotalCourseSemesterByID(id)
	if err != nil {
		log.Println(err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve total course semester"})
	}

	if totalCourseSemester == nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "Total course semester not found"})
	}

	return c.Status(http.StatusOK).JSON(totalCourseSemester)
}

func (h *CrudHandler) CreateTotalCourseSemester(c *fiber.Ctx) error {
	var totalCourseSemester models.TotalCourseSemester
	err := c.BodyParser(&totalCourseSemester)
	if err != nil {
		log.Println(err)
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	err = h.Service.CreateTotalCourseSemester(&totalCourseSemester)
	if err != nil {
		log.Println(err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create total course semester"})
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{"message": "Total course semester created successfully"})
}

func (h *CrudHandler) UpdateTotalCourseSemester(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		log.Println(err)
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	var totalCourseSemester models.TotalCourseSemester
	err = c.BodyParser(&totalCourseSemester)
	if err != nil {
		log.Println(err)
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	totalCourseSemester.ID = id
	err = h.Service.UpdateTotalCourseSemester(&totalCourseSemester)
	if err != nil {
		log.Println(err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update total course semester"})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "Total course semester updated successfully"})
}

func (h *CrudHandler) DeleteTotalCourseSemester(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		log.Println(err)
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	err = h.Service.DeleteTotalCourseSemester(id)
	if err != nil {
		log.Println(err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete total course semester"})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "Total course semester deleted successfully"})
}
