package services

import (
	"educational_program_creator/internal/models"
	"educational_program_creator/pkg/utils"
)

func (s *Service) GetAllTotalCourseSemesters(page int, perPage int) ([]models.TotalCourseSemester, error) {
	limit, offset := utils.CalculateLimitOffset(page, perPage)
	totalCourseSemesters, err := s.repo.GetAllTotalCourseSemesters(limit, offset)
	if err != nil {
		return nil, err
	}
	return totalCourseSemesters, nil
}

func (s *Service) GetTotalCourseSemesterByID(id int) (*models.TotalCourseSemester, error) {
	totalCourseSemester, err := s.repo.GetTotalCourseSemesterByID(id)
	if err != nil {
		return nil, err
	}
	return totalCourseSemester, nil
}

func (s *Service) CreateTotalCourseSemester(totalCourseSemester *models.TotalCourseSemester) error {
	return s.repo.CreateTotalCourseSemester(totalCourseSemester)
}

func (s *Service) UpdateTotalCourseSemester(totalCourseSemester *models.TotalCourseSemester) error {
	return s.repo.UpdateTotalCourseSemester(totalCourseSemester)
}

func (s *Service) DeleteTotalCourseSemester(id int) error {
	return s.repo.DeleteTotalCourseSemester(id)
}
