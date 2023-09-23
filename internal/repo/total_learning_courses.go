package repo

import "educational_program_creator/internal/models"

func (r *Repository) GetTotalLearningCourseByID(id int) (*models.TotalLearningCourse, error) {
	var totalLearningCourse models.TotalLearningCourse
	err := r.db.First(&totalLearningCourse, id).Error
	if err != nil {
		return nil, err
	}
	return &totalLearningCourse, nil
}

func (r *Repository) CreateTotalLearningCourse(totalLearningCourse *models.TotalLearningCourse) error {
	return r.db.Create(totalLearningCourse).Error
}

func (r *Repository) UpdateTotalLearningCourse(totalLearningCourse *models.TotalLearningCourse) error {
	return r.db.Save(totalLearningCourse).Error
}

func (r *Repository) DeleteTotalLearningCourse(id int) error {
	return r.db.Delete(&models.TotalLearningCourse{}, id).Error
}
