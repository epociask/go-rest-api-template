//+build wireinject

package main

import (
	"context"

	"github.com/epociask/go-rest-api-template/internal/config"
	"github.com/epociask/go-rest-api-template/internal/handlers"
	"github.com/epociask/go-rest-api-template/internal/server"
	"github.com/epociask/go-rest-api-template/internal/service"
	"github.com/google/wire"
)

func InitializeAndRun(ctx context.Context, cfg config.FilePath) (*server.Server, func(), error) {

	panic(
		wire.Build(
			config.NewConfig,
			config.NewServerConfig,
			serviceModule,
			handlersModule,
			server.New,
		),
	)
}

var serviceModule = wire.NewSet(
	service.Module,
	wire.Bind(new(service.Service), new(*service.ExampleService)),
)

var handlersModule = wire.NewSet(
	handlers.Module,
	wire.Bind(new(handlers.Handlers), new(*handlers.ExampleHandler)),
)
