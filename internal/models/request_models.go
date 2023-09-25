package models

type CreateUserRequest struct {
	Username     string `json:"username" binding:"required"`
	PasswordHash string `json:"password" binding:"required"`
	Name         string `json:"name" binding:"required"`
}

type UpdateUserRequest struct {
	Username string `json:"username"`
	Name     string `json:"name"`
}

type CreateProfessionalComponentRequest struct {
	CodeKz        string `json:"code_kz" binding:"required"`
	CodeRu        string `json:"code_ru" binding:"required"`
	CodeEn        string `json:"code_en" binding:"required"`
	DescriptionKz string `json:"description_kz" binding:"required"`
	DescriptionRu string `json:"description_ru" binding:"required"`
	DescriptionEn string `json:"description_en" binding:"required"`
	Order         int    `json:"order" binding:"required"`
}

type UpdateProfessionalComponentRequest struct {
	CodeKz        string `json:"code_kz"`
	CodeRu        string `json:"code_ru"`
	CodeEn        string `json:"code_en"`
	DescriptionKz string `json:"description_kz"`
	DescriptionRu string `json:"description_ru"`
	DescriptionEn string `json:"description_en"`
	Order         int    `json:"order"`
}

type CreateComponentRequest struct {
	ProfID        int    `json:"prof_id" binding:"required"`
	CodeKz        string `json:"code_kz" binding:"required"`
	CodeRu        string `json:"code_ru" binding:"required"`
	CodeEn        string `json:"code_en" binding:"required"`
	DescriptionKz string `json:"description_kz" binding:"required"`
	DescriptionRu string `json:"description_ru" binding:"required"`
	DescriptionEn string `json:"description_en" binding:"required"`
	Order         int    `json:"order" binding:"required"`
}

type UpdateComponentRequest struct {
	ProfID        int    `json:"prof_id"`
	CodeKz        string `json:"code_kz"`
	CodeRu        string `json:"code_ru"`
	CodeEn        string `json:"code_en"`
	DescriptionKz string `json:"description_kz"`
	DescriptionRu string `json:"description_ru"`
	DescriptionEn string `json:"description_en"`
	Order         int    `json:"order"`
}

type CreateModuleRequest struct {
	Code   string `json:"code" binding:"required"`
	NameKz string `json:"name_kz" binding:"required"`
	NameRu string `json:"name_ru" binding:"required"`
	NameEn string `json:"name_en" binding:"required"`
}

type UpdateModuleRequest struct {
	Code   string `json:"code"`
	NameKz string `json:"name_kz"`
	NameRu string `json:"name_ru"`
	NameEn string `json:"name_en"`
}

type CreateCycleRequest struct {
	ModuleID int    `json:"module_id" binding:"required"`
	CodeKz   string `json:"code_kz" binding:"required"`
	CodeRu   string `json:"code_ru" binding:"required"`
	CodeEn   string `json:"code_en" binding:"required"`
}

type UpdateCycleRequest struct {
	ModuleID int    `json:"module_id"`
	CodeKz   string `json:"code_kz"`
	CodeRu   string `json:"code_ru"`
	CodeEn   string `json:"code_en"`
}

type CreateDepartmentRequest struct {
	DescriptionKz string `json:"description_kz" binding:"required"`
	DescriptionRu string `json:"description_ru" binding:"required"`
	DescriptionEn string `json:"description_en" binding:"required"`
}

type UpdateDepartmentRequest struct {
	DescriptionKz string `json:"description_kz"`
	DescriptionRu string `json:"description_ru"`
	DescriptionEn string `json:"description_en"`
}

type CreateCourseRequest struct {
	CodeKz                  string `json:"code_kz" binding:"required"`
	CodeRu                  string `json:"code_ru" binding:"required"`
	CodeEn                  string `json:"code_en" binding:"required"`
	ECTS                    int    `json:"ects" binding:"required"`
	ModuleID                int    `json:"module_id" binding:"required"`
	DepartmentID            int    `json:"department_id" binding:"required"`
	ProfessionalComponentID int    `json:"professional_component_id" binding:"required"`
	NameKz                  string `json:"name_kz" binding:"required"`
	NameRu                  string `json:"name_ru" binding:"required"`
	NameEn                  string `json:"name_en" binding:"required"`
	LangOfTeachKz           string `json:"lang_of_teach_kz" binding:"required"`
	LangOfTeachRu           string `json:"lang_of_teach_ru" binding:"required"`
	LangOfTeachEn           string `json:"lang_of_teach_en" binding:"required"`
	ControlFormKz           string `json:"control_form_kz" binding:"required"`
	ControlFormRu           string `json:"control_form_ru" binding:"required"`
	ControlFormEn           string `json:"control_form_en" binding:"required"`
	LectureHour             int    `json:"lecture_hour" binding:"required"`
	SeminarHour             int    `json:"seminar_hour" binding:"required"`
	LabHour                 int    `json:"lab_hour" binding:"required"`
	SROHour                 int    `json:"sro_hour" binding:"required"`
}

type UpdateCourseRequest struct {
	CodeKz                  string `json:"code_kz"`
	CodeRu                  string `json:"code_ru"`
	CodeEn                  string `json:"code_en"`
	ECTS                    int    `json:"ects"`
	ModuleID                int    `json:"module_id"`
	DepartmentID            int    `json:"department_id"`
	ProfessionalComponentID int    `json:"professional_component_id"`
	NameKz                  string `json:"name_kz"`
	NameRu                  string `json:"name_ru"`
	NameEn                  string `json:"name_en"`
	LangOfTeachKz           string `json:"lang_of_teach_kz"`
	LangOfTeachRu           string `json:"lang_of_teach_ru"`
	LangOfTeachEn           string `json:"lang_of_teach_en"`
	ControlFormKz           string `json:"control_form_kz"`
	ControlFormRu           string `json:"control_form_ru"`
	ControlFormEn           string `json:"control_form_en"`
	LectureHour             int    `json:"lecture_hour"`
	SeminarHour             int    `json:"seminar_hour"`
	LabHour                 int    `json:"lab_hour"`
	SROHour                 int    `json:"sro_hour"`
}

type CreateSemesterCourse struct {
	CourseID   int `json:"course_id"`
	SemesterID int `json:"semester_id"`
}

type UpdateSemesterCourse struct {
	ID         int `json:"id"`
	CourseID   int `json:"course_id"`
	SemesterID int `json:"semester_id"`
}

type CreateTotalCourseSemesterRequest struct {
	TotalLearningCourseID int `json:"total_learning_course_id" binding:"required"`
	SemesterID            int `json:"semester_id" binding:"required"`
}

type UpdateTotalCourseSemesterRequest struct {
	TotalLearningCourseID int `json:"total_learning_course_id"`
	SemesterID            int `json:"semester_id"`
}

type CreateCoursePrerequisiteRequest struct {
	CourseID       int `json:"course_id" binding:"required"`
	PrerequisiteID int `json:"prerequisite_id" binding:"required"`
}

type UpdateCoursePrerequisiteRequest struct {
	CourseID       int `json:"course_id"`
	PrerequisiteID int `json:"prerequisite_id"`
}
