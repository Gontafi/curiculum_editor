package restapi

import (
	"educational_program_creator/internal/models"
	"github.com/gofiber/fiber/v2"
	"log"
	"net/http"
	"strconv"
)

func (h *CrudHandler) GetAllCoursePrerequisites(c *fiber.Ctx) error {
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

	prerequisites, err := h.Service.GetAllCoursePrerequisite(page, perPage)
	if err != nil {
		log.Println(err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve semesters"})
	}
	return c.Status(http.StatusOK).JSON(prerequisites)
}

func (h *CrudHandler) CreateCoursePrerequisite(c *fiber.Ctx) error {
	var req models.CreateCoursePrerequisiteRequest
	if err := c.BodyParser(&req); err != nil {
		log.Println(err)
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	coursePrerequisite := models.CoursePrerequisite{
		CourseID:       req.CourseID,
		PrerequisiteID: req.PrerequisiteID,
	}

	id, err := h.Service.CreateCoursePrerequisite(&coursePrerequisite)
	if err != nil {
		log.Println(err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create course prerequisite"})
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"message": "Course prerequisite created successfully",
		"id":      id,
	})
}

func (h *CrudHandler) GetCoursePrerequisiteByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		log.Println(err)
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	coursePrerequisite, err := h.Service.GetCoursePrerequisiteByID(id)
	if err != nil {
		log.Println(err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve course prerequisite"})
	}

	if coursePrerequisite == nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "Course prerequisite not found"})
	}

	return c.Status(http.StatusOK).JSON(coursePrerequisite)
}

func (h *CrudHandler) UpdateCoursePrerequisite(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		log.Println(err)
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	var req models.UpdateCoursePrerequisiteRequest
	if err := c.BodyParser(&req); err != nil {
		log.Println(err)
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	coursePrerequisite := models.CoursePrerequisite{
		ID:             id,
		CourseID:       req.CourseID,
		PrerequisiteID: req.PrerequisiteID,
	}

	if err := h.Service.UpdateCoursePrerequisite(&coursePrerequisite); err != nil {
		log.Println(err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update course prerequisite"})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "Course prerequisite updated successfully"})
}

func (h *CrudHandler) DeleteCoursePrerequisite(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		log.Println(err)
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	if err := h.Service.DeleteCoursePrerequisite(id); err != nil {
		log.Println(err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete course prerequisite"})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "Course prerequisite deleted successfully"})
}
