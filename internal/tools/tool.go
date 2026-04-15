package tools

import (
	"context"
	"fmt"

	"agentdemo/internal/core"
)

type Tool interface {
	Name() string
	Run(ctx context.Context, action core.Action) (core.ToolResult, error)
}

type Registry struct {
	items map[string]Tool
}

func NewRegistry() *Registry {
	return &Registry{items: make(map[string]Tool)}
}

func (r *Registry) Register(t Tool) {
	r.items[t.Name()] = t
}

func (r *Registry) Exists(name string) bool {
	_, ok := r.items[name]
	return ok
}

func (r *Registry) Execute(ctx context.Context, action core.Action) (core.ToolResult, error) {
	t, ok := r.items[action.ToolName]
	if !ok {
		return core.ToolResult{OK: false, Error: "tool not found"}, fmt.Errorf("tool not found: %s", action.ToolName)
	}
	return t.Run(ctx, action)
}
