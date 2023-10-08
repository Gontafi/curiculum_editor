package repo

import (
	"educational_program_creator/internal/models"
)

func (r *Repository) GetAllModules(limit int, offset int) ([]models.Module, error) {
	var modulesFromDB []models.Module
	err := r.db.Limit(limit).Offset(offset).Find(&modulesFromDB).Error
	if err != nil {
		return nil, err
	}

	return modulesFromDB, nil
}

func (r *Repository) GetModuleByID(id int) (*models.Module, error) {
	var moduleFromDB models.Module
	err := r.db.First(&moduleFromDB, id).Error
	if err != nil {
		return nil, err
	}

	return &moduleFromDB, nil
}

func (r *Repository) CreateModule(module *models.Module) (int, error) {
	if err := r.db.Create(&module).Error; err != nil {
		return 0, err
	}

	return module.ID, nil
}

func (r *Repository) UpdateModule(module *models.Module) error {
	if err := r.db.Save(module).Error; err != nil {
		return err
	}

	return nil
}

func (r *Repository) DeleteModule(id int) error {
	if err := r.db.Delete(&models.Module{}, id).Error; err != nil {
		return err
	}

	return nil
}

// You can remove these clearCached functions as they are specific to Redis caching
// func (r *Repository) clearCachedModules(ctx context.Context) {
//     // Remove Redis-specific cache clearing logic
// }

// func (r *Repository) clearCachedModule(id int) {
//     // Remove Redis-specific cache clearing logic
// }
