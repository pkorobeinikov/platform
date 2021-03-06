package docker

import (
	"context"
	"errors"
	"fmt"
	"os/exec"

	"github.com/pkorobeinikov/platform/platform-lib/service/deployment"
	"github.com/pkorobeinikov/platform/platform-lib/service/env"
	"platformctl/internal/cfg"
)

// EnsureSentinelNotRunning needs to be rethought.
func EnsureSentinelNotRunning(ctx context.Context, serviceName string) error {
	var (
		cmd  *exec.Cmd
		args []string
	)

	args = []string{
		cfg.PlatformFlavorContainerRuntimeCtl, "compose",
		"--file",
		deployment.DockerComposeFile,
		"--env-file",
		env.File,
		"ps",
		"--quiet",
		fmt.Sprintf("platform-sentinel-%s", serviceName),
	}

	cmd = exec.CommandContext(ctx, args[0], args[1:]...)
	sentinelStateInCurrentService, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}
	if len(sentinelStateInCurrentService) > 0 {
		return ErrCurrentServiceAlreayRunning
	}

	args = []string{
		cfg.PlatformFlavorContainerRuntimeCtl,
		"container",
		"ls",
		"--all",
		"--format={{.Names}}",
		"--filter=name=platform-sentinel",
	}

	cmd = exec.CommandContext(ctx, args[0], args[1:]...)
	sentinelStateGlobal, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}

	// But sentinel is running globally
	if len(sentinelStateGlobal) > 0 {
		return ErrSentinelAlreadyRunning
	}

	return nil
}

var (
	ErrCurrentServiceAlreayRunning = errors.New("current service already running")
	ErrSentinelAlreadyRunning      = errors.New("another service already running")
)
