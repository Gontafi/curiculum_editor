package repo

import (
	"educational_program_creator/internal/models"
)

func (r *Repository) GetAllCycles(limit int, offset int) ([]models.Cycle, error) {
	var cyclesFromDB []models.Cycle
	err := r.db.Preload("Module").Limit(limit).Offset(offset).Find(&cyclesFromDB).Error
	if err != nil {
		return nil, err
	}

	return cyclesFromDB, nil
}

func (r *Repository) GetCycleByID(id int) (*models.Cycle, error) {
	var cycleFromDB models.Cycle
	err := r.db.First(&cycleFromDB, id).Error
	if err != nil {
		return nil, err
	}

	return &cycleFromDB, nil
}

func (r *Repository) CreateCycle(cycle *models.Cycle) (int, error) {
	if err := r.db.Create(&cycle).Error; err != nil {
		return 0, err
	}

	return cycle.ID, nil
}

func (r *Repository) UpdateCycle(cycle *models.Cycle) error {
	if err := r.db.Save(cycle).Error; err != nil {
		return err
	}

	return nil
}

func (r *Repository) DeleteCycle(id int) error {
	if err := r.db.Delete(&models.Cycle{}, id).Error; err != nil {
		return err
	}

	return nil
}
