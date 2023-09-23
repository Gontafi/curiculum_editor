package services

import "educational_program_creator/internal/models"

func (s *Service) GetAllProfessionalComponents() ([]models.ProfessionalComponent, error) {
	components, err := s.repo.GetAllProfessionalComponents()
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

func (s *Service) CreateProfessionalComponent(component *models.ProfessionalComponent) error {
	return s.repo.CreateProfessionalComponent(component)
}

func (s *Service) UpdateProfessionalComponent(component *models.ProfessionalComponent) error {
	return s.repo.UpdateProfessionalComponent(component)
}

func (s *Service) DeleteProfessionalComponent(id int) error {
	return s.repo.DeleteProfessionalComponent(id)
}
