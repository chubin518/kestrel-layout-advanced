package service

import (
	"context"
	"os/exec"
)

type ShellService struct {
}

func NewShellService() *ShellService {
	return &ShellService{}
}

// Exec
func (s *ShellService) Exec(ctx context.Context, command string) (string, error) {
	cmd := exec.CommandContext(ctx, "bash", "-c", command)
	buf, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	return string(buf), nil
}
