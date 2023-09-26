package services

import (
	"context"
	"educational_program_creator/internal/models"
	"educational_program_creator/pkg/utils"
	"errors"
)

func (s *Service) GetAllComponents(ctx context.Context, page int, perPage int) ([]models.Component, error) {
	limit, offset := utils.CalculateLimitOffset(page, perPage)
	return s.repo.GetAllComponents(ctx, limit, offset)
}

func (s *Service) GetComponentByID(ctx context.Context, id int) (*models.Component, error) {
	if id <= 0 {
		return nil, errors.New("invalid component ID")
	}
	return s.repo.GetComponentByID(ctx, id)
}

func (s *Service) CreateComponent(ctx context.Context, component *models.Component) (int, error) {
	if component == nil {
		return 0, errors.New("nil component provided")
	}
	return s.repo.CreateComponent(ctx, component)
}

func (s *Service) UpdateComponent(ctx context.Context, component *models.Component) error {
	if component == nil || component.ID <= 0 {
		return errors.New("invalid component data")
	}
	return s.repo.UpdateComponent(ctx, component)
}

func (s *Service) DeleteComponent(ctx context.Context, id int) error {
	if id <= 0 {
		return errors.New("invalid component ID")
	}
	return s.repo.DeleteComponent(ctx, id)
}
