# 终端智能体框架设计（Demo）

## 1. 目标

构建一个最小可运行框架，验证以下能力：

1. 任务目标能够被自动拆解为步骤
2. 每一步通过统一工具协议执行
3. 执行结果回流到下一轮决策
4. 满足终止条件后自动收敛并产出结果

## 2. 核心模块

- `Runtime Engine`
  - 主循环：`Decide -> Guard -> Execute -> Observe -> Judge`
  - 负责步数限制、失败计数、状态推进

- `Model`
  - 统一决策接口：输入当前状态，输出结构化决策
  - Demo 使用 `MockModel` 模拟多轮动作

- `Tool Registry`
  - 统一注册和调用工具
  - 每个工具实现 `Run(ctx, action)`

- `Guardrail`
  - 控制可执行工具范围
  - 拦截非法或未知工具调用

- `Store`
  - 记录每一步动作与结果
  - 为后续回放、审计、重跑预留基础

## 3. 数据流

1. 用户输入任务目标
2. Runtime 读取当前状态，调用 Model 获取 `Decision`
3. Guard 校验 `Action`
4. Registry 执行工具，返回 `ToolResult`
5. Runtime 记录步骤并更新状态
6. 到达终止条件后输出 `FinalResult`

## 4. 关键协议

- `Decision`
  - `thought_summary`: 本轮简要思路
  - `next_action`: 下一步工具调用
  - `done`: 是否完成
  - `final_result`: 完成时的总结

- `Action`
  - `tool_name`: 工具名
  - `args`: 参数键值
  - `reason`: 执行原因

- `ToolResult`
  - `ok`: 是否成功
  - `output`: 文本输出
  - `error`: 错误信息
  - `artifacts`: 产物路径映射

## 5. Demo 执行流程

MockModel 默认三步：

1. 调用 `workspace.scan` 扫描目录
2. 调用 `report.write` 输出 `output/demo-report.md`
3. 输出 `done=true` 并返回最终总结

## 6. 后续演进建议

1. 将 `MockModel` 替换为真实 LLM（OpenAI 协议）
2. 增加 `bash.run`、`git.status`、`browser.*` 工具
3. 引入 SQLite 持久化与断点续跑
4. 增加失败恢复策略和更细粒度 guard 规则
