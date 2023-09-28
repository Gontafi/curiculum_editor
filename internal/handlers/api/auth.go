package api

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"net/http"
)

type SignInForm struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (h *UserHandler) SignUp(c *fiber.Ctx) error {

	var user SignInForm
	err := c.BodyParser(&user)
	if err != nil {
		log.Println()
		log.Println(err)
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	userID, err := h.Service.RegisterUser(user.Username, user.Password)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Error registering user"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"user_id": userID})
}

func (h *UserHandler) SignIn(c *fiber.Ctx) error {
	var form SignInForm

	err := c.BodyParser(&form)
	if err != nil {
		log.Println(err)
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	token, err := h.Service.GetTokenFromUser(form.Username, form.Password)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get Token check username and password",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"access_token": token})
}
