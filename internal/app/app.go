package app

import (
	"educational_program_creator/internal/config"
	"educational_program_creator/internal/handlers"
	"educational_program_creator/internal/handlers/api"
	"educational_program_creator/internal/handlers/middleware"
	"educational_program_creator/internal/handlers/restapi"
	"educational_program_creator/internal/repo"
	"educational_program_creator/internal/repo/db"
	"educational_program_creator/internal/routes"
	"educational_program_creator/internal/services"
	"github.com/gofiber/fiber/v2"
	"log"
)

func NewServer(cfg *config.AppConfig) {
	gdb, err := db.NewDB(cfg)
	if err != nil {
		log.Fatalln(err)
	}
	redisClient := db.NewRedisClient(cfg)
	repository := repo.NewRepository(gdb, redisClient)
	service := services.NewService(repository)
	handler := handlers.NewHandler(
		restapi.NewCrudHandler(service),
		api.NewUserHandler(service),
		middleware.NewMiddleware(service),
	)
	app := fiber.New(fiber.Config{})
	routes.InitRoutes(handler, app)

}
