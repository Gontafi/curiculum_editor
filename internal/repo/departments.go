package repo

import (
	"context"
	"educational_program_creator/internal/models"
	"encoding/json"
	"fmt"
	"time"
)

func (r *Repository) GetAllDepartments(ctx context.Context) ([]models.Department, error) {
	departments, err := r.getCachedDepartments(ctx)
	if err == nil && len(departments) > 0 {
		return departments, nil
	}

	var departmentsFromDB []models.Department
	err = r.db.Find(&departmentsFromDB).Error
	if err != nil {
		return nil, err
	}

	err = r.cacheDepartments(ctx, departmentsFromDB)
	if err != nil {
		fmt.Println("Failed to cache departments in Redis:", err)
	}

	return departmentsFromDB, nil
}

func (r *Repository) GetDepartmentByID(ctx context.Context, id int) (*models.Department, error) {
	department, err := r.getCachedDepartment(ctx, id)
	if err == nil {
		return department, nil
	}

	var departmentFromDB models.Department
	err = r.db.First(&departmentFromDB, id).Error
	if err != nil {
		return nil, err
	}

	err = r.cacheDepartment(ctx, &departmentFromDB)
	if err != nil {
		fmt.Println("Failed to cache department in Redis:", err)
	}

	return &departmentFromDB, nil
}

func (r *Repository) getCachedDepartments(ctx context.Context) ([]models.Department, error) {
	key := "department:" + "all"
	result, err := r.rd.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	var departments []models.Department
	err = json.Unmarshal([]byte(result), &departments)
	if err != nil {
		return nil, err
	}

	return departments, nil
}

func (r *Repository) cacheDepartments(ctx context.Context, departments []models.Department) error {
	key := "department:" + "all"
	data, err := json.Marshal(departments)
	if err != nil {
		return err
	}

	return r.rd.Set(ctx, key, data, time.Hour).Err()
}

func (r *Repository) getCachedDepartment(ctx context.Context, id int) (*models.Department, error) {
	key := "department:" + fmt.Sprintf("%d", id)
	result, err := r.rd.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	var department models.Department
	err = json.Unmarshal([]byte(result), &department)
	if err != nil {
		return nil, err
	}

	return &department, nil
}

func (r *Repository) cacheDepartment(ctx context.Context, department *models.Department) error {
	key := "department:" + fmt.Sprintf("%d", department.ID)
	data, err := json.Marshal(department)
	if err != nil {
		return err
	}

	return r.rd.Set(ctx, key, data, time.Hour).Err()
}

func (r *Repository) CreateDepartment(ctx context.Context, department *models.Department) error {
	if err := r.db.Create(department).Error; err != nil {
		return err
	}

	r.clearCachedDepartments(ctx)
	return nil
}

func (r *Repository) UpdateDepartment(ctx context.Context, department *models.Department) error {
	if err := r.db.Save(department).Error; err != nil {
		return err
	}

	r.clearCachedDepartment(ctx, department.ID)
	return nil
}

func (r *Repository) DeleteDepartment(ctx context.Context, id int) error {
	if err := r.db.Delete(&models.Department{}, id).Error; err != nil {
		return err
	}

	r.clearCachedDepartments(ctx)
	r.clearCachedDepartment(ctx, id)
	return nil
}

func (r *Repository) clearCachedDepartments(ctx context.Context) {
	key := "department:" + "all"
	r.rd.Del(ctx, key)
}

func (r *Repository) clearCachedDepartment(ctx context.Context, id int) {
	key := "department:" + fmt.Sprintf("%d", id)
	r.rd.Del(ctx, key)
}
