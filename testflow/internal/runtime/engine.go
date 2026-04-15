package runtime

import (
	"context"
	"fmt"
	"time"

	"testflow/internal/core"
	"testflow/internal/guard"
	"testflow/internal/model"
	"testflow/internal/store"
	"testflow/internal/tools"
)

type Engine struct {
	registry *tools.Registry
	guard    *guard.Policy
	model    model.Model
	store    *store.MemoryStore
}

func NewEngine(registry *tools.Registry, g *guard.Policy, m model.Model, s *store.MemoryStore) *Engine {
	return &Engine{
		registry: registry,
		guard:    g,
		model:    m,
		store:    s,
	}
}

type EventEmitter func(event string, data interface{})

func (e *Engine) Run(ctx context.Context, goal string, emit EventEmitter) error {
	e.store.SetStatus(&core.TestStatus{State: "running", CurrentStep: 0})

	state := core.StateView{
		Goal:          goal,
		StepCount:     0,
		DOMElements:   make([]*core.Element, 0),
		TestCases:     make([]*core.TestCase, 0),
		History:       make([]*core.HistoryEntry, 0),
		TestedSummary: make([]string, 0),
	}

	maxSteps := 100

	for state.StepCount < maxSteps {
		status := e.store.GetStatus()
		if status.State == "stopped" {
			break
		}

		decision, err := e.model.Decide(ctx, state)
		if err != nil {
			e.store.SetStatus(&core.TestStatus{
				State:       "error",
				CurrentStep: state.StepCount,
				LastError:   err.Error(),
			})
			return fmt.Errorf("model decide error: %w", err)
		}

		emit("ai_decision", decision)

		if decision.Done {
			e.store.SetStatus(&core.TestStatus{
				State:       "completed",
				CurrentStep: state.StepCount,
				TotalCases:  len(state.TestCases),
			})
			emit("test_complete", decision.FinalResult)
			return nil
		}

		if decision.NextAction == nil {
			break
		}

		if err := e.guard.CheckAction(decision.NextAction, e.registry); err != nil {
			e.store.SetStatus(&core.TestStatus{
				State:       "error",
				CurrentStep: state.StepCount,
				LastError:   fmt.Sprintf("guard check failed: %v", err),
			})
			return fmt.Errorf("guard check failed: %w", err)
		}

		result, err := e.registry.Execute(ctx, *decision.NextAction)
		if err != nil {
			e.store.SetStatus(&core.TestStatus{
				State:       "error",
				CurrentStep: state.StepCount,
				LastError:   err.Error(),
			})
			return fmt.Errorf("tool execution error: %w", err)
		}

		e.store.AddStep(decision.NextAction, &result)

		state.LastOutput = result.Output
		state.StepCount++

		emit("step_complete", map[string]interface{}{
			"step":   state.StepCount,
			"result": result,
		})

		if decision.TestCase != nil {
			e.store.AddTestCase(decision.TestCase)
			state.TestCases = append(state.TestCases, decision.TestCase)
			emit("test_case_generated", decision.TestCase)
		}

		time.Sleep(100 * time.Millisecond)
	}

	e.store.SetStatus(&core.TestStatus{
		State:       "completed",
		CurrentStep: state.StepCount,
	})
	return nil
}

func (e *Engine) Stop() {
	e.store.SetStatus(&core.TestStatus{State: "stopped"})
}
