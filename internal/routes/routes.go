package routes

import (
	"educational_program_creator/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

func InitRoutes(h *handlers.Handler, app *fiber.App) {
	api := app.Group("/api")

	auth := api.Group("/auth")
	{
		auth.Post("/sign-up", h.API.SignUp)
		auth.Post("/sign-in", h.API.SignIn)
	}
	rest := api.Group("/")
	{
		components := rest.Group("/component")
		{
			components.Get("/", h.CRUD.GetAllComponents)
			components.Get("/:id", h.CRUD.GetComponentByID)
			components.Post("/", h.CRUD.CreateComponent)
			components.Put("/:id", h.CRUD.UpdateComponent)
			components.Delete("/:id", h.CRUD.DeleteComponent)
		}
		courses := rest.Group("/course")
		{
			courses.Get("/", h.CRUD.GetAllCourses)
			courses.Get("/:id", h.CRUD.GetCourseByID)
			courses.Post("/", h.CRUD.CreateCourse)
			courses.Put("/:id", h.CRUD.UpdateCourse)
			courses.Delete("/:id", h.CRUD.DeleteCourse)
		}
		cycles := rest.Group("/cycle")
		{
			cycles.Get("/", h.CRUD.GetAllCycles)
			cycles.Get("/:id", h.CRUD.GetCycleByID)
			cycles.Post("/", h.CRUD.CreateCycle)
			cycles.Put("/:id", h.CRUD.UpdateCycle)
			cycles.Delete("/:id", h.CRUD.DeleteCycle)
		}
		departments := rest.Group("/department")
		{
			departments.Get("/", h.CRUD.GetAllDepartments)
			departments.Get("/:id", h.CRUD.GetDepartmentByID)
			departments.Post("/", h.CRUD.CreateDepartment)
			departments.Put("/:id", h.CRUD.UpdateDepartment)
			departments.Delete("/:id", h.CRUD.DeleteDepartment)
		}
		modules := rest.Group("/module")
		{
			modules.Get("/", h.CRUD.GetAllModules)
			modules.Get("/:id", h.CRUD.GetModuleByID)
			modules.Post("/", h.CRUD.CreateModule)
			modules.Put("/:id", h.CRUD.UpdateModule)
			modules.Delete("/:id", h.CRUD.DeleteModule)
		}
		professionalComponents := rest.Group("/professional-component")
		{
			professionalComponents.Get("/", h.CRUD.GetAllProfessionalComponents)
			professionalComponents.Get("/:id", h.CRUD.GetProfessionalComponentByID)
			professionalComponents.Post("/", h.CRUD.CreateProfessionalComponent)
			professionalComponents.Put("/:id", h.CRUD.UpdateProfessionalComponent)
			professionalComponents.Delete("/:id", h.CRUD.DeleteProfessionalComponent)
		}
		semesters := rest.Group("/semesters")
		{
			semesters.Get("/", h.CRUD.GetAllSemesters)
			semesters.Get("/:id", h.CRUD.GetSemesterByID)
			semesters.Post("/", h.CRUD.CreateSemester)
			semesters.Put("/:id", h.CRUD.UpdateSemester)
			semesters.Delete("/:id", h.CRUD.DeleteSemester)
		}
		totalCourseSemesters := rest.Group("/total-course-semesters")
		{
			totalCourseSemesters.Get("/", h.CRUD.GetAllTotalCourseSemesters)
			totalCourseSemesters.Get("/:id", h.CRUD.GetTotalCourseSemesterByID)
			totalCourseSemesters.Post("/", h.CRUD.CreateTotalCourseSemester)
			totalCourseSemesters.Put("/:id", h.CRUD.UpdateTotalCourseSemester)
			totalCourseSemesters.Delete("/:id", h.CRUD.DeleteTotalCourseSemester)
		}
		totalLearningCourses := api.Group("/total-learning-courses")
		{
			totalLearningCourses.Get("/:id", h.CRUD.GetTotalLearningCourseByID)
			totalLearningCourses.Post("/", h.CRUD.CreateTotalLearningCourse)
			totalLearningCourses.Delete("/:id", h.CRUD.DeleteTotalLearningCourse)
		}
		users := rest.Group("/users", h.Middleware.AdminRoleMiddleware())
		{
			users.Get("/", h.CRUD.GetAllUsers)
			users.Get("/:id", h.CRUD.GetUserByID)
			users.Get("/username/:username", h.CRUD.GetUserByUsername)
			users.Post("/", h.CRUD.AddUser)
			users.Put("/:id", h.CRUD.UpdateUser)
		}
	}
}
