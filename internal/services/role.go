package services

import (
	"educational_program_creator/internal/models"
	"educational_program_creator/internal/repo"
)

type RoleService struct {
	repository *repo.Repository
}

func NewRoleService(repository *repo.Repository) *RoleService {
	return &RoleService{
		repository: repository,
	}
}

func (s *RoleService) AddRole(role models.Role) (int, error) {
	return s.repository.AddRole(role)
}

func (s *RoleService) DeleteRole(id int) error {
	return s.repository.DeleteRole(id)
}

func (s *RoleService) UpdateRole(role models.Role) error {
	return s.repository.UpdateRole(role)
}

func (s *RoleService) AllRoles() ([]models.Role, error) {
	return s.repository.AllRoles()
}

func (s *RoleService) GetRoleByID(id int) (models.Role, error) {
	return s.repository.GetRoleByID(id)
}
