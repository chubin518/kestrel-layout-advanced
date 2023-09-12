package routes

import (
	"io/fs"
	"net/http"

	"github.com/chubin518/kestrel-layout-advanced/internal/routes/handler"
	"github.com/chubin518/kestrel-layout-advanced/pkg/logging"
	"github.com/chubin518/kestrel-layout-advanced/webui"
	"github.com/gin-gonic/gin"
)

// InitRoutes
func InitRoutes(fileHandler *handler.FileHandler, shellHandler *handler.ShellHandler) func(gin.IRouter) {
	return func(router gin.IRouter) {
		dist, err := fs.Sub(webui.StaticFS, "dist/assets")
		if err != nil {
			logging.Fatal("load assets failed: %v", err)
			return
		}
		router.StaticFS("/assets", http.FS(dist))

		router.GET("/", func(ctx *gin.Context) {
			rawIndex, _ := webui.StaticFS.ReadFile("dist/index.html")
			ctx.Writer.Write(rawIndex)
		})

		router.GET("/file/list", fileHandler.List)
		router.GET("/file/treeList", fileHandler.TreeList)
		router.GET("/file/preview", fileHandler.Preview)
		router.GET("/file/download", fileHandler.Download)

		router.POST("/shell/exec", shellHandler.Exec)
	}
}
