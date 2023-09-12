package handler

import (
	"github.com/chubin518/kestrel-layout-advanced/internal/service"
	"github.com/chubin518/kestrel-layout-advanced/pkg/result"
	"github.com/gin-gonic/gin"
)

type ShellHandler struct {
	shell *service.ShellService
}

func NewShellHandler(shell *service.ShellService) *ShellHandler {
	return &ShellHandler{shell: shell}
}

// Exec
func (h *ShellHandler) Exec(ctx *gin.Context) {
	command, ok := ctx.GetPostForm("cmd")
	if !ok || command == "" {
		result.BadRequest.JSON(ctx)
		return
	}
	output, err := h.shell.Exec(ctx, command)
	if err != nil {
		result.InternalServerError.WithMessage(err.Error()).JSON(ctx)
		return
	}
	result.OK.WithData(output).JSON(ctx)
}
