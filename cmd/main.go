package main

import (
	"context"
	"os"
	"time"

	"github.com/chubin518/kestrel-layout-advanced/pkg/graceful"
	"github.com/chubin518/kestrel-layout-advanced/pkg/logging"
	"github.com/chubin518/kestrel-layout-advanced/pkg/middleware"
	"github.com/gin-contrib/cors"
)

func main() {

	if err := graceful.
		CreateWebGraceful().
		UseAddr(":8080").
		UseCors(func(conf *cors.Config) {
			conf.AllowAllOrigins = true
			conf.AllowHeaders = []string{"*"}
			conf.AllowMethods = []string{"GET", "PUT", "POST", "PATCH", "DELETE", "OPTIONS"}
			conf.MaxAge = 12 * time.Hour
		}).
		UseMiddlewares(middleware.NewRequestId(), middleware.NewRecovery()).
		UseHealth().
		Configure(func(wg *graceful.WebGraceful) {
			contarter, err := BuildContainer(wg.GetConfig(), wg.GetLogging())
			if err != nil {
				logging.Fatal("Error creating container: %v", err)
				return
			}
			wg.UseRoutes(contarter.ConfigRoutes)
		}).
		RunWithContext(context.Background()); err != nil {
		logging.Error("error: %v", err)
		os.Exit(1)
	}
}
