package services

import (
	"context"
	"educational_program_creator/internal/models"
)

func (s *Service) GetAllModules(ctx context.Context) ([]models.Module, error) {
	modules, err := s.repo.GetAllModules(ctx)
	if err != nil {
		return nil, err
	}
	return modules, nil
}

func (s *Service) GetModuleByID(ctx context.Context, id int) (*models.Module, error) {
	module, err := s.repo.GetModuleByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return module, nil
}

func (s *Service) CreateModule(ctx context.Context, module *models.Module) error {
	return s.repo.CreateModule(ctx, module)
}

func (s *Service) UpdateModule(ctx context.Context, module *models.Module) error {
	return s.repo.UpdateModule(ctx, module)
}

func (s *Service) DeleteModule(ctx context.Context, id int) error {
	return s.repo.DeleteModule(ctx, id)
}
