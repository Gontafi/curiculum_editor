package repo

import (
	"educational_program_creator/internal/models"
	"fmt"
)

func (r *Repository) GetAllCourses(limit int, offset int) ([]models.Course, error) {
	var coursesFromDB []models.Course
	err := r.db.Preload("Module").Preload("Department").Preload("ProfessionalComponent").Preload("Prerequisites").Limit(limit).Offset(offset).Find(&coursesFromDB).Error
	if err != nil {
		return nil, err
	}

	if err != nil {
		fmt.Println("Failed to cache courses in Redis:", err)
	}

	return coursesFromDB, nil
}

func (r *Repository) GetCourseByID(id int) (*models.Course, error) {

	var courseFromDB models.Course
	err := r.db.First(&courseFromDB, id).Error
	if err != nil {
		return nil, err
	}

	return &courseFromDB, nil
}

func (r *Repository) CreateCourse(course *models.Course) (int, error) {
	if err := r.db.Create(&course).Error; err != nil {
		return 0, err
	}

	return course.ID, nil
}

func (r *Repository) UpdateCourse(course *models.Course) error {
	if err := r.db.Save(course).Error; err != nil {
		return err
	}

	return nil
}

func (r *Repository) DeleteCourse(id int) error {
	if err := r.db.Delete(&models.Course{}, id).Error; err != nil {
		return err
	}

	return nil
}
