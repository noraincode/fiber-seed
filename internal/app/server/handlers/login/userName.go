package login

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/google/wire"
	"github.com/noraincode/fiber-seed/internal/app/server/domains/auth"
)

var LoginSet = wire.NewSet(wire.Struct(new(UserLogin), "*"))

type UserLogin struct{}

func (l *UserLogin) UserNameLogin(c *fiber.Ctx) error {
	authOperator := auth.LoginOperator{}
	authOperator.SetOperator(&auth.UserPasswordLogin{})

	err := authOperator.DoLogin(c)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "login successful"})
}
