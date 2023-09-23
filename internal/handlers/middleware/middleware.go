package middleware

import (
	"educational_program_creator/internal/services"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"strings"
)

const (
	DefaultRoleId = 2
	AdminRoleId   = 1
)

type Middleware struct {
	Service *services.Service
}

func NewMiddleware(service *services.Service) *Middleware {
	return &Middleware{Service: service}
}

func (m *Middleware) AdminRoleMiddleware() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		tokenString := ctx.Get("Authorization")
		if tokenString == "" {
			return ctx.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
		}

		headerParts := strings.Split(tokenString, " ")
		if len(headerParts) != 2 || headerParts[0] != "Bearer" {
			return ctx.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token"})
		}

		claims, err := m.Service.ParseToken(headerParts[1])
		if err != nil {
			return ctx.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token"})
		}

		role, err := m.Service.GetUserRole(claims.UserID)
		if err != nil {
			return ctx.Status(http.StatusForbidden).JSON(fiber.Map{"error": "Failed to get role"})
		}

		if role.ID != AdminRoleId {
			return ctx.Status(http.StatusForbidden).JSON(fiber.Map{"error": "Access denied"})
		}

		return ctx.Next()
	}
}
