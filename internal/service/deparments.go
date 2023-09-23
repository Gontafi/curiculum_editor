package services

import "educational_program_creator/internal/models"

func (s *Service) GetAllDepartments() ([]models.Department, error) {
	departments, err := s.repo.GetAllDepartments()
	if err != nil {
		return nil, err
	}
	return departments, nil
}

func (s *Service) GetDepartmentByID(id int) (*models.Department, error) {
	department, err := s.repo.GetDepartmentByID(id)
	if err != nil {
		return nil, err
	}
	return department, nil
}

func (s *Service) CreateDepartment(department *models.Department) error {
	return s.repo.CreateDepartment(department)
}

func (s *Service) UpdateDepartment(department *models.Department) error {
	return s.repo.UpdateDepartment(department)
}

func (s *Service) DeleteDepartment(id int) error {
	return s.repo.DeleteDepartment(id)
}
