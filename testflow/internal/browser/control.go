package browser

import (
	"context"
	"fmt"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/rod/lib/proto"

	"testflow/internal/core"
)

type BrowserControl struct {
	browser *rod.Browser
	page    *rod.Page
}

func NewBrowserControl() *BrowserControl {
	return &BrowserControl{}
}

func (b *BrowserControl) Launch(ctx context.Context) error {
	u := launcher.New().Headless(true).MustLaunch()
	b.browser = rod.New().ControlURL(u).MustConnect()
	return nil
}

func (b *BrowserControl) Close() error {
	if b.page != nil {
		b.page.Close()
	}
	if b.browser != nil {
		return b.browser.Close()
	}
	return nil
}

func (b *BrowserControl) Navigate(ctx context.Context, url string) error {
	if b.browser == nil {
		if err := b.Launch(ctx); err != nil {
			return err
		}
	}
	var err error
	b.page, err = b.browser.Page(proto.TargetCreateTarget{URL: "about:blank"})
	if err != nil {
		return err
	}
	b.page = b.page.Timeout(30 * time.Second)
	return b.page.Navigate(url)
}

func (b *BrowserControl) Screenshot(ctx context.Context) ([]byte, error) {
	if b.page == nil {
		return nil, fmt.Errorf("page not initialized")
	}
	return b.page.Screenshot(true, &proto.PageCaptureScreenshot{Format: proto.PageCaptureScreenshotFormatPng})
}

func (b *BrowserControl) GetDOM(ctx context.Context) ([]*core.Element, error) {
	if b.page == nil {
		return nil, fmt.Errorf("page not initialized")
	}

	script := `
	(() => {
		const elements = [];
		const interactives = document.querySelectorAll('button, input, select, textarea, a, [onclick], [role="button"]');
		interactives.forEach((el, index) => {
			elements.push({
				tag: el.tagName.toLowerCase(),
				id: el.id || '',
				class: el.className || '',
				text: el.innerText || el.textContent || '',
				type: el.type || '',
				placeholder: el.placeholder || '',
				xpath: '//' + el.tagName.toLowerCase() + '[' + (el.id ? '@id="' + el.id + '"' : 'position()=' + (index + 1)) + ']'
			});
		});
		return elements;
	})()
	`

	result, err := b.page.Evaluate(&rod.EvalOptions{JS: script})
	if err != nil {
		return nil, err
	}

	elements := make([]*core.Element, 0)
	if result != nil {
		items := result.Value.Arr()
		for _, item := range items {
			m := item.Map()
			el := &core.Element{}
			if v, ok := m["tag"]; ok {
				el.Tag = v.Str()
			}
			if v, ok := m["id"]; ok {
				el.ID = v.Str()
			}
			if v, ok := m["class"]; ok {
				el.Class = v.Str()
			}
			if v, ok := m["text"]; ok {
				el.Text = v.Str()
			}
			if v, ok := m["type"]; ok {
				el.Type = v.Str()
			}
			if v, ok := m["placeholder"]; ok {
				el.Placeholder = v.Str()
			}
			if v, ok := m["xpath"]; ok {
				el.XPath = v.Str()
			}
			elements = append(elements, el)
		}
	}

	return elements, nil
}

func (b *BrowserControl) Click(ctx context.Context, selector string) error {
	if b.page == nil {
		return fmt.Errorf("page not initialized")
	}
	el, err := b.page.Search(selector)
	if err != nil {
		return err
	}
	return el.First.Click(proto.InputMouseButtonLeft, 1)
}

func (b *BrowserControl) Input(ctx context.Context, selector, text string) error {
	if b.page == nil {
		return fmt.Errorf("page not initialized")
	}
	el, err := b.page.Search(selector)
	if err != nil {
		return err
	}
	return el.First.Input(text)
}

func (b *BrowserControl) Select(ctx context.Context, selector, value string) error {
	if b.page == nil {
		return fmt.Errorf("page not initialized")
	}
	el, err := b.page.Search(selector)
	if err != nil {
		return err
	}
	return el.First.Select([]string{value}, true, "css-selector")
}

func (b *BrowserControl) Scroll(ctx context.Context, selector string, delta int) error {
	if b.page == nil {
		return fmt.Errorf("page not initialized")
	}
	if selector != "" {
		el, err := b.page.Search(selector)
		if err != nil {
			return err
		}
		return el.First.ScrollIntoView()
	}
	b.page.Mouse.Scroll(0, float64(delta), 10)
	return nil
}

func (b *BrowserControl) Hover(ctx context.Context, selector string) error {
	if b.page == nil {
		return fmt.Errorf("page not initialized")
	}
	el, err := b.page.Search(selector)
	if err != nil {
		return err
	}
	return el.First.Hover()
}

func (b *BrowserControl) Wait(ctx context.Context, timeout int) error {
	if b.page == nil {
		return fmt.Errorf("page not initialized")
	}
	time.Sleep(time.Duration(timeout) * time.Second)
	return nil
}

func (b *BrowserControl) GetPageURL() string {
	if b.page == nil {
		return ""
	}
	info, err := b.page.Info()
	if err != nil {
		return ""
	}
	return info.URL
}
