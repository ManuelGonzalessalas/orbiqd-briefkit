package codex

import (
	"context"
	"errors"
	"fmt"
	"os/exec"
	"regexp"
	"strings"

	"github.com/orbiqd/orbiqd-briefkit/internal/pkg/agent"
	"github.com/orbiqd/orbiqd-briefkit/internal/pkg/cli"
	"github.com/orbiqd/orbiqd-briefkit/internal/pkg/process"
)

var semverPattern = regexp.MustCompile(`\d+\.\d+\.\d+`)

const Codex = agent.RuntimeKind("codex")

type Runtime struct {
}

func NewRuntime() *Runtime {
	return &Runtime{}
}

func (runtime *Runtime) Execute(ctx context.Context, id agent.ExecutionID, config agent.RuntimeConfig, input agent.ExecutionInput) (agent.RuntimeInstance, error) {
	parsedConfig, err := createConfigFromRuntimeConfig(config)
	if err != nil {
		return nil, err
	}

	logDir, err := cli.ResolveRuntimeLogDir()
	if err != nil {
		return nil, err
	}

	instance, err := newInstance(ctx, id, logDir, parsedConfig, input)
	if err != nil {
		return nil, err
	}
	return instance, nil
}

func (runtime *Runtime) Discovery(ctx context.Context) (bool, error) {
	if err := ctx.Err(); err != nil {
		return false, err
	}

	_, err := process.LookupExecutable(ctx, []string{"codex"})
	if err == nil {
		return true, nil
	}

	if errors.Is(err, exec.ErrNotFound) {
		return false, nil
	}

	return false, err
}

func (runtime *Runtime) GetDefaultConfig(ctx context.Context) (agent.RuntimeConfig, error) {
	return nil, fmt.Errorf("not implemented")
}

func (runtime *Runtime) GetInfo(ctx context.Context) (agent.RuntimeInfo, error) {
	if err := ctx.Err(); err != nil {
		return agent.RuntimeInfo{}, err
	}

	path, err := process.LookupExecutable(ctx, []string{"codex"})
	if err != nil {
		return agent.RuntimeInfo{}, fmt.Errorf("lookup codex executable: %w", err)
	}

	output, err := exec.CommandContext(ctx, path, "--version").CombinedOutput()
	if err != nil {
		return agent.RuntimeInfo{}, fmt.Errorf("read codex version: %w", err)
	}

	version := semverPattern.FindString(string(output))
	if version == "" {
		return agent.RuntimeInfo{}, fmt.Errorf("parse codex version from output: %s", strings.TrimSpace(string(output)))
	}

	return agent.RuntimeInfo{Version: version}, nil
}
