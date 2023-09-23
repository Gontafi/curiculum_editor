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
	CodeKz        string      `gorm:"column:code_kz;not null" json:"code_kz"`
	CodeRu        string      `gorm:"column:code_ru;not null" json:"code_ru"`
	CodeEn        string      `gorm:"column:code_en;not null" json:"code_en"`
	DescriptionKz string      `gorm:"column:description_kz;not null" json:"description_kz"`
	DescriptionRu string      `gorm:"column:description_ru;not null" json:"description_ru"`
	DescriptionEn string      `gorm:"column:description_en;not null" json:"description_en"`
	Order         int         `gorm:"column:order;not null" json:"order"`
	Components    []Component `gorm:"foreignKey:ProfID"`
}

type Component struct {
	ID                    int                   `gorm:"primaryKey" json:"id"`
	ProfID                int                   `gorm:"column:prof_id;not null" json:"prof_id"`
	CodeKz                string                `gorm:"column:code_kz;not null" json:"code_kz"`
	CodeRu                string                `gorm:"column:code_ru;not null" json:"code_ru"`
	CodeEn                string                `gorm:"column:code_en;not null" json:"code_en"`
	DescriptionKz         string                `gorm:"column:description_kz;not null" json:"description_kz"`
	DescriptionRu         string                `gorm:"column:description_ru;not null" json:"description_ru"`
	DescriptionEn         string                `gorm:"column:description_en;not null" json:"description_en"`
	Order                 int                   `gorm:"column:order;not null" json:"order"`
	ProfessionalComponent ProfessionalComponent `gorm:"foreignKey:ProfID"`
}

type Module struct {
	ID     int    `gorm:"primaryKey" json:"id"`
	Code   string `gorm:"column:code;not null" json:"code"`
	NameKz string `gorm:"column:name_kz;not null" json:"name_kz"`
	NameRu string `gorm:"column:name_ru;not null" json:"name_ru"`
	NameEn string `gorm:"column:name_en;not null" json:"name_en"`

	Cycles []Cycle `gorm:"foreignKey:ModuleID"`
}

type Cycle struct {
	ID       int    `gorm:"primaryKey" json:"id"`
	ModuleID int    `gorm:"column:module_id;not null" json:"module_id"`
	CodeKz   string `gorm:"column:code_kz;not null" json:"code_kz"`
	CodeRu   string `gorm:"column:code_ru;not null" json:"code_ru"`
	CodeEn   string `gorm:"column:code_en;not null" json:"code_en"`

	Module Module `gorm:"foreignKey:ModuleID"`
}

type Department struct {
	ID            int    `gorm:"primaryKey" json:"id"`
	DescriptionKz string `gorm:"column:description_kz;not null" json:"description_kz"`
	DescriptionRu string `gorm:"column:description_ru;not null" json:"description_ru"`
	DescriptionEn string `gorm:"column:description_en;not null" json:"description_en"`
}

type Course struct {
	ID                      int    `gorm:"primaryKey" json:"id"`
	CodeKz                  string `gorm:"column:code_kz;not null" json:"code_kz"`
	CodeRu                  string `gorm:"column:code_ru;not null" json:"code_ru"`
	CodeEn                  string `gorm:"column:code_en;not null" json:"code_en"`
	ECTS                    int    `gorm:"column:ects;not null" json:"ects"`
	ModuleID                int    `gorm:"column:module_id;not null" json:"module_id"`
	DepartmentID            int    `gorm:"column:department_id;not null" json:"department_id"`
	ProfessionalComponentID int    `gorm:"column:professional_component_id;not null" json:"professional_component_id"`
	NameKz                  string `gorm:"column:name_kz;not null" json:"name_kz"`
	NameRu                  string `gorm:"column:name_ru;not null" json:"name_ru"`
	NameEn                  string `gorm:"column:name_en;not null" json:"name_en"`
	LangOfTeachKz           string `gorm:"column:lang_of_teach_kz;not null" json:"lang_of_teach_kz"`
	LangOfTeachRu           string `gorm:"column:lang_of_teach_ru;not null" json:"lang_of_teach_ru"`
	LangOfTeachEn           string `gorm:"column:lang_of_teach_en;not null" json:"lang_of_teach_en"`
	ControlFormKz           string `gorm:"column:control_form_kz;not null" json:"control_form_kz"`
	ControlFormRu           string `gorm:"column:control_form_ru;not null" json:"control_form_ru"`
	ControlFormEn           string `gorm:"column:control_form_en;not null" json:"control_form_en"`
	LectureHour             int    `gorm:"column:lecture_hour;not null" json:"lecture_hour"`
	SeminarHour             int    `gorm:"column:seminar_hour;not null" json:"seminar_hour"`
	LabHour                 int    `gorm:"column:lab_hour;not null" json:"lab_hour"`
	SROHour                 int    `gorm:"column:sro_hour;not null" json:"sro_hour"`

	Module                Module                `gorm:"foreignKey:ModuleID"`
	Department            Department            `gorm:"foreignKey:DepartmentID"`
	ProfessionalComponent ProfessionalComponent `gorm:"foreignKey:ProfessionalComponentID"`
}

type TotalLearningCourse struct {
	ID int `gorm:"primaryKey" json:"id"`
}

type Semester struct {
	ID       int `gorm:"primaryKey" json:"id"`
	CourseID int `gorm:"column:course_id;not null" json:"course_id"`

	Course Course `gorm:"foreignKey:CourseID"`
}

type TotalCourseSemester struct {
	ID                    int `gorm:"primaryKey" json:"id"`
	TotalLearningCourseID int `gorm:"column:total_learning_course_id;not null" json:"total_learning_course_id"`
	SemesterID            int `gorm:"column:semester_id;not null" json:"semester_id"`

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
