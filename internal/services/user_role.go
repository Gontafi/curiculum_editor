package services

import "educational_program_creator/internal/models"

func (s *Service) GetUserRole(userID int) (*models.Role, error) {
	return s.repo.GetUserRole(userID)
}
