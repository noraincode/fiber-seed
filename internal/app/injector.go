package app

import (
	"github.com/google/wire"
	"github.com/noraincode/fiber-seed/internal/app/server"
)

//lint:ignore U1000 For wire injection
var injectorSet = wire.NewSet(wire.Struct(new(Injector), "*"))

type Injector struct {
	Srv *server.Server
}
