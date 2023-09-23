package services

import (
	"context"
	"educational_program_creator/internal/models"
)

func (s *Service) GetAllCycles(ctx context.Context) ([]models.Cycle, error) {
	cycles, err := s.repo.GetAllCycles(ctx)
	if err != nil {
		return nil, err
	}
	return cycles, nil
}

func (s *Service) GetCycleByID(ctx context.Context, id int) (*models.Cycle, error) {
	cycle, err := s.repo.GetCycleByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return cycle, nil
}

func (s *Service) CreateCycle(ctx context.Context, cycle *models.Cycle) error {
	return s.repo.CreateCycle(ctx, cycle)
}

func (s *Service) UpdateCycle(ctx context.Context, cycle *models.Cycle) error {
	return s.repo.UpdateCycle(ctx, cycle)
}

func (s *Service) DeleteCycle(ctx context.Context, id int) error {
	return s.repo.DeleteCycle(ctx, id)
}
