package repo

import (
	"context"
	"educational_program_creator/internal/models"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type RepositoryInterface interface {
	GetAllCourses(ctx context.Context, limit int, offset int) ([]models.Course, error)
	GetCourseByID(ctx context.Context, id int) (*models.Course, error)
	getCachedCourses(ctx context.Context) ([]models.Course, error)
	cacheCourses(ctx context.Context, courses []models.Course) error
	getCachedCourse(ctx context.Context, id int) (*models.Course, error)
	cacheCourse(ctx context.Context, course *models.Course) error
	CreateCourse(ctx context.Context, course *models.Course) error
	UpdateCourse(ctx context.Context, course *models.Course) error
	DeleteCourse(ctx context.Context, id int) error
	clearCachedCourses(ctx context.Context)
	clearCachedCourse(ctx context.Context, id int)
	GetAllDepartments(ctx context.Context, limit int, offset int) ([]models.Department, error)
	GetDepartmentByID(ctx context.Context, id int) (*models.Department, error)
	getCachedDepartments(ctx context.Context) ([]models.Department, error)
	cacheDepartments(ctx context.Context, departments []models.Department) error
	getCachedDepartment(ctx context.Context, id int) (*models.Department, error)
	cacheDepartment(ctx context.Context, department *models.Department) error
	CreateDepartment(ctx context.Context, department *models.Department) error
	UpdateDepartment(ctx context.Context, department *models.Department) error
	DeleteDepartment(ctx context.Context, id int) error
	clearCachedDepartments(ctx context.Context)
	clearCachedDepartment(ctx context.Context, id int)
	GetAllTotalCourseSemesters(limit int, offset int) ([]models.TotalCourseSemester, error)
	GetTotalCourseSemesterByID(id int) (*models.TotalCourseSemester, error)
	CreateTotalCourseSemester(totalCourseSemester *models.TotalCourseSemester) error
	UpdateTotalCourseSemester(totalCourseSemester *models.TotalCourseSemester) error
	DeleteTotalCourseSemester(id int) error
	GetAllCycles(ctx context.Context, limit int, offset int) ([]models.Cycle, error)
	GetCycleByID(ctx context.Context, id int) (*models.Cycle, error)
	getCachedCycles(ctx context.Context) ([]models.Cycle, error)
	cacheCycles(ctx context.Context, cycles []models.Cycle) error
	getCachedCycle(ctx context.Context, id int) (*models.Cycle, error)
	cacheCycle(ctx context.Context, cycle *models.Cycle) error
	CreateCycle(ctx context.Context, cycle *models.Cycle) error
	UpdateCycle(ctx context.Context, cycle *models.Cycle) error
	DeleteCycle(ctx context.Context, id int) error
	clearCachedCycles(ctx context.Context)
	clearCachedCycle(ctx context.Context, id int)
	GetAllProfessionalComponents() ([]models.ProfessionalComponent, error)
	GetProfessionalComponentByID(id int) (*models.ProfessionalComponent, error)
	CreateProfessionalComponent(profComponent *models.ProfessionalComponent) error
	UpdateProfessionalComponent(profComponent *models.ProfessionalComponent) error
	DeleteProfessionalComponent(id int) error
	AddUser(user models.User) (int, error)
	UpdateUser(user models.User) error
	GetAllUsers() ([]models.User, error)
	GetUserById(id int) (*models.User, error)
	GetUserByUsername(username string) (*models.User, error)
	AddRole(role models.Role) (int, error)
	DeleteRole(id int) error
	UpdateRole(role models.Role) error
	AllRoles() ([]models.Role, error)
	GetRoleByID(id int) (models.Role, error)
	GetTotalLearningCourseByID(id int) (*models.TotalLearningCourse, error)
	CreateTotalLearningCourse(totalLearningCourse *models.TotalLearningCourse) error
	UpdateTotalLearningCourse(totalLearningCourse *models.TotalLearningCourse) error
	DeleteTotalLearningCourse(id int) error
	GetUserRole(userID int) (*models.Role, error)
	GetAllComponents(ctx context.Context, limit int, offset int) ([]models.Component, error)
	GetComponentByID(ctx context.Context, id int) (*models.Component, error)
	getCachedComponents(ctx context.Context) ([]models.Component, error)
	cacheComponents(ctx context.Context, components []models.Component) error
	getCachedComponent(ctx context.Context, id int) (*models.Component, error)
	cacheComponent(ctx context.Context, component *models.Component) error
	CreateComponent(ctx context.Context, component *models.Component) error
	UpdateComponent(ctx context.Context, component *models.Component) error
	DeleteComponent(ctx context.Context, id int) error
	clearCachedComponents(ctx context.Context)
	clearCachedComponent(ctx context.Context, id int)
	GetAllCoursePrerequisite() ([]models.CoursePrerequisite, error)
	CreateCoursePrerequisite(coursePrerequisite *models.CoursePrerequisite) error
	GetCoursePrerequisiteByID(id int) (*models.CoursePrerequisite, error)
	UpdateCoursePrerequisite(coursePrerequisite *models.CoursePrerequisite) error
	DeleteCoursePrerequisite(id int) error
	GetAllSemesters(limit int, offset int) ([]models.Semester, error)
	GetSemesterByID(id int) (*models.Semester, error)
	CreateSemester(semester *models.Semester) error
	UpdateSemester(semester *models.Semester) error
	DeleteSemester(id int) error
	GetAllModules(ctx context.Context, limit int, offset int) ([]models.Module, error)
	GetModuleByID(ctx context.Context, id int) (*models.Module, error)
	getCachedModules(ctx context.Context) ([]models.Module, error)
	cacheModules(ctx context.Context, modules []models.Module) error
	getCachedModule(ctx context.Context, id int) (*models.Module, error)
	cacheModule(ctx context.Context, module *models.Module) error
	CreateModule(ctx context.Context, module *models.Module) error
	UpdateModule(ctx context.Context, module *models.Module) error
	DeleteModule(ctx context.Context, id int) error
	clearCachedModules(ctx context.Context)
	clearCachedModule(ctx context.Context, id int)
}

type Repository struct {
	db *gorm.DB
	rd *redis.Client
}

func NewRepository(db *gorm.DB, rd *redis.Client) RepositoryInterface {
	return &Repository{db: db, rd: rd}
}
