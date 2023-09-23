package repo

import "educational_program_creator/internal/models"

func (r *Repository) GetAllProfessionalComponents() ([]models.ProfessionalComponent, error) {
	var professionalComponents []models.ProfessionalComponent
	err := r.db.Find(&professionalComponents).Error
	if err != nil {
		return nil, err
	}
	return professionalComponents, nil
}

func (r *Repository) GetProfessionalComponentByID(id int) (*models.ProfessionalComponent, error) {
	var professionalComponent models.ProfessionalComponent
	err := r.db.First(&professionalComponent, id).Error
	if err != nil {
		return nil, err
	}
	return &professionalComponent, nil
}

func (r *Repository) CreateProfessionalComponent(profComponent *models.ProfessionalComponent) error {
	return r.db.Create(profComponent).Error
}

func (r *Repository) UpdateProfessionalComponent(profComponent *models.ProfessionalComponent) error {
	return r.db.Save(profComponent).Error
}

func (r *Repository) DeleteProfessionalComponent(id int) error {
	return r.db.Delete(&models.ProfessionalComponent{}, id).Error
}
