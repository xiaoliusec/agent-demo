# AI TestFlow Framework Design

Feature Name: ai-testflow-framework
Updated: 2026-04-15

## Description

TestFlow 是一款 AI 驱动的全自动软件功能测试工具。其核心设计理念是：让大模型充当测试工程师，自主完成"看页面 → 做判断 → 执行操作 → 记录结果"的完整闭环。

## Architecture

### 整体架构图

```mermaid
graph TB
    subgraph Frontend["前端 UI 层 (Vue 3 + Element Plus)"]
        UI["三栏布局界面<br/>- 左：配置 + AI 响应流<br/>- 中：浏览器截图 + 日志<br/>- 右：进度 + 用例列表"]
    end

    subgraph Backend["后端业务逻辑层 (Go)"]
        Wails["Wails 方法绑定<br/>+ EventsEmit 推送"]
        Runtime["Agent Runtime 引擎<br/>Decide → Guard → Execute → Observe → Judge"]
        Model["AI 模型接口<br/>OpenAI Chat Completions 协议"]
        Store["步骤记录存储<br/>Memory Store"]
    end

    subgraph Browser["浏览器控制层 (go-rod)"]
        Browser["Chromium 实例<br/>+ WebSocket 连接"]
        DOM["DOM 提取<br/>JavaScript 注入"]
        Screenshot["页面截图<br/>实时同步"]
    end

    Frontend <-->|Wails RPC| Backend
    Backend -->|go-rod 控制| Browser
    Backend <-->|HTTP 调用| AIProvider["AI 提供商<br/>OpenAI/Claude/DeepSeek/..."]
```

### 技术选型理由

| 组件 | 选型 | 理由 |
|------|------|------|
| 后端语言 | Go | 编译为原生二进制，无运行时依赖；goroutine 天然适合并发处理 AI 流式响应 + 浏览器控制 |
| 桌面框架 | Wails v2 | Go 生态的 Electron 替代品，生成的 EXE 体积仅 ~15MB（Electron 动辄 150MB+） |
| 前端框架 | Vue 3 + Element Plus | 成熟的企业级 UI 方案，三栏布局开发效率高 |
| 浏览器控制 | go-rod | 纯 Go 编写的 Chrome DevTools Protocol 客户端，支持Headless浏览器控制 |
| AI 协议 | OpenAI Chat Completions | 主流 AI 提供商（包括国产 DeepSeek、豆包）都兼容这套协议 |

### 项目目录结构

```
testflow/
├── cmd/
│   └── testflow/
│       └── main.go              # Wails 应用入口
├── internal/
│   ├── browser/                  # 浏览器控制层
│   │   ├── control.go           # 浏览器启动、关闭控制
│   │   ├── dom.go               # DOM 提取逻辑
│   │   ├── screenshot.go        # 页面截图
│   │   └── actions.go           # 7种浏览器操作实现
│   ├── runtime/                  # Agent 运行时引擎
│   │   ├── engine.go            # 主循环：Decide → Guard → Execute → Observe → Judge
│   │   ├── state.go             # 状态管理
│   │   └── loop.go              # 循环控制与终止判断
│   ├── model/                    # AI 模型接口
│   │   ├── interface.go         # 统一接口定义
│   │   ├── openai.go            # OpenAI 实现
│   │   ├── claude.go            # Claude 实现
│   │   └── adaptive.go          # 多提供商适配
│   ├── perception/                # 多模态感知
│   │   ├── page.go             # 页面感知（截图 + DOM）
│   │   └── smartmode.go         # 智能模式切换
│   ├── tools/                    # 工具协议与实现
│   │   ├── registry.go         # 工具注册表
│   │   ├── click.go            # click 工具
│   │   ├── input.go            # input 工具
│   │   ├── select.go           # select 工具
│   │   ├── navigate.go         # navigate 工具
│   │   ├── scroll.go           # scroll 工具
│   │   ├── hover.go            # hover 工具
│   │   └── wait.go             # wait 工具
│   ├── guard/                    # 防护机制
│   │   └── policy.go           # 策略定义与检查
│   ├── store/                    # 存储层
│   │   ├── memory.go           # 内存存储
│   │   └── types.go            # 存储类型
│   ├── prompt/                   # 提示词管理
│   │   ├── default.go          # 默认提示词
│   │   └── generator.go        # AI 生成提示词
│   ├── history/                  # 对话历史管理
│   │   ├── manager.go          # 历史记录管理
│   │   └── context.go          # 上下文裁剪
│   └── testcase/                 # 测试用例生成
│       ├── generator.go        # 用例生成
│       └── exporter.go         # Excel 导出
├── frontend/
│   ├── src/
│   │   ├── App.vue             # 根组件
│   │   ├── main.ts             # 前端入口
│   │   ├── components/
│   │   │   ├── LeftPanel.vue   # 左栏：配置 + AI 响应
│   │   │   ├── CenterPanel.vue # 中栏：浏览器截图
│   │   │   ├── RightPanel.vue  # 右栏：进度 + 用例
│   │   │   └── ConfigForm.vue  # 配置表单
│   │   ├── views/
│   │   │   ├── MainView.vue    # 主界面
│   │   │   └── PromptView.vue  # 提示词管理
│   │   └── stores/
│   │       ├── config.ts       # 配置状态
│   │       └── test.ts         # 测试状态
│   ├── package.json
│   └── vite.config.ts
├── wails.json                   # Wails 项目配置
├── go.mod
└── README.md
```

