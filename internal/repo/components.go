package repo

import (
	"educational_program_creator/internal/models"
)

func (r *Repository) GetAllComponents(limit int, offset int) ([]models.Component, error) {
	var componentsFromDB []models.Component
	err := r.db.Preload("ProfessionalComponent").Limit(limit).Offset(offset).Find(&componentsFromDB).Error
	if err != nil {
		return nil, err
	}

	return componentsFromDB, nil
}

func (r *Repository) GetComponentByID(id int) (*models.Component, error) {
	var componentFromDB models.Component
	err := r.db.First(&componentFromDB, id).Error
	if err != nil {
		return nil, err
	}

	return &componentFromDB, nil
}

func (r *Repository) CreateComponent(component *models.Component) (int, error) {
	if err := r.db.Create(&component).Error; err != nil {
		return 0, err
	}

	return component.ID, nil
}

func (r *Repository) UpdateComponent(component *models.Component) error {
	if err := r.db.Save(component).Error; err != nil {
		return err
	}

	return nil
}

func (r *Repository) DeleteComponent(id int) error {
	if err := r.db.Delete(&models.Component{}, id).Error; err != nil {
		return err
	}

	return nil
}
