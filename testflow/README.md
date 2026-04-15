# TestFlow - AI 自动化测试工具

## 项目简介

TestFlow 是一款 AI 驱动的全自动软件功能测试工具。其核心设计理念是：让大模型充当测试工程师，自主完成"看页面 → 做判断 → 执行操作 → 记录结果"的完整闭环。

## 技术栈

| 组件 | 选型 | 说明 |
|------|------|------|
| 后端语言 | Go | 编译为原生二进制，无运行时依赖 |
| 桌面框架 | Wails v2 | Go 生态的 Electron 替代品 |
| 前端框架 | Vue 3 + Element Plus | 企业级 UI 方案 |
| 浏览器控制 | go-rod | 纯 Go 编写的 Chrome DevTools Protocol 客户端 |
| AI 协议 | OpenAI Chat Completions | 主流 AI 提供商兼容 |

## 功能特性

- **零脚本测试**：不需要编写测试代码
- **AI 自主决策**：大模型自主规划测试路径
- **多模态感知**：截图 + DOM 双通道理解页面
- **实时可视化**：浏览器画面实时同步
- **标准化输出**：自动生成符合 GB/T 25000 标准的测试报告
- **多 AI 支持**：OpenAI、Claude、DeepSeek、豆包、Ollama、vLLM 等

## 项目结构

```
testflow/
├── cmd/testflow/
│   └── main.go              # Wails 应用入口
├── internal/
│   ├── browser/             # 浏览器控制层
│   │   └── control.go      # go-rod 封装
│   ├── core/
│   │   └── types.go        # 核心类型定义
│   ├── guard/
│   │   └── policy.go       # 防护策略
│   ├── model/
│   │   └── interface.go    # AI 模型接口
│   ├── runtime/
│   │   └── engine.go       # Agent Runtime 引擎
│   ├── store/
│   │   └── memory.go       # 内存存储
│   └── tools/
│       ├── registry.go     # 工具注册表
│       └── browser_tool.go # 浏览器工具
├── frontend/
│   ├── src/
│   │   ├── App.vue         # 根组件（三栏布局）
│   │   └── components/
│   │       ├── LeftPanel.vue   # 左栏：配置 + AI 响应
│   │       ├── CenterPanel.vue  # 中栏：浏览器截图
│   │       └── RightPanel.vue   # 右栏：进度 + 用例
│   └── dist/               # 前端构建产物
├── go.mod
└── README.md
```

## 构建

### 前置依赖

- Go 1.22+
- Node.js 18+ (for frontend)
- Chrome/Chromium (for go-rod)

### 构建命令

```bash
# 下载依赖
cd testflow
go mod tidy

# 构建前端 (需要先安装 npm 依赖)
cd frontend
npm install
npm run build

# 返回根目录并构建应用
cd ..
go build -o testflow ./cmd/testflow
```

### 运行

```bash
./testflow
```

## 开发说明

### Phase 1 (当前)

- [x] 项目初始化 (Wails v2 + Vue 3)
- [x] 核心类型定义
- [x] Tool Registry 机制
- [x] Browser Control 浏览器控制
- [x] Guardrail 防护机制
- [x] Memory Store 存储层
- [x] AI Model 接口层
- [x] Wails 应用入口
- [x] UI 三栏布局
- [x] 配置表单

### Phase 2 (规划中)

- [ ] AI 测试循环引擎
- [ ] 多模态页面感知
- [ ] 浏览器操作引擎 (7种操作)
- [ ] 对话历史管理

### Phase 3 (规划中)

- [ ] 测试用例生成逻辑
- [ ] Excel 报告导出

### Phase 4 (规划中)

- [ ] 多 AI 提供商适配
- [ ] 端到端测试验证
- [ ] 性能优化

## 核心模块

### Agent Runtime

```
Decide → Guard → Execute → Observe → Judge
```

1. **Decide**: AI 模型分析当前状态，决定下一步操作
2. **Guard**: 防护机制检查操作是否安全
3. **Execute**: 执行工具操作
4. **Observe**: 观察执行结果
5. **Judge**: 判断是否满足终止条件

### 浏览器操作

- `click` - 点击元素
- `input` - 输入文本
- `select` - 选择下拉项
- `navigate` - 页面跳转
- `scroll` - 页面滚动
- `hover` - 鼠标悬停
- `wait` - 等待加载

## License

MIT
