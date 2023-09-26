package services

import (
	"educational_program_creator/internal/models"
	"educational_program_creator/pkg/utils"
)

func (s *Service) GetAllProfessionalComponents(page int, perPage int) ([]models.ProfessionalComponent, error) {
	limit, offset := utils.CalculateLimitOffset(page, perPage)
	components, err := s.repo.GetAllProfessionalComponents(limit, offset)
	if err != nil {
		return nil, err
	}
	return components, nil
}

func (s *Service) GetProfessionalComponentByID(id int) (*models.ProfessionalComponent, error) {
	component, err := s.repo.GetProfessionalComponentByID(id)
	if err != nil {
		return nil, err
	}
	return component, nil
}

func (s *Service) CreateProfessionalComponent(component *models.ProfessionalComponent) (int, error) {
	return s.repo.CreateProfessionalComponent(component)
}

func (s *Service) UpdateProfessionalComponent(component *models.ProfessionalComponent) error {
	return s.repo.UpdateProfessionalComponent(component)
}

func (s *Service) DeleteProfessionalComponent(id int) error {
	return s.repo.DeleteProfessionalComponent(id)
}
