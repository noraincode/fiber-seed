package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/wire"
	"github.com/noraincode/fiber-seed/internal/app/config"
)

var RouterSet = wire.NewSet(
	whoamiRouterSet,
	loginRouterSet,

	wire.Struct(new(Router), "*"),
	wire.Bind(new(IRouter), new(*Router)),
)

type Router struct {
	Config *config.Config

	WhoamiRouter WhoamiRouter
	LoginRouter  LoginRouter
}

type IRouter interface {
	Register(router *fiber.App)
}

func (r *Router) Register(router *fiber.App) {
	r.WhoamiRouter.Register(router)
	r.LoginRouter.Register(router)
}
