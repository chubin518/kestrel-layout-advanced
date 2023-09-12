//go:build wireinject
// +build wireinject

package main

import (
	"github.com/chubin518/kestrel-layout-advanced/internal/routes"
	"github.com/chubin518/kestrel-layout-advanced/internal/routes/handler"
	"github.com/chubin518/kestrel-layout-advanced/internal/service"
	"github.com/chubin518/kestrel-layout-advanced/pkg/config"
	"github.com/chubin518/kestrel-layout-advanced/pkg/logging"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

var ServiceSet = wire.NewSet(
	service.NewShellService,
	service.NewFileService,
)

var HandlerSet = wire.NewSet(
	handler.NewShellHandler,
	handler.NewFileHandler,
)

var RoutesSet = wire.NewSet(
	routes.InitRoutes,
)

func InitRoutes(config.IConfig, logging.ILogging) func(gin.IRouter) {
	panic(wire.Build(ServiceSet, HandlerSet, RoutesSet))
}
