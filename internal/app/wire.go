//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.
package app

import (
	"context"

	"github.com/google/wire"
	"github.com/noraincode/fiber-seed/internal/app/config"
	"github.com/noraincode/fiber-seed/internal/app/server"
	"github.com/noraincode/fiber-seed/internal/app/server/handlers"
	"github.com/noraincode/fiber-seed/internal/app/server/routes"
)

func BuildInjector(ctx context.Context) (*Injector, func(), error) {
	wire.Build(
		config.GetConfig,
		injectorSet,
		server.NewServer,
		routes.RouterSet,
		handlers.HandlersSet,
	)

	return new(Injector), nil, nil
}
