package auth

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type UserPasswordLogin struct{}

func (u *UserPasswordLogin) Do(ctx *fiber.Ctx) error {
	var credentials map[string]string

	if err := ctx.BodyParser(&credentials); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
	}

	return nil
}
