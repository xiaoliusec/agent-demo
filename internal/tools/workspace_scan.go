package tools

import (
	"context"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"agentdemo/internal/core"
)

type WorkspaceScanTool struct{}

func NewWorkspaceScanTool() *WorkspaceScanTool {
	return &WorkspaceScanTool{}
}

func (t *WorkspaceScanTool) Name() string {
	return "workspace.scan"
}

func (t *WorkspaceScanTool) Run(_ context.Context, action core.Action) (core.ToolResult, error) {
	path := action.Args["path"]
	if path == "" {
		path = "."
	}

	entries, err := os.ReadDir(path)
	if err != nil {
		return core.ToolResult{OK: false, Error: err.Error()}, err
	}

	names := make([]string, 0, len(entries))
	for _, e := range entries {
		name := e.Name()
		if e.IsDir() {
			name += "/"
		}
		names = append(names, filepath.Clean(name))
	}
	sort.Strings(names)

	return core.ToolResult{
		OK:     true,
		Output: strings.Join(names, "\n"),
	}, nil
}
