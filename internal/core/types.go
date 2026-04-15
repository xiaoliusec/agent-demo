package core

type Action struct {
	ToolName string            `json:"tool_name"`
	Args     map[string]string `json:"args"`
	Reason   string            `json:"reason"`
}

type Decision struct {
	ThoughtSummary string  `json:"thought_summary"`
	NextAction     *Action `json:"next_action,omitempty"`
	Done           bool    `json:"done"`
	FinalResult    string  `json:"final_result,omitempty"`
}

type ToolResult struct {
	OK        bool              `json:"ok"`
	Output    string            `json:"output,omitempty"`
	Error     string            `json:"error,omitempty"`
	Artifacts map[string]string `json:"artifacts,omitempty"`
}

type StepRecord struct {
	Index      int
	Thought    string
	Action     *Action
	ToolResult ToolResult
	ErrText    string
}

type StateView struct {
	Goal       string
	StepCount  int
	LastOutput string
}
