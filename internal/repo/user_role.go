package repo

import "educational_program_creator/internal/models"

func (r *Repository) GetUserRole(userID int) (*models.Role, error) {
	var user models.User

	if err := r.db.First(&user, userID).Error; err != nil {
		return nil, err
	}

	return &user.Role, nil
}
