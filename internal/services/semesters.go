package services

import "educational_program_creator/internal/models"

func (s *Service) GetAllSemesters() ([]models.Semester, error) {
	semesters, err := s.repo.GetAllSemesters()
	if err != nil {
		return nil, err
	}
	return semesters, nil
}

func (s *Service) GetSemesterByID(id int) (*models.Semester, error) {
	semester, err := s.repo.GetSemesterByID(id)
	if err != nil {
		return nil, err
	}
	return semester, nil
}

func (s *Service) CreateSemester(semester *models.Semester) error {
	return s.repo.CreateSemester(semester)
}

func (s *Service) UpdateSemester(semester *models.Semester) error {
	return s.repo.UpdateSemester(semester)
}

func (s *Service) DeleteSemester(id int) error {
	return s.repo.DeleteSemester(id)
}
