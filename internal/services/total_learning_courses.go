package services

import "educational_program_creator/internal/models"

func (s *Service) GetTotalLearningCourseByID(id int) (*models.TotalLearningCourse, error) {
	totalLearningCourse, err := s.repo.GetTotalLearningCourseByID(id)
	if err != nil {
		return nil, err
	}
	return totalLearningCourse, nil
}

func (s *Service) CreateTotalLearningCourse(totalLearningCourse *models.TotalLearningCourse) error {
	return s.repo.CreateTotalLearningCourse(totalLearningCourse)
}

func (s *Service) UpdateTotalLearningCourse(totalLearningCourse *models.TotalLearningCourse) error {
	return s.repo.UpdateTotalLearningCourse(totalLearningCourse)
}

func (s *Service) DeleteTotalLearningCourse(id int) error {
	return s.repo.DeleteTotalLearningCourse(id)
}
