package core

type Config struct {
	TargetURL   string `json:"target_url"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	AIProvider  string `json:"ai_provider"`
	APIKey      string `json:"api_key"`
	APIEndpoint string `json:"api_endpoint"`
	ModelName   string `json:"model_name"`
	MaxCases    int    `json:"max_cases"`
}

type TestStatus struct {
	State       string `json:"state"`
	CurrentStep int    `json:"current_step"`
	TotalCases  int    `json:"total_cases"`
	PassedCases int    `json:"passed_cases"`
	FailedCases int    `json:"failed_cases"`
	CurrentURL  string `json:"current_url"`
	LastError   string `json:"last_error"`
}

type Element struct {
	Tag         string `json:"tag"`
	ID          string `json:"id"`
	Class       string `json:"class"`
	Text        string `json:"text"`
	Type        string `json:"type"`
	Placeholder string `json:"placeholder"`
	XPath       string `json:"xpath"`
}

type Action struct {
	ToolName string            `json:"tool_name"`
	Args     map[string]string `json:"args"`
	Reason   string            `json:"reason"`
}

type TestCase struct {
	ID             string   `json:"id"`
	Module         string   `json:"module"`
	Name           string   `json:"name"`
	Precondition   string   `json:"precondition"`
	Steps          []string `json:"steps"`
	ExpectedResult string   `json:"expected_result"`
	ActualResult   string   `json:"actual_result"`
	Status         string   `json:"status"`
	Priority       string   `json:"priority"`
	Method         string   `json:"method"`
}

type Decision struct {
	ThoughtSummary string    `json:"thought_summary"`
	NextAction     *Action   `json:"next_action"`
	Done           bool      `json:"done"`
	FinalResult    string    `json:"final_result"`
	TestCase       *TestCase `json:"test_case"`
}

type HistoryEntry struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type StateView struct {
	Goal          string          `json:"goal"`
	StepCount     int             `json:"step_count"`
	LastOutput    string          `json:"last_output"`
	Screenshot    []byte          `json:"screenshot"`
	DOMElements   []*Element      `json:"dom_elements"`
	TestCases     []*TestCase     `json:"test_cases"`
	History       []*HistoryEntry `json:"history"`
	TestedSummary []string        `json:"tested_summary"`
}

type ToolResult struct {
	OK        bool              `json:"ok"`
	Output    string            `json:"output"`
	Error     string            `json:"error"`
	Artifacts map[string]string `json:"artifacts"`
}
