package codex

import (
	"encoding/json"
	"fmt"

	"github.com/mcuadros/go-defaults"

	"github.com/orbiqd/orbiqd-briefkit/internal/pkg/agent"
)

// Config defines runtime options for Codex execution.
type Config struct {
	// RequireWorkspaceRepository enforces that codex workdir must be a GIT repository.
	RequireWorkspaceRepository bool `json:"requireWorkspaceRepository" default:"true"`

	// EnableWebSearch allow codex to use internal web-search tool, without allowing it to access internet.
	EnableWebSearch *bool `json:"enableWebSearch"`

	// EnableNetworkAccess
	EnableNetworkAccess *bool `json:"enableNetworkAccess"`
}

// Args renders Codex CLI flags based on the config values.
func (config Config) Args() []string {
	var args []string

	if !config.RequireWorkspaceRepository {
		args = append(args, "--skip-git-repo-check")
	}

	if config.EnableNetworkAccess != nil {
		if *config.EnableNetworkAccess {
			args = append(args, "--config sandbox_workspace_write.network_access=true")
		} else {
			args = append(args, "--config sandbox_workspace_write.network_access=false")
		}
	}

	if config.EnableWebSearch != nil {
		if *config.EnableWebSearch {
			args = append(args, "--config features.web_search_request=true")
		} else {
			args = append(args, "--config features.web_search_request=false")
		}
	}

	return nil
}

func createConfigFromRuntimeConfig(config agent.RuntimeConfig) (Config, error) {
	var result Config

	switch typed := config.(type) {
	case nil:
		break
	case Config:
		result = typed
	case *Config:
		if typed != nil {
			result = *typed
		}
	default:
		payload, err := json.Marshal(config)
		if err != nil {
			return Config{}, fmt.Errorf("marshal codex config: %w", err)
		}

		if err := json.Unmarshal(payload, &result); err != nil {
			return Config{}, fmt.Errorf("unmarshal codex config: %w", err)
		}
	}

	defaults.SetDefaults(&result)

	return result, nil
}
