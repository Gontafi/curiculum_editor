package repo

import "educational_program_creator/internal/models"

func (r *Repository) GetAllCourses() ([]models.Course, error) {
	var courses []models.Course
	err := r.db.Find(&courses).Error
	if err != nil {
		return nil, err
	}
	return courses, nil
}

func (r *Repository) GetCourseByID(id int) (*models.Course, error) {
	var course models.Course
	err := r.db.First(&course, id).Error
	if err != nil {
		return nil, err
	}
	return &course, nil
}

func (r *Repository) CreateCourse(course *models.Course) error {
	return r.db.Create(course).Error
}

func (r *Repository) UpdateCourse(course *models.Course) error {
	return r.db.Save(course).Error
}

func (r *Repository) DeleteCourse(id int) error {
	return r.db.Delete(&models.Course{}, id).Error
}
