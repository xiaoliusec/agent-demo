package store

import (
	"sync"

	"testflow/internal/core"
)

type MemoryStore struct {
	mu         sync.RWMutex
	steps      []*StepRecord
	testCases  []*core.TestCase
	config     *core.Config
	testStatus *core.TestStatus
}

type StepRecord struct {
	Step      int
	Action    *core.Action
	Result    *core.ToolResult
	Timestamp string
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		steps:      make([]*StepRecord, 0),
		testCases:  make([]*core.TestCase, 0),
		testStatus: &core.TestStatus{State: "idle"},
	}
}

func (s *MemoryStore) AddStep(action *core.Action, result *core.ToolResult) {
	s.mu.Lock()
	defer s.mu.Unlock()
	record := &StepRecord{
		Step:   len(s.steps),
		Action: action,
		Result: result,
	}
	s.steps = append(s.steps, record)
	s.testStatus.CurrentStep = len(s.steps)
}

func (s *MemoryStore) GetSteps() []*StepRecord {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.steps
}

func (s *MemoryStore) AddTestCase(tc *core.TestCase) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.testCases = append(s.testCases, tc)
	s.testStatus.TotalCases = len(s.testCases)
	if tc.Status == "passed" {
		s.testStatus.PassedCases++
	} else if tc.Status == "failed" {
		s.testStatus.FailedCases++
	}
}

func (s *MemoryStore) GetTestCases() []*core.TestCase {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.testCases
}

func (s *MemoryStore) SetConfig(config *core.Config) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.config = config
}

func (s *MemoryStore) GetConfig() *core.Config {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.config
}

func (s *MemoryStore) SetStatus(status *core.TestStatus) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.testStatus = status
}

func (s *MemoryStore) GetStatus() *core.TestStatus {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.testStatus
}

func (s *MemoryStore) Reset() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.steps = make([]*StepRecord, 0)
	s.testCases = make([]*core.TestCase, 0)
	s.testStatus = &core.TestStatus{State: "idle"}
}
