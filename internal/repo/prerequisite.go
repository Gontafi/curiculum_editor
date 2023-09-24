package repo

import "educational_program_creator/internal/models"

func (r *Repository) CreateCoursePrerequisite(coursePrerequisite *models.CoursePrerequisite) error {
	return r.db.Create(coursePrerequisite).Error
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
