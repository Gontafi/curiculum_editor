package repo

import "educational_program_creator/internal/models"

func (r *Repository) GetAllTotalCourseSemesters() ([]models.TotalCourseSemester, error) {
	var totalCourseSemesters []models.TotalCourseSemester
	err := r.db.Find(&totalCourseSemesters).Error
	if err != nil {
		return nil, err
	}
	return totalCourseSemesters, nil
}

func (r *Repository) GetTotalCourseSemesterByID(id int) (*models.TotalCourseSemester, error) {
	var totalCourseSemester models.TotalCourseSemester
	err := r.db.First(&totalCourseSemester, id).Error
	if err != nil {
		return nil, err
	}
	return &totalCourseSemester, nil
}

func (r *Repository) CreateTotalCourseSemester(totalCourseSemester *models.TotalCourseSemester) error {
	return r.db.Create(totalCourseSemester).Error
}

func (r *Repository) UpdateTotalCourseSemester(totalCourseSemester *models.TotalCourseSemester) error {
	return r.db.Save(totalCourseSemester).Error
}

func (r *Repository) DeleteTotalCourseSemester(id int) error {
	return r.db.Delete(&models.TotalCourseSemester{}, id).Error
}
