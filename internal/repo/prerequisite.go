package repo

import "educational_program_creator/internal/models"

func (r *Repository) GetAllCoursePrerequisite(limit int, offset int) ([]models.CoursePrerequisite, error) {
	var prerequisites []models.CoursePrerequisite
	err := r.db.Preload("Course").Preload("Prerequisite").Limit(limit).Offset(offset).Find(&prerequisites).Error
	if err != nil {
		return nil, err
	}
	return prerequisites, nil
}

func (r *Repository) CreateCoursePrerequisite(coursePrerequisite *models.CoursePrerequisite) (int, error) {
	err := r.db.Create(&coursePrerequisite).Error
	if err != nil {
		return 0, err
	}

	return coursePrerequisite.ID, nil
}

func (r *Repository) GetCoursePrerequisiteByID(id int) (*models.CoursePrerequisite, error) {
	var coursePrerequisite models.CoursePrerequisite
	if err := r.db.First(&coursePrerequisite, id).Error; err != nil {
		return nil, err
	}
	return &coursePrerequisite, nil
}

func (r *Repository) UpdateCoursePrerequisite(coursePrerequisite *models.CoursePrerequisite) error {
	return r.db.Save(coursePrerequisite).Error
}

func (r *Repository) DeleteCoursePrerequisite(id int) error {
	return r.db.Delete(&models.CoursePrerequisite{}, id).Error
}
