package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/wire"
	"github.com/noraincode/fiber-seed/internal/app/server/handlers/login"
)

var loginRouterSet = wire.NewSet(wire.Struct(new(LoginRouter), "*"))

type LoginRouter struct {
	Login *login.UserLogin
}

func (r *LoginRouter) Register(router *fiber.App) {
	// Add router here
	router.Post("/login/username", r.Login.UserNameLogin)
}
