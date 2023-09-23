package handlers

import (
	"educational_program_creator/internal/handlers/api"
	"educational_program_creator/internal/handlers/middleware"
	"educational_program_creator/internal/handlers/restapi"
)

type Handler struct {
	CRUD       *restapi.CrudHandler
	API        *api.UserHandler
	Middleware *middleware.Middleware
}

func NewHandler(crud *restapi.CrudHandler, api *api.UserHandler, middleware *middleware.Middleware) *Handler {
	return &Handler{CRUD: crud, API: api, Middleware: middleware}
}
