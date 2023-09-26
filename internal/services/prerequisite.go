package services

import (
	"educational_program_creator/internal/models"
	"educational_program_creator/pkg/utils"
)

func (s *Service) GetAllCoursePrerequisite(page int, perPage int) ([]models.CoursePrerequisite, error) {
	limit, offset := utils.CalculateLimitOffset(page, perPage)
	return s.repo.GetAllCoursePrerequisite(limit, offset)
}
func (s *Service) CreateCoursePrerequisite(coursePrerequisite *models.CoursePrerequisite) (int, error) {
	return s.repo.CreateCoursePrerequisite(coursePrerequisite)
}

func (s *Service) GetCoursePrerequisiteByID(id int) (*models.CoursePrerequisite, error) {
	return s.repo.GetCoursePrerequisiteByID(id)
}

func (s *Service) UpdateCoursePrerequisite(coursePrerequisite *models.CoursePrerequisite) error {
	return s.repo.UpdateCoursePrerequisite(coursePrerequisite)
}

func (s *Service) DeleteCoursePrerequisite(id int) error {
	return s.repo.DeleteCoursePrerequisite(id)
}
