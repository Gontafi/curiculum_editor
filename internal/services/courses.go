package services

import (
	"educational_program_creator/internal/models"
	"educational_program_creator/pkg/utils"
)

func (s *Service) GetAllCourses(page int, perPage int) ([]models.Course, error) {
	limit, offset := utils.CalculateLimitOffset(page, perPage)
	courses, err := s.repo.GetAllCourses(limit, offset)
	if err != nil {
		return nil, err
	}
	return courses, nil
}

func (s *Service) GetCourseByID(id int) (*models.Course, error) {
	course, err := s.repo.GetCourseByID(id)
	if err != nil {
		return nil, err
	}
	return course, nil
}

func (s *Service) CreateCourse(course *models.Course) (int, error) {
	return s.repo.CreateCourse(course)
}

func (s *Service) UpdateCourse(course *models.Course) error {
	return s.repo.UpdateCourse(course)
}

func (s *Service) DeleteCourse(id int) error {
	return s.repo.DeleteCourse(id)
}
