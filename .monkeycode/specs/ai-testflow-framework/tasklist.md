# Task List: AI TestFlow Framework - Phase 1

## Phase 1: 框架与 UI 设计

### 1.1 项目初始化

- [x] 1.1.1 初始化 Wails v2 项目结构
- [x] 1.1.2 配置 go.mod 依赖（wails、go-rod 等）
- [x] 1.1.3 配置前端 Vue 3 + Element Plus
- [x] 1.1.4 验证项目能正常编译运行

### 1.2 后端核心类型定义

- [x] 1.2.1 定义 Config 配置结构体
- [x] 1.2.2 定义 TestStatus 测试状态结构体
- [x] 1.2.3 定义 Element DOM元素结构体
- [x] 1.2.4 定义 Action 工具调用结构体
- [x] 1.2.5 定义 Decision AI决策结构体
- [x] 1.2.6 定义 TestCase 测试用例结构体
- [x] 1.2.7 定义 StateView 状态视图结构体

### 1.3 Tool Registry 工具注册机制

- [x] 1.3.1 实现 Registry 结构体和注册方法
- [x] 1.3.2 实现工具查找和执行方法
- [x] 1.3.3 预留 7 种浏览器工具接口

### 1.4 Browser Control 浏览器控制层

- [x] 1.4.1 实现 BrowserControl 接口定义
- [x] 1.4.2 实现浏览器启动和关闭
- [x] 1.4.3 实现 Navigate 导航功能
- [x] 1.4.4 实现 Screenshot 截图功能
- [x] 1.4.5 实现 GetDOM DOM提取功能

### 1.5 Guardrail 防护机制

- [x] 1.5.1 定义防护策略结构
- [x] 1.5.2 实现工具白名单检查
- [x] 1.5.3 实现危险操作拦截

### 1.6 Store 存储层

- [x] 1.6.1 扩展 MemoryStore 支持 TestFlow 数据
- [x] 1.6.2 实现步骤记录存储
- [x] 1.6.3 实现测试用例存储

### 1.7 AI Model 接口层

- [x] 1.7.1 定义 Model 接口
- [x] 1.7.2 定义 StateView 状态视图
- [x] 1.7.3 实现 MockModel 用于测试

### 1.8 Wails 应用入口

- [x] 1.8.1 实现 App 后端结构体
- [x] 1.8.2 实现 SetConfig/GetConfig 配置方法
- [x] 1.8.3 实现 StartTest/StopTest 测试控制方法
- [x] 1.8.4 实现 GetTestStatus 状态查询方法
- [x] 1.8.5 实现 EventsEmit 事件推送

### 1.9 UI 三栏布局实现

- [x] 1.9.1 创建 LeftPanel 左栏组件
- [x] 1.9.2 创建 CenterPanel 中栏组件
- [x] 1.9.3 创建 RightPanel 右栏组件
- [x] 1.9.4 实现三栏拖拽调整宽度
- [x] 1.9.5 实现主布局整合

### 1.10 ConfigForm 配置表单

- [x] 1.10.1 实现 URL 输入框
- [x] 1.10.2 实现账号密码输入
- [x] 1.10.3 实现 AI 提供商选择
- [x] 1.10.4 实现 API Key/Endpoint 配置
- [x] 1.10.5 实现最大用例数配置

### 1.11 提示词管理界面

- [ ] 1.11.1 创建 PromptView 提示词管理页面
- [ ] 1.11.2 实现默认提示词展示
- [ ] 1.11.3 实现提示词编辑功能
- [ ] 1.11.4 实现自定义提示词创建

### 1.12 前端状态管理

- [ ] 1.12.1 创建 config store
- [ ] 1.12.2 创建 test store
- [ ] 1.12.3 实现状态持久化

### 1.13 前后端联调

- [ ] 1.13.1 验证 Wails RPC 通信
- [ ] 1.13.2 验证 EventsEmit 事件推送
- [ ] 1.13.3 验证配置读写功能

### 1.14 项目文档更新

- [x] 1.14.1 更新 README.md
- [ ] 1.14.2 更新 docs/FRAMEWORK.md
- [ ] 1.14.3 提交所有 Phase 1 变更

---

## Phase 2: 核心引擎开发（后续阶段）

- [ ] 2.1 AI 测试循环引擎实现
- [ ] 2.2 多模态页面感知系统
- [ ] 2.3 浏览器操作引擎（7种操作）
- [ ] 2.4 对话历史管理

## Phase 3: 测试用例与报告（后续阶段）

- [ ] 3.1 测试用例生成逻辑
- [ ] 3.2 Excel 报告导出功能

## Phase 4: 集成与优化（后续阶段）

- [ ] 4.1 多 AI 提供商适配
- [ ] 4.2 端到端测试验证
- [ ] 4.3 性能优化与错误处理