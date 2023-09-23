package handlers

import (
	"educational_program_creator/internal/handlers/api"
	"educational_program_creator/internal/handlers/middleware"
	"educational_program_creator/internal/handlers/rest"
)

type Handler struct {
	crud       *rest.CrudHandler
	api        *api.UserHandler
	middleware *middleware.Middleware
}

func NewHandler(crud *rest.CrudHandler) *Handler {
	return &Handler{crud: crud}
}
