package repo

import (
	"educational_program_creator/internal/models"
)

func (r *Repository) GetAllCycles() ([]models.Cycle, error) {
	var cycles []models.Cycle
	err := r.db.Find(&cycles).Error
	if err != nil {
		return nil, err
	}
	return cycles, nil
}

func (r *Repository) GetCycleByID(id int) (*models.Cycle, error) {
	var cycle models.Cycle
	err := r.db.First(&cycle, id).Error
	if err != nil {
		return nil, err
	}
	return &cycle, nil
}

func (r *Repository) CreateCycle(cycle *models.Cycle) error {
	return r.db.Create(cycle).Error
}

func (r *Repository) UpdateCycle(cycle *models.Cycle) error {
	return r.db.Save(cycle).Error
}

func (r *Repository) DeleteCycle(id int) error {
	return r.db.Delete(&models.Cycle{}, id).Error
}
