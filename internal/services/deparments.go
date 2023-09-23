package services

import (
	"context"
	"educational_program_creator/internal/models"
)

func (s *Service) GetAllDepartments(ctx context.Context) ([]models.Department, error) {
	departments, err := s.repo.GetAllDepartments(ctx)
	if err != nil {
		return nil, err
	}
	return departments, nil
}

func (s *Service) GetDepartmentByID(ctx context.Context, id int) (*models.Department, error) {
	department, err := s.repo.GetDepartmentByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return department, nil
}

func (s *Service) CreateDepartment(ctx context.Context, department *models.Department) error {
	return s.repo.CreateDepartment(ctx, department)
}

func (s *Service) UpdateDepartment(ctx context.Context, department *models.Department) error {
	return s.repo.UpdateDepartment(ctx, department)
}

func (s *Service) DeleteDepartment(ctx context.Context, id int) error {
	return s.repo.DeleteDepartment(ctx, id)
}