## Components and Interfaces

### 1. Browser Control Layer

#### BrowserControl

```go
type BrowserControl interface {
    Launch(ctx context.Context) error           // 启动浏览器
    Close() error                               // 关闭浏览器
    Navigate(ctx context.Context, url string) error  // 导航到 URL
    Screenshot(ctx context.Context) ([]byte, error)   // 获取截图
    GetDOM(ctx context.Context) ([]Element, error)    // 获取 DOM 元素
    Click(ctx context.Context, selector string) error
    Input(ctx context.Context, selector, text string) error
    Select(ctx context.Context, selector, value string) error
    Scroll(ctx context.Context, selector string, delta int) error
    Hover(ctx context.Context, selector string) error
    Wait(ctx context.Context, timeout int) error
}
```

#### Element

```go
type Element struct {
    Tag         string `json:"tag"`
    ID          string `json:"id"`
    Class       string `json:"class"`
    Text        string `json:"text"`
    Type        string `json:"type"`
    Placeholder string `json:"placeholder"`
    XPath       string `json:"xpath"`
}
```

### 2. Agent Runtime

#### Decision

```go
type Decision struct {
    ThoughtSummary string     `json:"thought_summary"`  // 本轮思路
    NextAction     *Action     `json:"next_action"`       // 下一步动作
    Done           bool        `json:"done"`               // 是否完成
    FinalResult    string      `json:"final_result"`       // 最终结果
    TestCase       *TestCase   `json:"test_case"`          // 生成的测试用例
}
```

#### Action

```go
type Action struct {
    ToolName string            `json:"tool_name"`  // 工具名
    Args     map[string]string `json:"args"`       // 参数
    Reason   string            `json:"reason"`     // 执行原因
}
```

#### TestCase

```go
type TestCase struct {
    ID            string   `json:"id"`             // 用例编号
    Module        string   `json:"module"`         // 所属模块
    Name          string   `json:"name"`           // 用例名称
    Precondition  string   `json:"precondition"`   // 前置条件
    Steps         []string `json:"steps"`          // 测试步骤
    ExpectedResult string  `json:"expected_result"`// 预期结果
    ActualResult  string   `json:"actual_result"`  // 实际结果
    Status        string   `json:"status"`          // 状态：passed/failed/blocked
    Priority      string   `json:"priority"`        // 优先级
    Method        string   `json:"method"`          // 测试方法
}
```

### 3. AI Model Interface

