package model

import (
	"context"

	"testflow/internal/core"
)

type Model interface {
	Decide(ctx context.Context, state core.StateView) (core.Decision, error)
}

type MockModel struct{}

func NewMockModel() *MockModel {
	return &MockModel{}
}

func (m *MockModel) Decide(_ context.Context, state core.StateView) (core.Decision, error) {
	switch state.StepCount {
	case 0:
		return core.Decision{
			ThoughtSummary: "先扫描页面，了解当前可交互元素",
			NextAction: &core.Action{
				ToolName: "browser.navigate",
				Args: map[string]string{
					"url": "https://example.com",
				},
				Reason: "导航到目标系统",
			},
		}, nil
	case 1:
		return core.Decision{
			ThoughtSummary: "获取页面截图和 DOM 元素",
			NextAction: &core.Action{
				ToolName: "browser.screenshot",
				Args:     map[string]string{},
				Reason:   "获取当前页面状态",
			},
		}, nil
	case 2:
		return core.Decision{
			ThoughtSummary: "分析 DOM 元素，准备进行测试",
			NextAction: &core.Action{
				ToolName: "browser.get_dom",
				Args:     map[string]string{},
				Reason:   "提取可交互元素",
			},
		}, nil
	default:
		return core.Decision{
			ThoughtSummary: "测试流程已完成",
			Done:           true,
			FinalResult:    "Mock 测试流程完成，共执行 3 步",
		}, nil
	}
}
