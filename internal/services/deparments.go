package services

import (
	"educational_program_creator/internal/models"
	"educational_program_creator/pkg/utils"
)

func (s *Service) GetAllDepartments(page int, perPage int) ([]models.Department, error) {
	limit, offset := utils.CalculateLimitOffset(page, perPage)
	departments, err := s.repo.GetAllDepartments(limit, offset)
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

func (s *Service) CreateDepartment(department *models.Department) (int, error) {
	return s.repo.CreateDepartment(department)
}

func (s *Service) UpdateDepartment(department *models.Department) error {
	return s.repo.UpdateDepartment(department)
}

func (s *Service) DeleteDepartment(id int) error {
	return s.repo.DeleteDepartment(id)
}
