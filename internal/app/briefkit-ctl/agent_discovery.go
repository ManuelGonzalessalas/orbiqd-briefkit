package briefkitctl

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/orbiqd/orbiqd-briefkit/internal/pkg/agent"
)

type AgentDiscoveryCmd struct{}

type AgentDiscoveryOutputItem struct {
	RuntimeKind agent.RuntimeKind `json:"runtimeKind"`
	Found       bool              `json:"found"`
	Version     string            `json:"version"`
}

type AgentDiscoveryOutput struct {
	Items []AgentDiscoveryOutputItem `json:"items"`
	Count int                        `json:"count"`
}

func (command *AgentDiscoveryCmd) Run(ctx context.Context, registry agent.RuntimeRegistry) error {
	kinds, err := registry.List(ctx)
	if err != nil {
		return fmt.Errorf("list runtimes: %w", err)
	}

	items := make([]AgentDiscoveryOutputItem, 0, len(kinds))
	for _, kind := range kinds {
		runtime, err := registry.Get(ctx, kind)
		if err != nil {
			return fmt.Errorf("get runtime %s: %w", kind, err)
		}

		info, err := runtime.GetInfo(ctx)
		if err != nil {
			return fmt.Errorf("get runtime info %s: %w", kind, err)
		}

		found, err := runtime.Discovery(ctx)
		if err != nil {
			return fmt.Errorf("discover runtime %s: %w", kind, err)
		}

		items = append(items, AgentDiscoveryOutputItem{
			RuntimeKind: kind,
			Found:       found,
			Version:     info.Version,
		})
	}

	output := AgentDiscoveryOutput{
		Items: items,
		Count: len(items),
	}

	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(output); err != nil {
		return fmt.Errorf("encode agent discovery output: %w", err)
	}

	return nil
}
