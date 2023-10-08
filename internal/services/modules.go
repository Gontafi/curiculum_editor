package services

import (
	"educational_program_creator/internal/models"
	"educational_program_creator/pkg/utils"
)

func (s *Service) GetAllModules(page int, perPage int) ([]models.Module, error) {
	limit, offset := utils.CalculateLimitOffset(page, perPage)
	modules, err := s.repo.GetAllModules(limit, offset)
	if err != nil {
		return nil, err
	}
	return modules, nil
}

func (s *Service) GetModuleByID(id int) (*models.Module, error) {
	module, err := s.repo.GetModuleByID(id)
	if err != nil {
		return nil, err
	}
	return module, nil
}

func (s *Service) CreateModule(module *models.Module) (int, error) {
	return s.repo.CreateModule(module)
}

func (s *Service) UpdateModule(module *models.Module) error {
	return s.repo.UpdateModule(module)
}

func (s *Service) DeleteModule(id int) error {
	return s.repo.DeleteModule(id)
}
