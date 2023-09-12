package graceful

import (
	"fmt"
	"os"
	"runtime"

	"github.com/chubin518/kestrel-layout-advanced/buildinfo"
	"github.com/chubin518/kestrel-layout-advanced/pkg/result"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/process"
)

const (
	B  = 1
	KB = 1024 * B
	MB = 1024 * KB
	GB = 1024 * MB
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
func (app *WebGraceful) UseRoutes(routes ...func(gin.IRouter)) *WebGraceful {
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
	return app
}

// health web server health check
func health(ctx *gin.Context) {
	metrics := make(map[string]any)
	if osCpu, err := cpu.Percent(0, false); err == nil && len(osCpu) > 0 {
		metrics["cpu_percent"] = osCpu[0]
	}
	if cores, err := cpu.Counts(false); err == nil {
		metrics["cpu_cores"] = cores
	}
	if osRAM, err := mem.VirtualMemory(); err == nil {
		metrics["ram_total"] = osRAM.Total / MB
		metrics["ram_used"] = osRAM.Used / MB
		metrics["ram_percent"] = osRAM.UsedPercent
	}

	if osDisk, err := disk.Usage("/"); err == nil {
		metrics["disk_total"] = osDisk.Total / MB
		metrics["disk_used"] = osDisk.Used / MB
		metrics["disk_percent"] = osDisk.UsedPercent
	}

	app := map[string]any{
		"name":          buildinfo.Name,
		"version":       buildinfo.Version,
		"env":           buildinfo.Environment,
		"build_time":    buildinfo.BuildTime,
		"build_version": buildinfo.BuildVersion,
	}
	if proc, err := process.NewProcess(int32(os.Getpid())); err == nil {
		if cpu, err := proc.CPUPercent(); err == nil {
			app["proc_cpu"] = cpu
		}
		if ram, err := proc.MemoryInfo(); err == nil {
			app["proc_ram"] = ram.RSS / MB
		}
	}

	data := map[string]any{
		"status":  "UP",
		"app":     app,
		"metrics": metrics,
		"runtime": map[string]any{
			"version":    fmt.Sprintf("go version %s %s/%s", runtime.Version(), runtime.GOOS, runtime.GOARCH),
			"go_routine": runtime.NumGoroutine(),
		},
	}

	result.OK.WithData(data).JSON(ctx)
}
