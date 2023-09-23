package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID           int       `gorm:"primaryKey" json:"id"`
	Username     string    `gorm:"uniqueIndex;not null" json:"username"`
	PasswordHash string    `gorm:"not null" json:"-"`
	Name         string    `gorm:"not null" json:"name"`
	CreatedAt    time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

type ProfessionalComponent struct {
	ID            int         `gorm:"primaryKey" json:"id"`
	CodeKz        string      `gorm:"column:code_kz" json:"code_kz"`
	CodeRu        string      `gorm:"column:code_ru" json:"code_ru"`
	CodeEn        string      `gorm:"column:code_en" json:"code_en"`
	DescriptionKz string      `gorm:"column:description_kz" json:"description_kz"`
	DescriptionRu string      `gorm:"column:description_ru" json:"description_ru"`
	DescriptionEn string      `gorm:"column:description_en" json:"description_en"`
	Order         int         `gorm:"column:order" json:"order"`
	Components    []Component `gorm:"foreignKey:ProfID"`
}

type Component struct {
	ID                    int                   `gorm:"primaryKey" json:"id"`
	ProfID                int                   `gorm:"column:prof_id" json:"prof_id"`
	CodeKz                string                `gorm:"column:code_kz" json:"code_kz"`
	CodeRu                string                `gorm:"column:code_ru" json:"code_ru"`
	CodeEn                string                `gorm:"column:code_en" json:"code_en"`
	DescriptionKz         string                `gorm:"column:description_kz" json:"description_kz"`
	DescriptionRu         string                `gorm:"column:description_ru" json:"description_ru"`
	DescriptionEn         string                `gorm:"column:description_en" json:"description_en"`
	Order                 int                   `gorm:"column:order" json:"order"`
	ProfessionalComponent ProfessionalComponent `gorm:"foreignKey:ProfID"`
}

type Module struct {
	ID     int    `gorm:"primaryKey" json:"id"`
	Code   string `gorm:"column:code" json:"code"`
	NameKz string `gorm:"column:name_kz" json:"name_kz"`
	NameRu string `gorm:"column:name_ru" json:"name_ru"`
	NameEn string `gorm:"column:name_en" json:"name_en"`

	Cycles []Cycle `gorm:"foreignKey:ModuleID"`
}

type Cycle struct {
	ID       int    `gorm:"primaryKey" json:"id"`
	ModuleID int    `gorm:"column:module_id" json:"module_id"`
	CodeKz   string `gorm:"column:code_kz" json:"code_kz"`
	CodeRu   string `gorm:"column:code_ru" json:"code_ru"`
	CodeEn   string `gorm:"column:code_en" json:"code_en"`

	Module Module `gorm:"foreignKey:ModuleID"`
}

type Department struct {
	ID            int    `gorm:"primaryKey" json:"id"`
	DescriptionKz string `gorm:"column:description_kz" json:"description_kz"`
	DescriptionRu string `gorm:"column:description_ru" json:"description_ru"`
	DescriptionEn string `gorm:"column:description_en" json:"description_en"`
}

type Course struct {
	ID                      int    `gorm:"primaryKey" json:"id"`
	CodeKz                  string `gorm:"column:code_kz" json:"code_kz"`
	CodeRu                  string `gorm:"column:code_ru" json:"code_ru"`
	CodeEn                  string `gorm:"column:code_en" json:"code_en"`
	ECTS                    int    `gorm:"column:ects" json:"ects"`
	ModuleID                int    `gorm:"column:module_id" json:"module_id"`
	DepartmentID            int    `gorm:"column:department_id" json:"department_id"`
	ProfessionalComponentID int    `gorm:"column:professional_component_id" json:"professional_component_id"`
	NameKz                  string `gorm:"column:name_kz" json:"name_kz"`
	NameRu                  string `gorm:"column:name_ru" json:"name_ru"`
	NameEn                  string `gorm:"column:name_en" json:"name_en"`
	LangOfTeachKz           string `gorm:"column:lang_of_teach_kz" json:"lang_of_teach_kz"`
	LangOfTeachRu           string `gorm:"column:lang_of_teach_ru" json:"lang_of_teach_ru"`
	LangOfTeachEn           string `gorm:"column:lang_of_teach_en" json:"lang_of_teach_en"`
	ControlFormKz           string `gorm:"column:control_form_kz" json:"control_form_kz"`
	ControlFormRu           string `gorm:"column:control_form_ru" json:"control_form_ru"`
	ControlFormEn           string `gorm:"column:control_form_en" json:"control_form_en"`
	LectureHour             int    `gorm:"column:lecture_hour" json:"lecture_hour"`
	SeminarHour             int    `gorm:"column:seminar_hour" json:"seminar_hour"`
	LabHour                 int    `gorm:"column:lab_hour" json:"lab_hour"`
	SROHour                 int    `gorm:"column:sro_hour" json:"sro_hour"`

	Module                Module                `gorm:"foreignKey:ModuleID"`
	Department            Department            `gorm:"foreignKey:DepartmentID"`
	ProfessionalComponent ProfessionalComponent `gorm:"foreignKey:ProfessionalComponentID"`
}

type TotalLearningCourse struct {
	ID int `gorm:"primaryKey" json:"id"`
}

type Semester struct {
	ID       int `gorm:"primaryKey" json:"id"`
	CourseID int `gorm:"column:course_id" json:"course_id"`

	Course Course `gorm:"foreignKey:CourseID"`
}

type TotalCourseSemester struct {
	ID                    int `gorm:"primaryKey" json:"id"`
	TotalLearningCourseID int `gorm:"column:total_learning_course_id" json:"total_learning_course_id"`
	SemesterID            int `gorm:"column:semester_id" json:"semester_id"`

	TotalLearningCourse TotalLearningCourse `gorm:"foreignKey:TotalLearningCourseID"`
	Semester            Semester            `gorm:"foreignKey:SemesterID"`
}

func AutoMigrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&User{},
		&ProfessionalComponent{},
		&Component{},
		&Module{},
		&Cycle{},
		&Department{},
		&Course{},
		&TotalLearningCourse{},
		&Semester{},
		&TotalCourseSemester{},
	)
	if err != nil {
		return
	}
}
