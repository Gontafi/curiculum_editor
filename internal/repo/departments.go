package repo

import (
	"educational_program_creator/internal/models"
)

func (r *Repository) GetAllDepartments(limit int, offset int) ([]models.Department, error) {
	var departmentsFromDB []models.Department
	err := r.db.Limit(limit).Offset(offset).Find(&departmentsFromDB).Error
	if err != nil {
		return nil, err
	}

	return departmentsFromDB, nil
}

func (r *Repository) GetDepartmentByID(id int) (*models.Department, error) {
	var departmentFromDB models.Department
	err := r.db.First(&departmentFromDB, id).Error
	if err != nil {
		return nil, err
	}

	return &departmentFromDB, nil
}

func (r *Repository) CreateDepartment(department *models.Department) (int, error) {
	if err := r.db.Create(&department).Error; err != nil {
		return 0, err
	}

	return department.ID, nil
}

func (r *Repository) UpdateDepartment(department *models.Department) error {
	if err := r.db.Save(department).Error; err != nil {
		return err
	}

	return nil
}

func (r *Repository) DeleteDepartment(id int) error {
	if err := r.db.Delete(&models.Department{}, id).Error; err != nil {
		return err
	}

	return nil
}
