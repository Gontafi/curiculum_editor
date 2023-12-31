package services

import (
	"educational_program_creator/internal/repo"
)

type Service struct {
	repo repo.RepositoryInterface
}

func NewService(repo repo.RepositoryInterface) *Service {
	return &Service{repo: repo}
}
