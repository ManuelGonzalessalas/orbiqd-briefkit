package agent

import (
	"context"
	"errors"
)

type Config struct {
	RuntimeKind   RuntimeKind   `json:"runtimeKind"`
	RuntimeConfig RuntimeConfig `json:"runtimeConfig"`
}

type ConfigRepository interface {
	Exists(ctx context.Context, id AgentID) (bool, error)
	Get(ctx context.Context, id AgentID) (Config, error)
	Update(ctx context.Context, id AgentID, config Config) error
	List(ctx context.Context) ([]AgentID, error)
}

var (
	// ErrAgentConfigNotFound indicates the agent configuration does not exist.
	ErrAgentConfigNotFound = errors.New("agent config not found")

	// ErrAgentIDInvalid indicates the agent identifier is missing or invalid.
	ErrAgentIDInvalid = errors.New("agent id invalid")
)
