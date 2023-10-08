package services

import (
	"educational_program_creator/internal/models"
	"educational_program_creator/pkg/utils"
)

func (s *Service) GetAllCycles(page int, perPage int) ([]models.Cycle, error) {
	limit, offset := utils.CalculateLimitOffset(page, perPage)
	cycles, err := s.repo.GetAllCycles(limit, offset)
	if err != nil {
		return nil, err
	}
	return cycles, nil
}

func (s *Service) GetCycleByID(id int) (*models.Cycle, error) {
	cycle, err := s.repo.GetCycleByID(id)
	if err != nil {
		return nil, err
	}
	return cycle, nil
}

func (s *Service) CreateCycle(cycle *models.Cycle) (int, error) {
	return s.repo.CreateCycle(cycle)
}

func (s *Service) UpdateCycle(cycle *models.Cycle) error {
	return s.repo.UpdateCycle(cycle)
}

func (s *Service) DeleteCycle(id int) error {
	return s.repo.DeleteCycle(id)
}
