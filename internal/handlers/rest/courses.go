package rest

import (
	"educational_program_creator/internal/models"
	"github.com/gofiber/fiber/v2"
	"log"
	"net/http"
	"strconv"
)

func (h *CrudHandler) GetAllCourses(c *fiber.Ctx) error {
	courses, err := h.Service.GetAllCourses(c.Context())
	if err != nil {
		log.Println(err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve courses"})
	}
	return c.Status(http.StatusOK).JSON(courses)
}

func (h *CrudHandler) GetCourseByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		log.Println(err)
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	course, err := h.Service.GetCourseByID(c.Context(), id)
	if err != nil {
		log.Println(err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve course"})
	}

	if course == nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "Course not found"})
	}

	return c.Status(http.StatusOK).JSON(course)
}

func (h *CrudHandler) CreateCourse(c *fiber.Ctx) error {
	var req models.CreateCourseRequest
	if err := c.BodyParser(&req); err != nil {
		log.Println(err)
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	course := models.Course{
		CodeKz:                  req.CodeKz,
		CodeRu:                  req.CodeRu,
		CodeEn:                  req.CodeEn,
		ECTS:                    req.ECTS,
		ModuleID:                req.ModuleID,
		DepartmentID:            req.DepartmentID,
		ProfessionalComponentID: req.ProfessionalComponentID,
		NameKz:                  req.NameKz,
		NameRu:                  req.NameRu,
		NameEn:                  req.NameEn,
		LangOfTeachKz:           req.LangOfTeachKz,
		LangOfTeachRu:           req.LangOfTeachRu,
		LangOfTeachEn:           req.LangOfTeachEn,
		ControlFormKz:           req.ControlFormKz,
		ControlFormRu:           req.ControlFormRu,
		ControlFormEn:           req.ControlFormEn,
		LectureHour:             req.LectureHour,
		SeminarHour:             req.SeminarHour,
		LabHour:                 req.LabHour,
		SROHour:                 req.SROHour,
	}

	if err := h.Service.CreateCourse(c.Context(), &course); err != nil {
		log.Println(err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create course"})
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{"message": "Course created successfully"})
}

func (h *CrudHandler) UpdateCourse(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		log.Println(err)
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	var req models.UpdateCourseRequest
	if err := c.BodyParser(&req); err != nil {
		log.Println(err)
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	course := models.Course{
		ID:                      id,
		CodeKz:                  req.CodeKz,
		CodeRu:                  req.CodeRu,
		CodeEn:                  req.CodeEn,
		ECTS:                    req.ECTS,
		ModuleID:                req.ModuleID,
		DepartmentID:            req.DepartmentID,
		ProfessionalComponentID: req.ProfessionalComponentID,
		NameKz:                  req.NameKz,
		NameRu:                  req.NameRu,
		NameEn:                  req.NameEn,
		LangOfTeachKz:           req.LangOfTeachKz,
		LangOfTeachRu:           req.LangOfTeachRu,
		LangOfTeachEn:           req.LangOfTeachEn,
		ControlFormKz:           req.ControlFormKz,
		ControlFormRu:           req.ControlFormRu,
		ControlFormEn:           req.ControlFormEn,
		LectureHour:             req.LectureHour,
		SeminarHour:             req.SeminarHour,
		LabHour:                 req.LabHour,
		SROHour:                 req.SROHour,
	}

	if err := h.Service.UpdateCourse(c.Context(), &course); err != nil {
		log.Println(err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update course"})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "Course updated successfully"})
}

func (h *CrudHandler) DeleteCourse(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		log.Println(err)
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	if err := h.Service.DeleteCourse(c.Context(), id); err != nil {
		log.Println(err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete course"})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "Course deleted successfully"})
}
