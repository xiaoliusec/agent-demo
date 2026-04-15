package tools

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"agentdemo/internal/core"
)

type ReportWriteTool struct{}

func NewReportWriteTool() *ReportWriteTool {
	return &ReportWriteTool{}
}

func (t *ReportWriteTool) Name() string {
	return "report.write"
}

func (t *ReportWriteTool) Run(_ context.Context, action core.Action) (core.ToolResult, error) {
	path := action.Args["path"]
	content := action.Args["content"]
	if path == "" {
		return core.ToolResult{OK: false, Error: "missing path"}, fmt.Errorf("missing path")
	}

	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return core.ToolResult{OK: false, Error: err.Error()}, err
	}

	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		return core.ToolResult{OK: false, Error: err.Error()}, err
	}

	return core.ToolResult{
		OK:     true,
		Output: "report written",
		Artifacts: map[string]string{
			"report": path,
		},
	}, nil
}
