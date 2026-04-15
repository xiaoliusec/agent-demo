package store

import (
	"sync"

	"agentdemo/internal/core"
)

type MemoryStore struct {
	mu    sync.Mutex
	steps []core.StepRecord
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{steps: make([]core.StepRecord, 0, 32)}
}

func (s *MemoryStore) Append(step core.StepRecord) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.steps = append(s.steps, step)
}

func (s *MemoryStore) List() []core.StepRecord {
	s.mu.Lock()
	defer s.mu.Unlock()
	out := make([]core.StepRecord, len(s.steps))
	copy(out, s.steps)
	return out
}
