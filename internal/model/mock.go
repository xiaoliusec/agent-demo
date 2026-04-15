package model

import (
	"context"
	"fmt"

	"agentdemo/internal/core"
)

type MockModel struct{}

func NewMockModel() *MockModel {
	return &MockModel{}
}

func (m *MockModel) Decide(_ context.Context, state core.StateView) (core.Decision, error) {
	switch state.StepCount {
	case 0:
		return core.Decision{
			ThoughtSummary: "先扫描工作区，了解当前可用文件",
			NextAction: &core.Action{
				ToolName: "workspace.scan",
				Args: map[string]string{
					"path": ".",
				},
				Reason: "执行任何任务前先收集上下文",
			},
		}, nil
	case 1:
		report := fmt.Sprintf("# Demo 执行报告\n\n- 任务目标: %s\n- 最近观察:\n\n```text\n%s\n```\n", state.Goal, state.LastOutput)
		return core.Decision{
			ThoughtSummary: "根据扫描结果输出一份报告，作为 Demo 产物",
			NextAction: &core.Action{
				ToolName: "report.write",
				Args: map[string]string{
					"path":    "output/demo-report.md",
					"content": report,
				},
				Reason: "生成可见产物，证明流程完整可运行",
			},
		}, nil
	default:
		return core.Decision{
			ThoughtSummary: "核心流程已完成",
			Done:           true,
			FinalResult:    "Demo 已自动完成扫描与报告输出",
		}, nil
	}
}
