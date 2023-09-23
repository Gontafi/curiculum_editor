package repo

import (
	"context"
	"educational_program_creator/internal/models"
	"encoding/json"
	"fmt"
	"time"
)

func (r *Repository) GetAllCycles(ctx context.Context) ([]models.Cycle, error) {
	cycles, err := r.getCachedCycles(ctx)
	if err == nil && len(cycles) > 0 {
		return cycles, nil
	}

	var cyclesFromDB []models.Cycle
	err = r.db.Find(&cyclesFromDB).Error
	if err != nil {
		return nil, err
	}

	err = r.cacheCycles(ctx, cyclesFromDB)
	if err != nil {
		fmt.Println("Failed to cache cycles in Redis:", err)
	}

	return cyclesFromDB, nil
}

func (r *Repository) GetCycleByID(ctx context.Context, id int) (*models.Cycle, error) {
	cycle, err := r.getCachedCycle(ctx, id)
	if err == nil {
		return cycle, nil
	}

	var cycleFromDB models.Cycle
	err = r.db.First(&cycleFromDB, id).Error
	if err != nil {
		return nil, err
	}

	err = r.cacheCycle(ctx, &cycleFromDB)
	if err != nil {
		fmt.Println("Failed to cache cycle in Redis:", err)
	}

	return &cycleFromDB, nil
}

func (r *Repository) getCachedCycles(ctx context.Context) ([]models.Cycle, error) {
	key := "cycle:" + "all"
	result, err := r.rd.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	var cycles []models.Cycle
	err = json.Unmarshal([]byte(result), &cycles)
	if err != nil {
		return nil, err
	}

	return cycles, nil
}

func (r *Repository) cacheCycles(ctx context.Context, cycles []models.Cycle) error {
	key := "cycle:" + "all"
	data, err := json.Marshal(cycles)
	if err != nil {
		return err
	}

	return r.rd.Set(ctx, key, data, time.Hour).Err()
}

func (r *Repository) getCachedCycle(ctx context.Context, id int) (*models.Cycle, error) {
	key := "cycle:" + fmt.Sprintf("%d", id)
	result, err := r.rd.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	var cycle models.Cycle
	err = json.Unmarshal([]byte(result), &cycle)
	if err != nil {
		return nil, err
	}

	return &cycle, nil
}

func (r *Repository) cacheCycle(ctx context.Context, cycle *models.Cycle) error {
	key := "cycle:" + fmt.Sprintf("%d", cycle.ID)
	data, err := json.Marshal(cycle)
	if err != nil {
		return err
	}

	return r.rd.Set(ctx, key, data, time.Hour).Err()
}

func (r *Repository) CreateCycle(ctx context.Context, cycle *models.Cycle) error {
	if err := r.db.Create(cycle).Error; err != nil {
		return err
	}

	r.clearCachedCycles(ctx)
	return nil
}

func (r *Repository) UpdateCycle(ctx context.Context, cycle *models.Cycle) error {
	if err := r.db.Save(cycle).Error; err != nil {
		return err
	}

	r.clearCachedCycle(ctx, cycle.ID)
	return nil
}

func (r *Repository) DeleteCycle(ctx context.Context, id int) error {
	if err := r.db.Delete(&models.Cycle{}, id).Error; err != nil {
		return err
	}

	r.clearCachedCycles(ctx)
	r.clearCachedCycle(ctx, id)
	return nil
}

func (r *Repository) clearCachedCycles(ctx context.Context) {
	key := "cycle:" + "all"
	r.rd.Del(ctx, key)
}

func (r *Repository) clearCachedCycle(ctx context.Context, id int) {
	key := "cycle:" + fmt.Sprintf("%d", id)
	r.rd.Del(ctx, key)
}
