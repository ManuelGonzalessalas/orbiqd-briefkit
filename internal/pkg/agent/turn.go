package agent

import "time"

type TurnID string

type TurnRequest struct {
}

type TurnResponse struct {
}

type TurnStatus struct {
	CreatedAt  time.Duration
	UpdatedAt  time.Duration
	FinishedAt time.Duration
}
