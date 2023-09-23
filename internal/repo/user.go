package repo

import (
	"educational_program_creator/internal/models"
	"educational_program_creator/pkg/utils"
)

func (r *Repository) AddUser(user models.User) (int, error) {
	result := r.db.Create(&user)
	if result.Error != nil {
		return 0, result.Error
	}

	return user.ID, nil
}

func (r *Repository) UpdateUser(user models.User) error {
	result := r.db.UpdateColumns(&user)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return utils.UserNotFoundErr
	}
	r.db.Save(&user)
	return nil
}

func (r *Repository) GetAllUsers() ([]models.User, error) {
	var users []models.User
	result := r.db.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}

func (r *Repository) GetUserById(id int) (*models.User, error) {
	var user *models.User
	result := r.db.First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func (r *Repository) GetUserByUsername(username string) (*models.User, error) {
	var user *models.User
	result := r.db.Where("username = ?", username).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}
