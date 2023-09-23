package restapi

import services "educational_program_creator/internal/services"

type CrudHandler struct {
	Service *services.Service
}

func NewCrudHandler(service *services.Service) *CrudHandler {
	return &CrudHandler{Service: service}
}
