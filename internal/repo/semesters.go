package repo

import "educational_program_creator/internal/models"

func (r *Repository) GetAllSemesters(limit int, offset int) ([]models.SemesterCourse, error) {
	var semesters []models.SemesterCourse
	err := r.db.Limit(limit).Offset(offset).Find(&semesters).Error
	if err != nil {
		return nil, err
	}
	return semesters, nil
}

func (r *Repository) GetSemesterByID(id int) (*models.SemesterCourse, error) {
	var semester models.SemesterCourse
	err := r.db.First(&semester, id).Error
	if err != nil {
		return nil, err
	}
	return &semester, nil
}

func (r *Repository) CreateSemester(semester *models.SemesterCourse) (int, error) {
	err := r.db.Create(&semester).Error
	if err != nil {
		return 0, err
	}
	return semester.ID, nil
}

func (r *Repository) UpdateSemester(semester *models.SemesterCourse) error {
	return r.db.Save(semester).Error
}

func (r *Repository) DeleteSemester(id int) error {
	return r.db.Delete(&models.SemesterCourse{}, id).Error
}
