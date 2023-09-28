package repo

import "educational_program_creator/internal/models"

func (r *Repository) GetAllTotalCourseSemesters(limit int, offset int) ([]models.TotalCourseSemester, error) {
	var totalCourseSemesters []models.TotalCourseSemester
	err := r.db.Preload("Course").Preload("Semester").Limit(limit).Offset(offset).Find(&totalCourseSemesters).Error
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

func (r *Repository) CreateTotalCourseSemester(totalCourseSemester *models.TotalCourseSemester) (int, error) {
	err := r.db.Create(&totalCourseSemester).Error
	if err != nil {
		return 0, err
	}
	return totalCourseSemester.ID, err
}

func (r *Repository) UpdateTotalCourseSemester(totalCourseSemester *models.TotalCourseSemester) error {
	return r.db.Save(totalCourseSemester).Error
}

func (r *Repository) DeleteTotalCourseSemester(id int) error {
	return r.db.Delete(&models.TotalCourseSemester{}, id).Error
}
