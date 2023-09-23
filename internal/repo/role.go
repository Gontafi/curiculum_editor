package repo

import (
	"educational_program_creator/internal/models"
)

func (r *Repository) AddRole(role models.Role) (int, error) {
	if err := r.db.Create(&role).Error; err != nil {
		return 0, err
	}
	return int(role.ID), nil
}

func (r *Repository) DeleteRole(id int) error {
	if err := r.db.Delete(&models.Role{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (r *Repository) UpdateRole(role models.Role) error {
	if err := r.db.Save(&role).Error; err != nil {
		return err
	}
	return nil
}

func (r *Repository) AllRoles() ([]models.Role, error) {
	var roles []models.Role
	if err := r.db.Find(&roles).Error; err != nil {
		return []models.Role{}, err
	}
	return roles, nil
}

func (r *Repository) GetRoleByID(id int) (models.Role, error) {
	var role models.Role
	if err := r.db.First(&role, id).Error; err != nil {
		return models.Role{}, err
	}
	return role, nil
}
