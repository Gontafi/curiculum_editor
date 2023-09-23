package rest

import (
	"educational_program_creator/internal/models"
	"github.com/gofiber/fiber/v2"
	"log"
	"net/http"
	"strconv"
)

func (h *CrudHandler) GetTotalLearningCourseByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		log.Println(err)
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	totalLearningCourse, err := h.Service.GetTotalLearningCourseByID(id)
	if err != nil {
		log.Println(err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve total learning course"})
	}

	if totalLearningCourse == nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "Total learning course not found"})
	}

	return c.Status(http.StatusOK).JSON(totalLearningCourse)
}

func (h *CrudHandler) CreateTotalLearningCourse(c *fiber.Ctx) error {
	var totalLearningCourse models.TotalLearningCourse
	err := c.BodyParser(&totalLearningCourse)
	if err != nil {
		log.Println(err)
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	err = h.Service.CreateTotalLearningCourse(&totalLearningCourse)
	if err != nil {
		log.Println(err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create total learning course"})
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{"message": "Total learning course created successfully"})
}

func (h *CrudHandler) DeleteTotalLearningCourse(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		log.Println(err)
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	err = h.Service.DeleteTotalLearningCourse(id)
	if err != nil {
		log.Println(err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete total learning course"})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "Total learning course deleted successfully"})
}