```go
type Model interface {
    Decide(ctx context.Context, state StateView) (Decision, error)
}

type StateView struct {
    Goal         string           // 用户目标
    StepCount    int              // 当前步数
    LastOutput   string           // 上一步输出
    Screenshot   []byte           // 当前截图
    DOMElements  []Element        // DOM 元素列表
    TestCases    []*TestCase      // 已生成用例
    History      []*HistoryEntry  // 对话历史
    TestedSummary []string        // 已测功能摘要
}
```

### 4. Frontend-Backend Communication

#### Wails 绑定方法

```go
// Frontend 调用后端的方法
type App struct {
    // 配置相关
    SetConfig(config Config) error
    GetConfig() Config

    // 测试控制
    StartTest(goal string) error
    StopTest() error
    GetTestStatus() TestStatus

    // 实时数据订阅
    OnAIToken(callback func(string))   // AI 流式响应
    OnScreenshot(callback func([]byte)) // 截图更新
    OnTestCase(callback func(*TestCase)) // 新用例生成
    OnLog(callback func(string))        // 日志推送
}
```

## Data Models

### Config

```go
type Config struct {
    TargetURL    string `json:"target_url"`    // 被测系统 URL
    Username     string `json:"username"`     // 登录账号
    Password     string `json:"password"`      // 登录密码
    AIProvider   string `json:"ai_provider"`  // AI 提供商
    APIKey       string `json:"api_key"`      // API Key
    APIEndpoint  string `json:"api_endpoint"` // API Endpoint
    ModelName    string `json:"model_name"`   // 模型名称
    MaxCases     int    `json:"max_cases"`    // 最大用例数
}
```

### TestStatus

```go
type TestStatus struct {
    State       string      `json:"state"`        // idle/running/paused/completed/stopped
    CurrentStep int         `json:"current_step"`  // 当前步数
    TotalCases  int         `json:"total_cases"`   // 已生成用例数
    PassedCases int         `json:"passed_cases"`  // 通过用例数
    FailedCases int         `json:"failed_cases"`  // 失败用例数
    CurrentURL  string      `json:"current_url"`   // 当前页面 URL
    LastError   string     `json:"last_error"`    // 最近错误
}
```

## Correctness Properties

### 终止条件

测试循环在满足以下任一条件时终止：

1. AI 显式声明 `done: true`
2. 生成用例数达到配置上限 (`MaxCases`)
3. 用户手动停止测试
4. 发生不可恢复错误

### 防护策略

1. **工具调用约束**: 仅允许已注册的工具被调用
2. **操作安全检查**: 禁止执行危险操作（如删除文件）
3. **超时控制**: 每个操作设置最大执行时间
4. **重试机制**: 失败操作最多重试 3 次

### 智能模式切换

```go
// SmartMode 决策逻辑
func ShouldUseScreenshot(domElements []Element, consecutiveDOMRounds int) bool {
    if len(domElements) < 10 {
        return true  // DOM 元素稀少，需要截图辅助
    }
    if consecutiveDOMRounds >= 5 {
        return true  // 连续 DOM 模式已达 5 轮，切换截图
    }
    return false
}
```

## Error Handling

| 错误类型 | 处理策略 | 用户提示 |
|----------|----------|----------|
| 浏览器启动失败 | 重试 3 次，提示检查 Chrome 安装 | "浏览器启动失败，请检查 Chrome 是否安装" |
| AI API 调用失败 | 重试 3 次，指数退避 | "AI 服务暂时不可用，正在重试..." |
| 元素定位失败 | 记录错误，继续下一步 | "无法定位元素，跳过当前操作" |
| 页面加载超时 | 重试或跳过 | "页面加载超时" |
| Token 超限 | 触发上下文裁剪 | 自动裁剪对话历史 |

## Test Strategy

### 单元测试

- BrowserControl 各方法测试
- Agent Runtime 循环逻辑测试
- AI Response JSON 解析测试
- 提示词模板渲染测试

### 集成测试

- 前后端 Wails 通信测试
- AI 模型实际调用测试（需配置真实 API Key）
- 端到端测试流程验证

### 测试覆盖目标

- 核心业务逻辑覆盖率 > 80%
- 边界条件和错误处理覆盖率 > 70%