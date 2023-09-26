package restapi

import (
	"educational_program_creator/internal/models"
	"github.com/gofiber/fiber/v2"
	"log"
	"net/http"
	"strconv"
)

func (h *CrudHandler) GetAllSemesters(c *fiber.Ctx) error {
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

	semesters, err := h.Service.GetAllSemesters(page, perPage)
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

func (h *CrudHandler) CreateSemesterCourse(c *fiber.Ctx) error {
	var req models.CreateSemesterCourse
	if err := c.BodyParser(&req); err != nil {
		log.Println(err)
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}
	//log help
	semester := models.SemesterCourse{
		CourseID:   req.CourseID,
		SemesterID: req.SemesterID,
	}

	id, err := h.Service.CreateSemester(&semester)
	if err != nil {
		log.Println(err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create semester"})
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"message": "Semester created successfully",
		"id":      id,
	})
}

func (h *CrudHandler) UpdateSemester(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		log.Println(err)
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	var req models.UpdateSemesterCourse
	if err := c.BodyParser(&req); err != nil {
		log.Println(err)
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	semester := models.SemesterCourse{
		ID:         id,
		CourseID:   req.CourseID,
		SemesterID: req.SemesterID,
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
