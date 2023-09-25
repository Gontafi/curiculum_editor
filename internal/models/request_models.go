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
	Order         int    `json:"order,string" binding:"required"`
}

type UpdateProfessionalComponentRequest struct {
	CodeKz        string `json:"code_kz"`
	CodeRu        string `json:"code_ru"`
	CodeEn        string `json:"code_en"`
	DescriptionKz string `json:"description_kz"`
	DescriptionRu string `json:"description_ru"`
	DescriptionEn string `json:"description_en"`
	Order         int    `json:"order,string"`
}

type CreateComponentRequest struct {
	ProfID        int    `json:"prof_id" binding:"required"`
	CodeKz        string `json:"code_kz" binding:"required"`
	CodeRu        string `json:"code_ru" binding:"required"`
	CodeEn        string `json:"code_en" binding:"required"`
	DescriptionKz string `json:"description_kz" binding:"required"`
	DescriptionRu string `json:"description_ru" binding:"required"`
	DescriptionEn string `json:"description_en" binding:"required"`
	Order         int    `json:"order,string" binding:"required"`
}

type UpdateComponentRequest struct {
	ProfID        int    `json:"prof_id"`
	CodeKz        string `json:"code_kz"`
	CodeRu        string `json:"code_ru"`
	CodeEn        string `json:"code_en"`
	DescriptionKz string `json:"description_kz"`
	DescriptionRu string `json:"description_ru"`
	DescriptionEn string `json:"description_en"`
	Order         int    `json:"order,string"`
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
	ModuleID int    `json:"module_id,string" binding:"required"`
	CodeKz   string `json:"code_kz" binding:"required"`
	CodeRu   string `json:"code_ru" binding:"required"`
	CodeEn   string `json:"code_en" binding:"required"`
}

type UpdateCycleRequest struct {
	ModuleID int    `json:"module_id,string"`
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
	ECTS                    int    `json:"ects,string" binding:"required"`
	ModuleID                int    `json:"module_id,string" binding:"required"`
	DepartmentID            int    `json:"department_id,string" binding:"required"`
	ProfessionalComponentID int    `json:"professional_component_id,string" binding:"required"`
	NameKz                  string `json:"name_kz" binding:"required"`
	NameRu                  string `json:"name_ru" binding:"required"`
	NameEn                  string `json:"name_en" binding:"required"`
	LangOfTeachKz           string `json:"lang_of_teach_kz" binding:"required"`
	LangOfTeachRu           string `json:"lang_of_teach_ru" binding:"required"`
	LangOfTeachEn           string `json:"lang_of_teach_en" binding:"required"`
	ControlFormKz           string `json:"control_form_kz" binding:"required"`
	ControlFormRu           string `json:"control_form_ru" binding:"required"`
	ControlFormEn           string `json:"control_form_en" binding:"required"`
	LectureHour             int    `json:"lecture_hour,string" binding:"required"`
	SeminarHour             int    `json:"seminar_hour,string" binding:"required"`
	LabHour                 int    `json:"lab_hour,string" binding:"required"`
	SROHour                 int    `json:"sro_hour,string" binding:"required"`
}

type UpdateCourseRequest struct {
	CodeKz                  string `json:"code_kz"`
	CodeRu                  string `json:"code_ru"`
	CodeEn                  string `json:"code_en"`
	ECTS                    int    `json:"ects,string"`
	ModuleID                int    `json:"module_id,string"`
	DepartmentID            int    `json:"department_id,string"`
	ProfessionalComponentID int    `json:"professional_component_id,string"`
	NameKz                  string `json:"name_kz"`
	NameRu                  string `json:"name_ru"`
	NameEn                  string `json:"name_en"`
	LangOfTeachKz           string `json:"lang_of_teach_kz"`
	LangOfTeachRu           string `json:"lang_of_teach_ru"`
	LangOfTeachEn           string `json:"lang_of_teach_en"`
	ControlFormKz           string `json:"control_form_kz"`
	ControlFormRu           string `json:"control_form_ru"`
	ControlFormEn           string `json:"control_form_en"`
	LectureHour             int    `json:"lecture_hour,string"`
	SeminarHour             int    `json:"seminar_hour,string"`
	LabHour                 int    `json:"lab_hour,string"`
	SROHour                 int    `json:"sro_hour,string"`
}

type CreateSemesterCourse struct {
	CourseID   int `json:"course_id,string"`
	SemesterID int `json:"semester_id,string"`
}

type UpdateSemesterCourse struct {
	ID         int `json:"id,string"`
	CourseID   int `json:"course_id,string"`
	SemesterID int `json:"semester_id,string"`
}

type CreateTotalCourseSemesterRequest struct {
	TotalLearningCourseID int `json:"total_learning_course_id,string" binding:"required"`
	SemesterID            int `json:"semester_id,string" binding:"required"`
}

type UpdateTotalCourseSemesterRequest struct {
	TotalLearningCourseID int `json:"total_learning_course_id,string"`
	SemesterID            int `json:"semester_id,string"`
}

type CreateCoursePrerequisiteRequest struct {
	CourseID       int `json:"course_id,string" binding:"required"`
	PrerequisiteID int `json:"prerequisite_id,string" binding:"required"`
}

type UpdateCoursePrerequisiteRequest struct {
	CourseID       int `json:"course_id,string"`
	PrerequisiteID int `json:"prerequisite_id,string"`
}
