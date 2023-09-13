//go:build wireinject
// +build wireinject

package main

import (
	"github.com/chubin518/kestrel-layout-advanced/internal/routes"
	"github.com/chubin518/kestrel-layout-advanced/internal/routes/handler"
	"github.com/chubin518/kestrel-layout-advanced/internal/service"
	"github.com/chubin518/kestrel-layout-advanced/pkg/config"
	"github.com/chubin518/kestrel-layout-advanced/pkg/graceful"
	"github.com/chubin518/kestrel-layout-advanced/pkg/logging"
	"github.com/google/wire"
)

type Container struct {
	ConfigRoutes graceful.RouteFunc
}

var ContainerSet = wire.NewSet(wire.Struct(new(Container), "*"))

var ServiceSet = wire.NewSet(
	service.NewShellService,
	service.NewFileService,
)

var HandlerSet = wire.NewSet(
	handler.NewShellHandler,
	handler.NewFileHandler,
)

var RoutesSet = wire.NewSet(
	routes.RegisterRoutes,
)

func BuildContainer(config.IConfig, logging.ILogging) (*Container, error) {
	panic(wire.Build(ServiceSet, HandlerSet, RoutesSet, ContainerSet))
}
