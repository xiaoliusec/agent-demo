package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"agentdemo/internal/guard"
	"agentdemo/internal/model"
	"agentdemo/internal/runtime"
	"agentdemo/internal/store"
	"agentdemo/internal/tools"
)

func main() {
	goal := flag.String("goal", "先扫描工作区，再输出一份 demo 报告", "任务目标")
	maxSteps := flag.Int("max-steps", 10, "最大执行步数")
	maxConsecutiveFails := flag.Int("max-fails", 3, "最大连续失败次数")
	flag.Parse()

	registry := tools.NewRegistry()
	registry.Register(tools.NewWorkspaceScanTool())
	registry.Register(tools.NewReportWriteTool())

	policy := guard.NewPolicy(registry)
	memoryStore := store.NewMemoryStore()
	mockModel := model.NewMockModel()

	engine := runtime.NewEngine(mockModel, registry, policy, memoryStore)
	state := runtime.NewState(*goal, *maxSteps, *maxConsecutiveFails)

	if err := engine.Run(context.Background(), state); err != nil {
		log.Fatalf("agent run failed: %v", err)
	}

	fmt.Printf("Status: %s\n", state.Status)
	fmt.Printf("Steps: %d\n", state.StepCount)
	fmt.Printf("Result: %s\n", state.FinalResult)

	fmt.Println("--- Step Logs ---")
	for _, step := range memoryStore.List() {
		actionName := "<none>"
		if step.Action != nil {
			actionName = step.Action.ToolName
		}
		fmt.Printf("[%d] %s -> %s (ok=%v)\n", step.Index, step.Thought, actionName, step.ToolResult.OK)
	}
}
