package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/wire"
	"github.com/noraincode/fiber-seed/internal/app/server/handlers/whoami"
)

var whoamiRouterSet = wire.NewSet(wire.Struct(new(WhoamiRouter), "*"))

type WhoamiRouter struct {
	Whoami *whoami.Whoami
}

func (r *WhoamiRouter) Register(router *fiber.App) {
	// Add router here
	router.Get("/whoami", r.Whoami.Whoami)
}
