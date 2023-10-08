package services

import (
	"educational_program_creator/internal/models"
	"educational_program_creator/pkg/utils"
	"errors"
)

func (s *Service) GetAllComponents(page int, perPage int) ([]models.Component, error) {
	limit, offset := utils.CalculateLimitOffset(page, perPage)
	return s.repo.GetAllComponents(limit, offset)
}

func (s *Service) GetComponentByID(id int) (*models.Component, error) {
	if id <= 0 {
		return nil, errors.New("invalid component ID")
	}
	return s.repo.GetComponentByID(id)
}

func (s *Service) CreateComponent(component *models.Component) (int, error) {
	if component == nil {
		return 0, errors.New("nil component provided")
	}
	return s.repo.CreateComponent(component)
}

func (s *Service) UpdateComponent(component *models.Component) error {
	if component == nil || component.ID <= 0 {
		return errors.New("invalid component data")
	}
	return s.repo.UpdateComponent(component)
}

func (s *Service) DeleteComponent(id int) error {
	if id <= 0 {
		return errors.New("invalid component ID")
	}
	return s.repo.DeleteComponent(id)
}
