package repo

import (
	"context"
	"educational_program_creator/internal/models"
	"encoding/json"
	"fmt"
	"time"
)

func (r *Repository) GetAllCourses(ctx context.Context, limit int, offset int) ([]models.Course, error) {
	courses, err := r.getCachedCourses(ctx)
	if err == nil && len(courses) > 0 {
		return courses, nil
	}

	var coursesFromDB []models.Course
	err = r.db.Find(&coursesFromDB).Limit(limit).Offset(offset).Error
	if err != nil {
		return nil, err
	}

	err = r.cacheCourses(ctx, coursesFromDB)
	if err != nil {
		fmt.Println("Failed to cache courses in Redis:", err)
	}

	return coursesFromDB, nil
}

func (r *Repository) GetCourseByID(ctx context.Context, id int) (*models.Course, error) {
	course, err := r.getCachedCourse(ctx, id)
	if err == nil {
		return course, nil
	}

	var courseFromDB models.Course
	err = r.db.First(&courseFromDB, id).Error
	if err != nil {
		return nil, err
	}

	err = r.cacheCourse(ctx, &courseFromDB)
	if err != nil {
		fmt.Println("Failed to cache course in Redis:", err)
	}

	return &courseFromDB, nil
}

func (r *Repository) getCachedCourses(ctx context.Context) ([]models.Course, error) {
	key := "course:" + "all"
	result, err := r.rd.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	var courses []models.Course
	err = json.Unmarshal([]byte(result), &courses)
	if err != nil {
		return nil, err
	}

	return courses, nil
}

func (r *Repository) cacheCourses(ctx context.Context, courses []models.Course) error {
	key := "course:" + "all"
	data, err := json.Marshal(courses)
	if err != nil {
		return err
	}

	return r.rd.Set(ctx, key, data, time.Hour).Err()
}

func (r *Repository) getCachedCourse(ctx context.Context, id int) (*models.Course, error) {
	key := "course:" + fmt.Sprintf("%d", id)
	result, err := r.rd.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	var course models.Course
	err = json.Unmarshal([]byte(result), &course)
	if err != nil {
		return nil, err
	}

	return &course, nil
}

func (r *Repository) cacheCourse(ctx context.Context, course *models.Course) error {
	key := "course:" + fmt.Sprintf("%d", course.ID)
	data, err := json.Marshal(course)
	if err != nil {
		return err
	}

	return r.rd.Set(ctx, key, data, time.Hour).Err()
}

func (r *Repository) CreateCourse(ctx context.Context, course *models.Course) error {
	if err := r.db.Create(course).Error; err != nil {
		return err
	}

	r.clearCachedCourses(ctx)
	return nil
}

func (r *Repository) UpdateCourse(ctx context.Context, course *models.Course) error {
	if err := r.db.Save(course).Error; err != nil {
		return err
	}

	r.clearCachedCourse(ctx, course.ID)
	return nil
}

func (r *Repository) DeleteCourse(ctx context.Context, id int) error {
	if err := r.db.Delete(&models.Course{}, id).Error; err != nil {
		return err
	}

	r.clearCachedCourses(ctx)
	r.clearCachedCourse(ctx, id)
	return nil
}

func (r *Repository) clearCachedCourses(ctx context.Context) {
	key := "course:" + "all"
	r.rd.Del(ctx, key)
}

func (r *Repository) clearCachedCourse(ctx context.Context, id int) {
	key := "course:" + fmt.Sprintf("%d", id)
	r.rd.Del(ctx, key)
}
