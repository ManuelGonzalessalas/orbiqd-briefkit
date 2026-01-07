package gemini

import (
	"context"

	"github.com/orbiqd/orbiqd-briefkit/internal/pkg/agent"
)

type Instance struct {
}

func (instance *Instance) Events() <-chan agent.RuntimeEvent {
	//TODO implement me
	panic("implement me")
}

func (instance *Instance) Wait(ctx context.Context) (agent.RuntimeResult, error) {
	//TODO implement me
	panic("implement me")
}
