package repo

import "educational_program_creator/internal/models"

func (r *Repository) GetAllModules() ([]models.Module, error) {
	var modules []models.Module
	err := r.db.Find(&modules).Error
	if err != nil {
		return nil, err
	}
	return modules, nil
}

func (r *Repository) GetModuleByID(id int) (*models.Module, error) {
	var module models.Module
	err := r.db.First(&module, id).Error
	if err != nil {
		return nil, err
	}
	return &module, nil
}

func (r *Repository) CreateModule(module *models.Module) error {
	return r.db.Create(module).Error
}

func (r *Repository) UpdateModule(module *models.Module) error {
	return r.db.Save(module).Error
}

func (r *Repository) DeleteModule(id int) error {
	return r.db.Delete(&models.Module{}, id).Error
}
