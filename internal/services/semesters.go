package services

import (
	"educational_program_creator/internal/models"
	"educational_program_creator/pkg/utils"
)

func (s *Service) GetAllSemesters(page int, perPage int) ([]models.SemesterCourse, error) {
	limit, offset := utils.CalculateLimitOffset(page, perPage)
	semesters, err := s.repo.GetAllSemesters(limit, offset)
	if err != nil {
		return nil, err
	}
	return semesters, nil
}

func (s *Service) GetSemesterByID(id int) (*models.SemesterCourse, error) {
	semester, err := s.repo.GetSemesterByID(id)
	if err != nil {
		return nil, err
	}
	return semester, nil
}

func (s *Service) CreateSemester(semester *models.SemesterCourse) error {
	return s.repo.CreateSemester(semester)
}

func (s *Service) UpdateSemester(semester *models.SemesterCourse) error {
	return s.repo.UpdateSemester(semester)
}

func (s *Service) DeleteSemester(id int) error {
	return s.repo.DeleteSemester(id)
}
