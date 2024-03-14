package handlers

import (
	"github.com/google/wire"
	"github.com/noraincode/fiber-seed/internal/app/server/handlers/login"
	"github.com/noraincode/fiber-seed/internal/app/server/handlers/whoami"
)

var HandlersSet = wire.NewSet(
	whoami.WireSet,
	login.LoginSet,
)
