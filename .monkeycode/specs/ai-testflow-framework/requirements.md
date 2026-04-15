# Requirements Document

## Introduction

本项目旨在开发 TestFlow——一款 AI 驱动的全自动软件功能测试工具。核心目标是让大模型充当测试工程师，自主完成"看页面 → 做判断 → 执行操作 → 记录结果"的完整闭环。

## Glossary

- **TestFlow**: AI 驱动的自动化功能测试工具
- **Agent Runtime**: Agent 主循环，负责决策、执行、观察的闭环控制
- **Tool Registry**: 工具注册表，统一管理所有可调用工具
- **Guardrail**: 防护机制，约束工具调用范围
- **Multi-modal Perception**: 多模态感知，同时支持截图和 DOM 信息理解
- **Test Case**: 测试用例，记录测试步骤和预期结果

## Requirements

### Requirement 1: 技术栈选型

**User Story:** AS 开发者，我需要选择成熟且高效的技术栈，以便快速构建可靠的桌面应用。

#### Acceptance Criteria

1. The system SHALL use Go 作为后端开发语言，利用其原生二进制编译和并发处理能力
2. The system SHALL use Wails v2 作为桌面框架，生成单文件 EXE（~15MB）
3. The system SHALL use Vue 3 + Element Plus 作为前端 UI 方案
4. The system SHALL use go-rod 作为浏览器自动化控制库

### Requirement 2: 整体架构设计

**User Story:** AS 开发者，我需要清晰的模块划分，以便后续扩展和维护。

#### Acceptance Criteria

1. The system SHALL 采用三层架构：UI 层（Vue 3）、业务逻辑层（Go）、浏览器控制层（go-rod）
2. The system SHALL 支持前后端通过 Wails 方法绑定直接通信，而非 HTTP
3. The system SHALL 支持通过 Wails EventsEmit 从 Go 端推送实时数据到前端

### Requirement 3: UI 布局设计

**User Story:** AS 测试工程师，我需要一个清晰直观的界面，方便监控测试进度和结果。

#### Acceptance Criteria

1. The system SHALL 采用三栏布局：左栏（系统信息 + AI 响应流）、中栏（浏览器实时截图 + 操作日志）、右栏（测试进度 + 用例列表）
2. The system SHALL 支持用户拖拽调整三栏宽度
3. The system SHALL 在左栏实时展示 AI 响应的 JSON 格式化输出
4. The system SHALL 在中栏显示浏览器画面的实时截图
5. The system SHALL 在右栏按模块分组展示测试用例

### Requirement 4: 系统配置功能

**User Story:** AS 测试工程师，我需要配置被测系统信息和 AI 模型参数。

#### Acceptance Criteria

1. The system SHALL 支持配置被测系统的 URL
2. The system SHALL 支持配置登录账号密码（可选）
3. The system SHALL 支持选择 AI 模型提供商（OpenAI、Claude、DeepSeek、豆包、Ollama、vLLM）
4. The system SHALL 支持配置 API Key 和 API Endpoint
5. The system SHALL 支持本地私有化部署（Ollama、vLLM）无需 API Key

### Requirement 5: 提示词管理

**User Story:** AS 测试工程师，我需要管理和定制 AI 的测试行为提示词。

#### Acceptance Criteria

1. The system SHALL 提供默认系统提示词，赋予 AI"资深测试工程师"角色
2. The system SHALL 支持查看和编辑默认提示词
3. The system SHALL 支持创建针对特定系统的自定义提示词
4. The system SHALL 支持 AI 自动生成定制化提示词（根据系统名称和 URL 推测系统类型）

### Requirement 6: AI 测试循环引擎

**User Story:** AS 测试工程师，我需要 AI 自动执行完整的测试闭环。

#### Acceptance Criteria

