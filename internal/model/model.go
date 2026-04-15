package model

import (
	"context"

	"agentdemo/internal/core"
)

type Model interface {
	Decide(ctx context.Context, state core.StateView) (core.Decision, error)
}
