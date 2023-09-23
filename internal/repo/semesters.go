package repo

import "educational_program_creator/internal/models"

func (r *Repository) GetAllSemesters() ([]models.Semester, error) {
	var semesters []models.Semester
	err := r.db.Find(&semesters).Error
	if err != nil {
		return nil, err
	}
	return semesters, nil
}

func (r *Repository) GetSemesterByID(id int) (*models.Semester, error) {
	var semester models.Semester
	err := r.db.First(&semester, id).Error
	if err != nil {
		return nil, err
	}
	return &semester, nil
}

func (r *Repository) CreateSemester(semester *models.Semester) error {
	return r.db.Create(semester).Error
}

func (r *Repository) UpdateSemester(semester *models.Semester) error {
	return r.db.Save(semester).Error
}

func (r *Repository) DeleteSemester(id int) error {
	return r.db.Delete(&models.Semester{}, id).Error
}
