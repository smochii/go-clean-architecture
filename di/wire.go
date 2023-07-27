//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	"github.com/smochii/go-clean-architecture/handler"
	"github.com/smochii/go-clean-architecture/infrastructure/database"
	"github.com/smochii/go-clean-architecture/route"
	"github.com/smochii/go-clean-architecture/service"
)

var providerSet = wire.NewSet(
	database.NewConnection,
	database.NewUserRepository,

	service.NewUserService,
	service.NewAuthService,

	handler.NewMiddlewareHandler,
	handler.NewAuthHandler,
	handler.NewUserHandler,
)

func Route() *route.Route {
	wire.Build(
		providerSet,
		route.NewRoute,
	)
	return nil
}
