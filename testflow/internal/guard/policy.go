package guard

import (
	"fmt"

	"testflow/internal/core"
	"testflow/internal/tools"
)

type Policy struct {
	allowedTools      map[string]bool
	dangerousPatterns []string
	maxRetries        int
}

func NewPolicy() *Policy {
	return &Policy{
		allowedTools: make(map[string]bool),
		dangerousPatterns: []string{
			"rm -rf",
			"delete",
			"drop",
			"format",
		},
		maxRetries: 3,
	}
}

func (p *Policy) SetAllowedTools(tools []string) {
	for _, t := range tools {
		p.allowedTools[t] = true
	}
}

func (p *Policy) CheckTool(toolName string) error {
	if len(p.allowedTools) == 0 {
		return nil
	}
	if !p.allowedTools[toolName] {
		return fmt.Errorf("tool '%s' is not allowed by policy", toolName)
	}
	return nil
}

func (p *Policy) CheckAction(action *core.Action, registry *tools.Registry) error {
	if action == nil {
		return fmt.Errorf("action is nil")
	}
	if !registry.Exists(action.ToolName) {
		return fmt.Errorf("tool '%s' does not exist", action.ToolName)
	}
	if err := p.CheckTool(action.ToolName); err != nil {
		return err
	}
	return nil
}

func (p *Policy) GetMaxRetries() int {
	return p.maxRetries
}
