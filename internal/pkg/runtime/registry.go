package runtime

import (
	"context"
	"sort"

	"github.com/orbiqd/orbiqd-briefkit/internal/pkg/agent"
	"github.com/orbiqd/orbiqd-briefkit/internal/pkg/runtime/claude"
	"github.com/orbiqd/orbiqd-briefkit/internal/pkg/runtime/codex"
	"github.com/orbiqd/orbiqd-briefkit/internal/pkg/runtime/gemini"
)

type Registry struct {
	runtime map[agent.RuntimeKind]agent.Runtime
}

func NewRegistry() *Registry {
	return &Registry{
		runtime: map[agent.RuntimeKind]agent.Runtime{
			gemini.Gemini: gemini.NewRuntime(),
			claude.Claude: claude.NewRuntime(),
			codex.Codex:   codex.NewRuntime(),
		},
	}
}

func (registry Registry) Get(ctx context.Context, kind agent.RuntimeKind) (agent.Runtime, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	runtime, ok := registry.runtime[kind]
	if !ok {
		return nil, agent.ErrRuntimeNotFound
	}

	return runtime, nil
}

func (registry Registry) List(ctx context.Context) ([]agent.RuntimeKind, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	kinds := make([]agent.RuntimeKind, 0, len(registry.runtime))
	for kind := range registry.runtime {
		if err := ctx.Err(); err != nil {
			return nil, err
		}

		kinds = append(kinds, kind)
	}

	sort.Slice(kinds, func(i, j int) bool {
		return kinds[i] < kinds[j]
	})

	return kinds, nil
}
