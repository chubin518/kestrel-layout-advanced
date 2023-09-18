package graceful

import (
	"github.com/chubin518/kestrel-layout-advanced/pkg/result"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// DefaultNoMethod
func DefaultNoMethod(ctx *gin.Context) {
	result.MethodNotAllowed.JSON(ctx)
}

// DefaultNoRoute
func DefaultNoRoute(ctx *gin.Context) {
	result.NotFound.JSON(ctx)
}

// NoMethod
func (app *WebGraceful) NoMethod(handlers ...gin.HandlerFunc) *WebGraceful {
	app.router.NoMethod(handlers...)
	return app
}

// NoRoute
func (app *WebGraceful) NoRoute(handlers ...gin.HandlerFunc) *WebGraceful {
	app.router.NoRoute(handlers...)
	return app
}

// UseRoutes
func (app *WebGraceful) UseRoutes(routes ...RouteFunc) *WebGraceful {
	for _, apply := range routes {
		if apply != nil {
			apply(app.router)
		}
	}
	return app
}

// UseCors
func (app *WebGraceful) UseCors(configure func(*cors.Config)) *WebGraceful {
	conf := cors.DefaultConfig()
	configure(&conf)
	app.router.Use(cors.New(conf))
	return app
}

// UseMiddlewares
func (app *WebGraceful) UseMiddlewares(middlewares ...gin.HandlerFunc) *WebGraceful {
	app.router.Use(middlewares...)
	return app
}

// UseHealth
func (app *WebGraceful) UseHealth() *WebGraceful {
	app.router.GET("/health", health)
	app.router.GET("/health/metrics", metrics)
	return app
}
