package repo

import (
	"educational_program_creator/internal/models"
	"gorm.io/gorm"
)

type RepositoryInterface interface {
	AddUser(user models.User) (int, error)
	UpdateUser(user models.User) error
	GetAllUsers() ([]models.User, error)
	GetUserById(id int) (*models.User, error)
	GetUserByUsername(username string) (*models.User, error)
	GetAllComponents() ([]models.Component, error)
	GetComponentByID(id int) (*models.Component, error)
	CreateComponent(component *models.Component) error
	UpdateComponent(component *models.Component) error
	DeleteComponent(id int) error
	GetAllCourses() ([]models.Course, error)
	GetCourseByID(id int) (*models.Course, error)
	CreateCourse(course *models.Course) error
	UpdateCourse(course *models.Course) error
	DeleteCourse(id int) error
	GetAllCycles() ([]models.Cycle, error)
	GetCycleByID(id int) (*models.Cycle, error)
	CreateCycle(cycle *models.Cycle) error
	UpdateCycle(cycle *models.Cycle) error
	DeleteCycle(id int) error
	GetAllDepartments() ([]models.Department, error)
	GetDepartmentByID(id int) (*models.Department, error)
	CreateDepartment(department *models.Department) error
	UpdateDepartment(department *models.Department) error
	DeleteDepartment(id int) error
	GetAllModules() ([]models.Module, error)
	GetModuleByID(id int) (*models.Module, error)
	CreateModule(module *models.Module) error
	UpdateModule(module *models.Module) error
	DeleteModule(id int) error
	GetAllProfessionalComponents() ([]models.ProfessionalComponent, error)
	GetProfessionalComponentByID(id int) (*models.ProfessionalComponent, error)
	CreateProfessionalComponent(profComponent *models.ProfessionalComponent) error
	UpdateProfessionalComponent(profComponent *models.ProfessionalComponent) error
	DeleteProfessionalComponent(id int) error
	GetAllSemesters() ([]models.Semester, error)
	GetSemesterByID(id int) (*models.Semester, error)
	CreateSemester(semester *models.Semester) error
	UpdateSemester(semester *models.Semester) error
	DeleteSemester(id int) error
	GetAllTotalCourseSemesters() ([]models.TotalCourseSemester, error)
	GetTotalCourseSemesterByID(id int) (*models.TotalCourseSemester, error)
	CreateTotalCourseSemester(totalCourseSemester *models.TotalCourseSemester) error
	UpdateTotalCourseSemester(totalCourseSemester *models.TotalCourseSemester) error
	DeleteTotalCourseSemester(id int) error
	GetTotalLearningCourseByID(id int) (*models.TotalLearningCourse, error)
	CreateTotalLearningCourse(totalLearningCourse *models.TotalLearningCourse) error
	UpdateTotalLearningCourse(totalLearningCourse *models.TotalLearningCourse) error
	DeleteTotalLearningCourse(id int) error
}

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) RepositoryInterface {
	return &Repository{db: db}
}
