package runtime

import (
	"context"
	"fmt"

	"agentdemo/internal/core"
	"agentdemo/internal/guard"
	"agentdemo/internal/model"
	"agentdemo/internal/store"
	"agentdemo/internal/tools"
)

type Engine struct {
	model model.Model
	tools *tools.Registry
	guard *guard.Policy
	store *store.MemoryStore
}

func NewEngine(model model.Model, tools *tools.Registry, guard *guard.Policy, store *store.MemoryStore) *Engine {
	return &Engine{model: model, tools: tools, guard: guard, store: store}
}

func (e *Engine) Run(ctx context.Context, state *State) error {
	for state.Status == StatusRunning {
		if state.StepCount >= state.MaxSteps {
			state.Status = StatusFailed
			return fmt.Errorf("reach max steps: %d", state.MaxSteps)
		}

		decision, err := e.model.Decide(ctx, core.StateView{
			Goal:       state.Goal,
			StepCount:  state.StepCount,
			LastOutput: state.LastOutput,
		})
		if err != nil {
			state.FailCount++
			if state.FailCount >= state.MaxConsecutive {
				state.Status = StatusFailed
				return fmt.Errorf("model decide failed too many times: %w", err)
			}
			continue
		}

		if decision.Done {
			state.Status = StatusDone
			state.FinalResult = decision.FinalResult
			return nil
		}

		if decision.NextAction == nil {
			state.FailCount++
			if state.FailCount >= state.MaxConsecutive {
				state.Status = StatusFailed
				return fmt.Errorf("empty next action")
			}
			continue
		}

		if err := e.guard.Check(*decision.NextAction); err != nil {
			state.FailCount++
			if state.FailCount >= state.MaxConsecutive {
				state.Status = StatusFailed
				return err
			}
			continue
		}

		result, execErr := e.tools.Execute(ctx, *decision.NextAction)
		record := core.StepRecord{
			Index:      state.StepCount,
			Thought:    decision.ThoughtSummary,
			Action:     decision.NextAction,
			ToolResult: result,
		}
		if execErr != nil {
			record.ErrText = execErr.Error()
		}
		e.store.Append(record)

		state.StepCount++
		if execErr != nil {
			state.FailCount++
			if state.FailCount >= state.MaxConsecutive {
				state.Status = StatusFailed
				return fmt.Errorf("tool failed too many times: %w", execErr)
			}
			continue
		}

		state.FailCount = 0
		state.LastOutput = result.Output
	}

	return nil
}
