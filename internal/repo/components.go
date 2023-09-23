package repo

import "educational_program_creator/internal/models"

func (r *Repository) GetAllComponents() ([]models.Component, error) {
	var components []models.Component
	err := r.db.Find(&components).Error
	if err != nil {
		return nil, err
	}
	return components, nil
}

func (r *Repository) GetComponentByID(id int) (*models.Component, error) {
	var component models.Component
	err := r.db.First(&component, id).Error
	if err != nil {
		return nil, err
	}
	return &component, nil
}

func (r *Repository) CreateComponent(component *models.Component) error {
	return r.db.Create(component).Error
}

func (r *Repository) UpdateComponent(component *models.Component) error {
	return r.db.Save(component).Error
}

func (r *Repository) DeleteComponent(id int) error {
	return r.db.Delete(&models.Component{}, id).Error
}
