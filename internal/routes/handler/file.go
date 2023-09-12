package handler

import (
	"fmt"
	"net/http"
	"os"

	"github.com/chubin518/kestrel-layout-advanced/internal/service"
	"github.com/chubin518/kestrel-layout-advanced/pkg/logging"
	"github.com/chubin518/kestrel-layout-advanced/pkg/result"
	"github.com/gin-gonic/gin"
)

type FileHandler struct {
	fileService *service.FileService
}

func NewFileHandler(fileService *service.FileService) *FileHandler {
	return &FileHandler{fileService: fileService}
}

// Preview
func (h *FileHandler) Preview(ctx *gin.Context) {
	path, ok := ctx.GetQuery("path")
	if !ok || path == "" {
		ctx.String(http.StatusInternalServerError, "path is required")
		return
	}
	ctx.File(path)
}

// Download
func (h *FileHandler) Download(ctx *gin.Context) {
	filePath, ok := ctx.GetQuery("path")
	if !ok || filePath == "" {
		ctx.String(http.StatusInternalServerError, "path is required")
		return
	}

	file, err := os.Open(filePath)
	if err != nil {
		ctx.String(http.StatusInternalServerError, fmt.Sprintf("Error: %s", err.Error()))
		return
	}
	defer file.Close()

	fileInfo, _ := file.Stat()

	ctx.Writer.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", fileInfo.Name()))
	ctx.Writer.Header().Set("Content-Type", "application/octet-stream")
	ctx.Writer.Header().Set("Content-Length", fmt.Sprintf("%d", fileInfo.Size()))

	http.ServeContent(ctx.Writer, ctx.Request, fileInfo.Name(), fileInfo.ModTime(), file)
}

// List
func (h *FileHandler) List(ctx *gin.Context) {
	dir, ok := ctx.GetQuery("dir")
	if !ok || dir == "" {
		result.BadRequest.WithMessage("dir is required").JSON(ctx)
		return
	}
	list, err := h.fileService.List(ctx, dir)
	if err != nil {
		result.ServiceError.WithMessage(err.Error()).JSON(ctx)
		return
	}
	result.OK.WithData(list).JSON(ctx)
}

// TreeList
func (h *FileHandler) TreeList(ctx *gin.Context) {
	dir := ctx.Query("dir")
	tn, err := h.fileService.TreeList(ctx, dir)
	if err != nil {
		logging.ErrorContext(ctx, "failed to get tree list: %v", err)
		result.ServiceError.WithMessage(err.Error()).JSON(ctx)
		return
	}
	logging.InfoContext(ctx, "dictionary[%s] get tree list [%v]", dir, tn)
	result.OK.WithData(tn).JSON(ctx)
}
