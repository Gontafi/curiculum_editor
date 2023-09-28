package repo

import (
	"context"
	"educational_program_creator/internal/models"
	"encoding/json"
	"fmt"
	"time"
)

func (r *Repository) GetAllModules(ctx context.Context, limit int, offset int) ([]models.Module, error) {
	modules, err := r.getCachedModules(ctx)
	if err == nil && len(modules) > 0 {
		return modules, nil
	}

	var modulesFromDB []models.Module
	err = r.db.Preload("Cycle").Limit(limit).Offset(offset).Find(&modulesFromDB).Error
	if err != nil {
		return nil, err
	}

	err = r.cacheModules(ctx, modulesFromDB)
	if err != nil {
		fmt.Println("Failed to cache modules in Redis:", err)
	}

	return modulesFromDB, nil
}

func (r *Repository) GetModuleByID(ctx context.Context, id int) (*models.Module, error) {
	module, err := r.getCachedModule(ctx, id)
	if err == nil {
		return module, nil
	}

	var moduleFromDB models.Module
	err = r.db.First(&moduleFromDB, id).Error
	if err != nil {
		return nil, err
	}

	err = r.cacheModule(ctx, &moduleFromDB)
	if err != nil {
		fmt.Println("Failed to cache module in Redis:", err)
	}

	return &moduleFromDB, nil
}

func (r *Repository) getCachedModules(ctx context.Context) ([]models.Module, error) {
	key := "module:" + "all"
	result, err := r.rd.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	var modules []models.Module
	err = json.Unmarshal([]byte(result), &modules)
	if err != nil {
		return nil, err
	}

	return modules, nil
}

func (r *Repository) cacheModules(ctx context.Context, modules []models.Module) error {
	key := "module:" + "all"
	data, err := json.Marshal(modules)
	if err != nil {
		return err
	}

	return r.rd.Set(ctx, key, data, time.Hour).Err()
}

func (r *Repository) getCachedModule(ctx context.Context, id int) (*models.Module, error) {
	key := "module:" + fmt.Sprintf("%d", id)
	result, err := r.rd.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	var module models.Module
	err = json.Unmarshal([]byte(result), &module)
	if err != nil {
		return nil, err
	}

	return &module, nil
}

func (r *Repository) cacheModule(ctx context.Context, module *models.Module) error {
	key := "module:" + fmt.Sprintf("%d", module.ID)
	data, err := json.Marshal(module)
	if err != nil {
		return err
	}

	return r.rd.Set(ctx, key, data, time.Hour).Err()
}

func (r *Repository) CreateModule(ctx context.Context, module *models.Module) (int, error) {
	if err := r.db.Create(&module).Error; err != nil {
		return 0, err
	}

	r.clearCachedModules(ctx)
	return module.ID, nil
}

func (r *Repository) UpdateModule(ctx context.Context, module *models.Module) error {
	if err := r.db.Save(module).Error; err != nil {
		return err
	}

	r.clearCachedModule(ctx, module.ID)
	return nil
}

func (r *Repository) DeleteModule(ctx context.Context, id int) error {
	if err := r.db.Delete(&models.Module{}, id).Error; err != nil {
		return err
	}

	r.clearCachedModules(ctx)
	r.clearCachedModule(ctx, id)
	return nil
}

func (r *Repository) clearCachedModules(ctx context.Context) {
	key := "module:" + "all"
	r.rd.Del(ctx, key)
}

func (r *Repository) clearCachedModule(ctx context.Context, id int) {
	key := "module:" + fmt.Sprintf("%d", id)
	r.rd.Del(ctx, key)
}
