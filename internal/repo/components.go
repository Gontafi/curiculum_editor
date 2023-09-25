package repo

import (
	"context"
	"educational_program_creator/internal/models"
	"encoding/json"
	"fmt"
	"time"
)

func (r *Repository) GetAllComponents(ctx context.Context, limit int, offset int) ([]models.Component, error) {
	components, err := r.getCachedComponents(ctx)
	if err == nil && len(components) > 0 {
		return components, nil
	}

	var componentsFromDB []models.Component
	err = r.db.Find(&componentsFromDB).Limit(limit).Offset(offset).Error
	if err != nil {
		return nil, err
	}

	err = r.cacheComponents(ctx, componentsFromDB)
	if err != nil {
		fmt.Println("Failed to cache components in Redis:", err)
	}

	return componentsFromDB, nil
}

func (r *Repository) GetComponentByID(ctx context.Context, id int) (*models.Component, error) {
	component, err := r.getCachedComponent(ctx, id)
	if err == nil {
		return component, nil
	}

	var componentFromDB models.Component
	err = r.db.First(&componentFromDB, id).Error
	if err != nil {
		return nil, err
	}

	err = r.cacheComponent(ctx, &componentFromDB)
	if err != nil {
		fmt.Println("Failed to cache component in Redis:", err)
	}

	return &componentFromDB, nil
}

func (r *Repository) getCachedComponents(ctx context.Context) ([]models.Component, error) {
	key := "component:" + "all"
	result, err := r.rd.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	var components []models.Component
	err = json.Unmarshal([]byte(result), &components)
	if err != nil {
		return nil, err
	}

	return components, nil
}

func (r *Repository) cacheComponents(ctx context.Context, components []models.Component) error {
	key := "component:" + "all"
	data, err := json.Marshal(components)
	if err != nil {
		return err
	}

	return r.rd.Set(ctx, key, data, time.Hour).Err()
}

func (r *Repository) getCachedComponent(ctx context.Context, id int) (*models.Component, error) {
	key := "component:" + fmt.Sprintf("%d", id)
	result, err := r.rd.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	var component models.Component
	err = json.Unmarshal([]byte(result), &component)
	if err != nil {
		return nil, err
	}

	return &component, nil
}

func (r *Repository) cacheComponent(ctx context.Context, component *models.Component) error {
	key := "component:" + fmt.Sprintf("%d", component.ID)
	data, err := json.Marshal(component)
	if err != nil {
		return err
	}

	return r.rd.Set(ctx, key, data, time.Hour).Err()
}

func (r *Repository) CreateComponent(ctx context.Context, component *models.Component) error {
	if err := r.db.Create(component).Error; err != nil {
		return err
	}

	r.clearCachedComponents(ctx)
	return nil
}

func (r *Repository) UpdateComponent(ctx context.Context, component *models.Component) error {
	if err := r.db.Save(component).Error; err != nil {
		return err
	}

	r.clearCachedComponent(ctx, component.ID)
	return nil
}

func (r *Repository) DeleteComponent(ctx context.Context, id int) error {
	if err := r.db.Delete(&models.Component{}, id).Error; err != nil {
		return err
	}

	r.clearCachedComponents(ctx)
	r.clearCachedComponent(ctx, id)
	return nil
}

func (r *Repository) clearCachedComponents(ctx context.Context) {
	key := "component:" + "all"
	r.rd.Del(ctx, key)
}

func (r *Repository) clearCachedComponent(ctx context.Context, id int) {
	key := "component:" + fmt.Sprintf("%d", id)
	r.rd.Del(ctx, key)
}
