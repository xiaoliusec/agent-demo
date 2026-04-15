package main

import (
	"context"
	"fmt"
	"log"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"

	"testflow"
	"testflow/internal/browser"
	"testflow/internal/core"
	"testflow/internal/guard"
	"testflow/internal/model"
	agentruntime "testflow/internal/runtime"
	"testflow/internal/store"
	"testflow/internal/tools"
)

type App struct {
	ctx      context.Context
	runtime  interface{}
	registry *tools.Registry
	guard    *guard.Policy
	model    model.Model
	store    *store.MemoryStore
	browser  *browser.BrowserControl
	config   *core.Config
}

func NewApp() *App {
	return &App{
		registry: tools.NewRegistry(),
		guard:    guard.NewPolicy(),
		model:    model.NewMockModel(),
		store:    store.NewMemoryStore(),
		browser:  browser.NewBrowserControl(),
		config:   &core.Config{MaxCases: 100},
	}
}

func (a *App) SetConfig(config core.Config) error {
	a.config = &config
	a.store.SetConfig(&config)

	switch config.AIProvider {
	case "openai", "claude", "deepseek", "doubao", "ollama", "vllm":
		a.model = model.NewMockModel()
	default:
		a.model = model.NewMockModel()
	}

	a.guard.SetAllowedTools(a.registry.ListTools())

	return nil
}

func (a *App) GetConfig() core.Config {
	if a.config == nil {
		return core.Config{MaxCases: 100}
	}
	return *a.config
}

func (a *App) StartTest(goal string) error {
	status := a.store.GetStatus()
	if status.State == "running" {
		return fmt.Errorf("test already running")
	}

	a.store.Reset()

	engine := agentruntime.NewEngine(a.registry, a.guard, a.model, a.store)

	go func() {
		_ = engine.Run(context.Background(), goal, func(event string, data interface{}) {
			log.Printf("Event: %s, Data: %v", event, data)
		})
	}()

	return nil
}

func (a *App) StopTest() error {
	status := a.store.GetStatus()
	if status.State != "running" {
		return fmt.Errorf("no test running")
	}
	a.store.SetStatus(&core.TestStatus{State: "stopped"})
	return nil
}

func (a *App) GetTestStatus() core.TestStatus {
	return *a.store.GetStatus()
}

func (a *App) GetSteps() []*store.StepRecord {
	return a.store.GetSteps()
}

func (a *App) GetTestCases() []*core.TestCase {
	return a.store.GetTestCases()
}

func (a *App) GetLastScreenshot() []byte {
	screenshot, err := a.browser.Screenshot(context.Background())
	if err != nil {
		return nil
	}
	return screenshot
}

func (a *App) RegisterTools() {
	browserTool := tools.NewBrowserTool(a.browser)
	a.registry.Register(browserTool)
}

func (a *App) OnStartup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) OnDomReady(ctx context.Context) {
	a.RegisterTools()
}

func (a *App) OnBeforeClose(ctx context.Context) bool {
	a.browser.Close()
	return false
}

func main() {
	app := NewApp()

	err := wails.Run(&options.App{
		Title:  "TestFlow - AI Automated Testing Tool",
		Width:  1400,
		Height: 900,
		AssetServer: &assetserver.Options{
			Assets: testflow.Assets,
		},
		OnStartup:        app.OnStartup,
		OnDomReady:       app.OnDomReady,
		OnBeforeClose:    app.OnBeforeClose,
		BackgroundColour: &options.RGBA{R: 255, G: 255, B: 255, A: 255},
	})
	if err != nil {
		log.Fatal(err)
	}
}
