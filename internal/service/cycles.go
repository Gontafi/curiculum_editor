package services

import "educational_program_creator/internal/models"

func (s *Service) GetAllCycles() ([]models.Cycle, error) {
	cycles, err := s.repo.GetAllCycles()
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

func (s *Service) CreateCycle(cycle *models.Cycle) error {
	return s.repo.CreateCycle(cycle)
}

func (s *Service) UpdateCycle(cycle *models.Cycle) error {
	return s.repo.UpdateCycle(cycle)
}

func (s *Service) DeleteCycle(id int) error {
	return s.repo.DeleteCycle(id)
}
