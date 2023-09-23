package restapi

import (
	"educational_program_creator/internal/models"
	"github.com/gofiber/fiber/v2"
	"log"
	"net/http"
	"strconv"
)

func (h *CrudHandler) GetAllSemesters(c *fiber.Ctx) error {
	semesters, err := h.Service.GetAllSemesters()
	if err != nil {
		log.Println(err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve semesters"})
	}
	return c.Status(http.StatusOK).JSON(semesters)
}

func (h *CrudHandler) GetSemesterByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		log.Println(err)
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	semester, err := h.Service.GetSemesterByID(id)
	if err != nil {
		log.Println(err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve semester"})
	}

	if semester == nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "Semester not found"})
	}

	return c.Status(http.StatusOK).JSON(semester)
}

func (h *CrudHandler) CreateSemester(c *fiber.Ctx) error {
	var req models.CreateSemesterRequest
	if err := c.BodyParser(&req); err != nil {
		log.Println(err)
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	semester := models.Semester{
		CourseID: req.CourseID,
	}

	if err := h.Service.CreateSemester(&semester); err != nil {
		log.Println(err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create semester"})
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{"message": "Semester created successfully"})
}

func (h *CrudHandler) UpdateSemester(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		log.Println(err)
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	var req models.UpdateSemesterRequest
	if err := c.BodyParser(&req); err != nil {
		log.Println(err)
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	semester := models.Semester{
		ID:       id,
		CourseID: req.CourseID,
	}

	if err := h.Service.UpdateSemester(&semester); err != nil {
		log.Println(err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update semester"})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "Semester updated successfully"})
}

func (h *CrudHandler) DeleteSemester(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		log.Println(err)
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	if err := h.Service.DeleteSemester(id); err != nil {
		log.Println(err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete semester"})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "Semester deleted successfully"})
}
