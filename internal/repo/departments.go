package repo

import "educational_program_creator/internal/models"

func (r *Repository) GetAllDepartments() ([]models.Department, error) {
	var departments []models.Department
	err := r.db.Find(&departments).Error
	if err != nil {
		return nil, err
	}
	return departments, nil
}

// GetDepartmentByID retrieves a department by its ID from the database.
func (r *Repository) GetDepartmentByID(id int) (*models.Department, error) {
	var department models.Department
	err := r.db.First(&department, id).Error
	if err != nil {
		return nil, err
	}
	return &department, nil
}

// CreateDepartment creates a new department in the database.
func (r *Repository) CreateDepartment(department *models.Department) error {
	return r.db.Create(department).Error
}

// UpdateDepartment updates an existing department in the database.
func (r *Repository) UpdateDepartment(department *models.Department) error {
	return r.db.Save(department).Error
}

// DeleteDepartment deletes a department from the database by its ID.
func (r *Repository) DeleteDepartment(id int) error {
	return r.db.Delete(&models.Department{}, id).Error
}
