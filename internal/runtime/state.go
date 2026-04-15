package runtime

type Status string

const (
	StatusRunning Status = "RUNNING"
	StatusDone    Status = "DONE"
	StatusFailed  Status = "FAILED"
)

type State struct {
	Goal           string
	Status         Status
	StepCount      int
	FailCount      int
	MaxSteps       int
	MaxConsecutive int
	LastOutput     string
	FinalResult    string
}

func NewState(goal string, maxSteps int, maxConsecutive int) *State {
	return &State{
		Goal:           goal,
		Status:         StatusRunning,
		MaxSteps:       maxSteps,
		MaxConsecutive: maxConsecutive,
	}
}
