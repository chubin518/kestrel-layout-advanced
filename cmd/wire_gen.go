// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

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

// Injectors from wire.go:

func BuildContainer(iConfig config.IConfig, iLogging logging.ILogging) (*Container, error) {
	fileService := service.NewFileService()
	fileHandler := handler.NewFileHandler(fileService)
	shellService := service.NewShellService()
	shellHandler := handler.NewShellHandler(shellService)
	routeFunc := routes.RegisterRoutes(fileHandler, shellHandler)
	container := &Container{
		ConfigRoutes: routeFunc,
	}
	return container, nil
}

// wire.go:

type Container struct {
	ConfigRoutes graceful.RouteFunc
}

var ContainerSet = wire.NewSet(wire.Struct(new(Container), "*"))

var ServiceSet = wire.NewSet(service.NewShellService, service.NewFileService)

var HandlerSet = wire.NewSet(handler.NewShellHandler, handler.NewFileHandler)

var RoutesSet = wire.NewSet(routes.RegisterRoutes)
