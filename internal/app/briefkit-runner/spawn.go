package briefkit_runner

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"syscall"

	"github.com/orbiqd/orbiqd-briefkit/internal/pkg/agent"
	"github.com/orbiqd/orbiqd-briefkit/internal/pkg/process"
)

const executableName = "briefkit-runner"

func Spawn(ctx context.Context, executionId agent.ExecutionID) error {
	var executablePath string

	if envPath, ok := os.LookupEnv("BRIEFKIT_RUNNER_PATH"); ok {
		if _, err := os.Stat(envPath); err != nil {
			return fmt.Errorf("executable from BRIEFKIT_RUNNER_PATH not found: %w", err)
		}
		executablePath = envPath
	} else {
		if self, err := os.Executable(); err == nil {
			candidate := filepath.Join(filepath.Dir(self), executableName)
			if _, err := os.Stat(candidate); err == nil {
				executablePath = candidate
			}
		}

		if executablePath == "" {
			var err error
			executablePath, err = process.LookupExecutable(ctx, []string{executableName})
			if err != nil {
				return fmt.Errorf("lookup executable: %w", err)
			}
		}
	}

	cmd := exec.Command(executablePath, string(executionId))

	cmd.SysProcAttr = &syscall.SysProcAttr{
		Setsid: true,
	}

	cmd.Stdin = nil
	cmd.Stdout = nil
	cmd.Stderr = nil

	if err := cmd.Start(); err != nil {
		return fmt.Errorf("start process: %w", err)
	}

	if err := cmd.Process.Release(); err != nil {
		return fmt.Errorf("release process: %w", err)
	}

	return nil
}