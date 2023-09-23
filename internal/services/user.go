package services

import (
	"educational_program_creator/internal/models"
)

func (s *Service) AddUser(user models.User) (int, error) {
	return s.repo.AddUser(user)
}

func (s *Service) UpdateUser(user models.User) error {
	return s.repo.UpdateUser(user)
}

func (s *Service) GetAllUsers() ([]models.User, error) {
	return s.repo.GetAllUsers()
}

func (s *Service) GetUserByID(id int) (*models.User, error) {
	return s.repo.GetUserById(id)
}

func (s *Service) GetUserByUsername(username string) (*models.User, error) {
	return s.repo.GetUserByUsername(username)
}
