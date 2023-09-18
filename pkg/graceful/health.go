package graceful

import (
	"fmt"
	"os"
	"runtime"

	"github.com/chubin518/kestrel-layout-advanced/buildinfo"
	"github.com/chubin518/kestrel-layout-advanced/pkg/result"
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

// metrics
func metrics(ctx *gin.Context) {
	server := make(map[string]any)
	if osCpu, err := cpu.Percent(0, false); err == nil && len(osCpu) > 0 {
		server["cpu_percent"] = osCpu[0]
	}
	if cores, err := cpu.Counts(false); err == nil {
		server["cpu_cores"] = cores
	}
	if osRAM, err := mem.VirtualMemory(); err == nil {
		server["ram_total"] = osRAM.Total / GB
		server["ram_used"] = osRAM.Used / GB
		server["ram_percent"] = osRAM.UsedPercent
	}

	if osDisk, err := disk.Usage("/"); err == nil {
		server["disk_total"] = osDisk.Total / GB
		server["disk_used"] = osDisk.Used / GB
		server["disk_percent"] = osDisk.UsedPercent
	}

	app := map[string]any{
		"name":       buildinfo.Name,
		"version":    buildinfo.Version,
		"env":        buildinfo.Environment,
		"build_time": buildinfo.BuildTime,
		"go_version": buildinfo.GoVersion,
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
		"app":    app,
		"server": server,
		"runtime": map[string]any{
			"version":   fmt.Sprintf("go version %s %s/%s", runtime.Version(), runtime.GOOS, runtime.GOARCH),
			"goroutine": runtime.NumGoroutine(),
		},
	}

	result.OK.WithData(data).JSON(ctx)
}

// health web server health check
func health(ctx *gin.Context) {
	result.OK.WithData(map[string]any{
		"status":     "UP",
		"name":       buildinfo.Name,
		"version":    buildinfo.Version,
		"env":        buildinfo.Environment,
		"build_time": buildinfo.BuildTime,
		"go_version": buildinfo.GoVersion,
	}).JSON(ctx)
}
