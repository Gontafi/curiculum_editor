package services

import "educational_program_creator/internal/models"

func (s *Service) CreateCoursePrerequisite(coursePrerequisite *models.CoursePrerequisite) error {
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