1. WHEN 测试启动时，The system SHALL 执行页面感知步骤，截取浏览器截图并提取 DOM 可交互元素
2. WHEN 页面感知完成，The system SHALL 将截图和 DOM 数据发送给多模态大模型
3. WHEN AI 返回决策，The system SHALL 解析 JSON 格式响应，包含操作指令和测试用例
4. WHEN 操作指令存在，The system SHALL 在浏览器中执行对应操作（click、input、select、navigate、scroll、hover、wait）
5. WHEN AI 判断当前操作构成完整测试场景，The system SHALL 自动生成标准化测试用例
6. The system SHALL 检查终止条件（AI 声明完成 / 达到用例上限 / 用户手动停止），不满足则继续循环

### Requirement 7: 多模态页面感知

**User Story:** AS 测试工程师，我需要系统能够准确理解页面结构和内容。

#### Acceptance Criteria

1. The system SHALL 通过注入 JavaScript 提取页面所有可交互元素（按钮、输入框、链接等）的属性（tag、id、class、text、type、placeholder）
2. The system SHALL 将提取结果组成结构化 JSON 数组供 AI 分析
3. WHEN DOM 可交互元素 ≥ 10 个，The system SHALL 仅使用 DOM 数据进行快速测试
4. WHEN DOM 元素 < 10 个 或 连续 DOM 模式已达 5 轮，The system SHALL 自动加入截图辅助理解
5. The system SHALL 检测模型是否支持视觉能力，不支持时自动降级为纯文本 DOM 模式

### Requirement 8: 浏览器操作引擎

**User Story:** AS 测试工程师，我需要系统能够执行各种浏览器操作。

#### Acceptance Criteria

1. The system SHALL 支持 click 操作，点击指定元素
2. The system SHALL 支持 input 操作，向输入框填写文本
3. The system SHALL 支持 select 操作，选择下拉框选项
4. The system SHALL 支持 navigate 操作，跳转到指定 URL
5. The system SHALL 支持 scroll 操作，滚动页面
6. The system SHALL 支持 hover 操作，鼠标悬停
7. The system SHALL 支持 wait 操作，等待页面加载
8. The system SHALL 透明兼容 jQuery 风格的 :contains() 选择器

### Requirement 9: 对话历史管理

**User Story:** AS 测试工程师，我需要系统管理长对话上下文，避免超出 Token 限制。

#### Acceptance Criteria

1. WHEN 消息历史超过 42 条，The system SHALL 自动裁剪，保留最近 40 条
2. The system SHALL 每轮交互将已测功能摘要注入提示词，避免重复测试
3. The system SHALL 在内存中维护 testedKeys 去重 Map，实现双重防重复机制

### Requirement 10: 测试用例管理

**User Story:** AS 测试工程师，我需要系统自动生成标准化的测试用例。

#### Acceptance Criteria

1. The system SHALL 生成符合 GB/T 25000 标准的测试用例
2. The system SHALL 为每个测试用例记录：编号、模块、名称、前置条件、步骤、预期结果、实际结果、状态、优先级、测试方法
3. The system SHALL 支持按模块分组展示测试用例
4. The system SHALL 支持一键导出 Excel 测试报告

### Requirement 11: 测试报告导出

**User Story:** AS 测试工程师，我需要导出专业的测试报告。

#### Acceptance Criteria

1. The system SHALL 支持导出 Excel 格式测试报告
2. The system SHALL 在报告中包含测试汇总页：总用例数、通过/失败/阻塞统计
3. The system SHALL 按模块分 Sheet，每个功能模块独立一个工作表
4. The system SHALL 对状态进行着色：通过为绿色、失败为红色

## 项目阶段规划

### Phase 1: 框架与 UI 设计（当前阶段）

- 技术栈选型与项目初始化
- 整体架构设计
- UI 三栏布局实现
- 前后端通信机制实现
- 系统配置界面实现
- 提示词管理界面实现

### Phase 2: 核心引擎开发

- AI 测试循环引擎
- 多模态页面感知系统
- 浏览器操作引擎
- 对话历史管理

### Phase 3: 测试用例与报告

- 测试用例生成逻辑
- 测试报告导出功能

### Phase 4: 集成与优化

- 多 AI 提供商适配
- 端到端测试验证
- 性能优化与错误处理