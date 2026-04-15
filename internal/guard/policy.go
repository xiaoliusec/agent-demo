package guard

import (
	"fmt"

	"agentdemo/internal/core"
	"agentdemo/internal/tools"
)

type Policy struct {
	registry *tools.Registry
}

func NewPolicy(registry *tools.Registry) *Policy {
	return &Policy{registry: registry}
}

func (p *Policy) Check(action core.Action) error {
	if action.ToolName == "" {
		return fmt.Errorf("empty tool name")
	}
	if !p.registry.Exists(action.ToolName) {
		return fmt.Errorf("tool not allowed: %s", action.ToolName)
	}
	return nil
}
