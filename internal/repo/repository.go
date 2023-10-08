package repo

import (
	"educational_program_creator/internal/models"
	"gorm.io/gorm"
)

type RepositoryInterface interface {
	GetAllCoursePrerequisite(limit int, offset int) ([]models.CoursePrerequisite, error)
	CreateCoursePrerequisite(coursePrerequisite *models.CoursePrerequisite) (int, error)
	GetCoursePrerequisiteByID(id int) (*models.CoursePrerequisite, error)
	UpdateCoursePrerequisite(coursePrerequisite *models.CoursePrerequisite) error
	DeleteCoursePrerequisite(id int) error
	GetAllSemesters(limit int, offset int) ([]models.SemesterCourse, error)
	GetSemesterByID(id int) (*models.SemesterCourse, error)
	CreateSemester(semester *models.SemesterCourse) (int, error)
	UpdateSemester(semester *models.SemesterCourse) error
	DeleteSemester(id int) error
	AddUser(user models.User) (int, error)
	UpdateUser(user models.User) error
	GetAllUsers() ([]models.User, error)
	GetUserById(id int) (*models.User, error)
	GetUserByUsername(username string) (*models.User, error)
	GetAllTotalCourseSemesters(limit int, offset int) ([]models.TotalCourseSemester, error)
	GetTotalCourseSemesterByID(id int) (*models.TotalCourseSemester, error)
	CreateTotalCourseSemester(totalCourseSemester *models.TotalCourseSemester) (int, error)
	UpdateTotalCourseSemester(totalCourseSemester *models.TotalCourseSemester) error
	DeleteTotalCourseSemester(id int) error
	GetAllDepartments(limit int, offset int) ([]models.Department, error)
	GetDepartmentByID(id int) (*models.Department, error)
	CreateDepartment(department *models.Department) (int, error)
	UpdateDepartment(department *models.Department) error
	DeleteDepartment(id int) error
	AddRole(role models.Role) (int, error)
	DeleteRole(id int) error
	UpdateRole(role models.Role) error
	AllRoles() ([]models.Role, error)
	GetRoleByID(id int) (models.Role, error)
	GetTotalLearningCourseByID(id int) (*models.TotalLearningCourse, error)
	CreateTotalLearningCourse(totalLearningCourse *models.TotalLearningCourse) (int, error)
	UpdateTotalLearningCourse(totalLearningCourse *models.TotalLearningCourse) error
	DeleteTotalLearningCourse(id int) error
	GetAllModules(limit int, offset int) ([]models.Module, error)
	GetModuleByID(id int) (*models.Module, error)
	CreateModule(module *models.Module) (int, error)
	UpdateModule(module *models.Module) error
	DeleteModule(id int) error
	GetAllComponents(limit int, offset int) ([]models.Component, error)
	GetComponentByID(id int) (*models.Component, error)
	CreateComponent(component *models.Component) (int, error)
	UpdateComponent(component *models.Component) error
	DeleteComponent(id int) error
	GetAllProfessionalComponents(limit int, offset int) ([]models.ProfessionalComponent, error)
	GetProfessionalComponentByID(id int) (*models.ProfessionalComponent, error)
	CreateProfessionalComponent(profComponent *models.ProfessionalComponent) (int, error)
	UpdateProfessionalComponent(profComponent *models.ProfessionalComponent) error
	DeleteProfessionalComponent(id int) error
	GetAllCourses(limit int, offset int) ([]models.Course, error)
	GetCourseByID(id int) (*models.Course, error)
	CreateCourse(course *models.Course) (int, error)
	UpdateCourse(course *models.Course) error
	DeleteCourse(id int) error
	GetAllCycles(limit int, offset int) ([]models.Cycle, error)
	GetCycleByID(id int) (*models.Cycle, error)
	CreateCycle(cycle *models.Cycle) (int, error)
	UpdateCycle(cycle *models.Cycle) error
	DeleteCycle(id int) error
	GetUserRole(userID int) (*models.Role, error)
}

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) RepositoryInterface {
	return &Repository{db: db}
}
