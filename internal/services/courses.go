package services

import (
	"context"
	"educational_program_creator/internal/models"
	"educational_program_creator/pkg/utils"
)

func (s *Service) GetAllCourses(ctx context.Context, page int, perPage int) ([]models.Course, error) {
	limit, offset := utils.CalculateLimitOffset(page, perPage)
	courses, err := s.repo.GetAllCourses(ctx, limit, offset)
	if err != nil {
		return nil, err
	}
	return courses, nil
}

func (s *Service) GetCourseByID(ctx context.Context, id int) (*models.Course, error) {
	course, err := s.repo.GetCourseByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return course, nil
}

func (s *Service) CreateCourse(ctx context.Context, course *models.Course) (int, error) {
	return s.repo.CreateCourse(ctx, course)
}

func (s *Service) UpdateCourse(ctx context.Context, course *models.Course) error {
	return s.repo.UpdateCourse(ctx, course)
}

func (s *Service) DeleteCourse(ctx context.Context, id int) error {
	return s.repo.DeleteCourse(ctx, id)
}
