package api

import services "educational_program_creator/internal/services"

type UserHandler struct {
	Service *services.Service
}

func NewUserHandler(service *services.Service) *UserHandler {
	return &UserHandler{Service: service}
}
