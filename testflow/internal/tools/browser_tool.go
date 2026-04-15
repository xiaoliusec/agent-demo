package tools

import (
	"context"

	"testflow/internal/browser"
	"testflow/internal/core"
)

type BrowserTool struct {
	browser *browser.BrowserControl
}

func NewBrowserTool(b *browser.BrowserControl) *BrowserTool {
	return &BrowserTool{browser: b}
}

func (t *BrowserTool) Name() string {
	return "browser.control"
}

func (t *BrowserTool) Run(ctx context.Context, action core.Action) (core.ToolResult, error) {
	switch action.ToolName {
	case "browser.navigate":
		url := action.Args["url"]
		if err := t.browser.Navigate(ctx, url); err != nil {
			return core.ToolResult{OK: false, Error: err.Error()}, err
		}
		return core.ToolResult{OK: true, Output: "Navigated to " + url}, nil

	case "browser.screenshot":
		img, err := t.browser.Screenshot(ctx)
		if err != nil {
			return core.ToolResult{OK: false, Error: err.Error()}, err
		}
		return core.ToolResult{OK: true, Output: "Screenshot taken", Artifacts: map[string]string{"screenshot": "data:image/png;base64," + string(img)}}, nil

	case "browser.get_dom":
		_, err := t.browser.GetDOM(ctx)
		if err != nil {
			return core.ToolResult{OK: false, Error: err.Error()}, err
		}
		return core.ToolResult{OK: true, Output: "DOM extracted"}, nil

	case "browser.click":
		selector := action.Args["selector"]
		if err := t.browser.Click(ctx, selector); err != nil {
			return core.ToolResult{OK: false, Error: err.Error()}, err
		}
		return core.ToolResult{OK: true, Output: "Clicked " + selector}, nil

	case "browser.input":
		selector := action.Args["selector"]
		text := action.Args["text"]
		if err := t.browser.Input(ctx, selector, text); err != nil {
			return core.ToolResult{OK: false, Error: err.Error()}, err
		}
		return core.ToolResult{OK: true, Output: "Input " + text}, nil

	case "browser.select":
		selector := action.Args["selector"]
		value := action.Args["value"]
		if err := t.browser.Select(ctx, selector, value); err != nil {
			return core.ToolResult{OK: false, Error: err.Error()}, err
		}
		return core.ToolResult{OK: true, Output: "Selected " + value}, nil

	case "browser.scroll":
		selector := action.Args["selector"]
		delta := 100
		if _, ok := action.Args["delta"]; ok {
			delta = 100
		}
		if err := t.browser.Scroll(ctx, selector, delta); err != nil {
			return core.ToolResult{OK: false, Error: err.Error()}, err
		}
		return core.ToolResult{OK: true, Output: "Scrolled"}, nil

	case "browser.hover":
		selector := action.Args["selector"]
		if err := t.browser.Hover(ctx, selector); err != nil {
			return core.ToolResult{OK: false, Error: err.Error()}, err
		}
		return core.ToolResult{OK: true, Output: "Hovered " + selector}, nil

	case "browser.wait":
		timeout := 5
		if _, ok := action.Args["timeout"]; ok {
			timeout = 5
		}
		if err := t.browser.Wait(ctx, timeout); err != nil {
			return core.ToolResult{OK: false, Error: err.Error()}, err
		}
		return core.ToolResult{OK: true, Output: "Waited"}, nil

	default:
		return core.ToolResult{OK: false, Error: "unknown action: " + action.ToolName}, nil
	}
}
