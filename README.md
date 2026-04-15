# Agent Demo

这是一个可运行的终端智能体框架 Demo，用于验证“目标驱动 -> 工具调用 -> 结果观察 -> 自动收敛”的执行闭环。

## Demo 能力

- 运行统一的 Agent Runtime 循环
- 通过 Tool Registry 调度工具
- 通过 Guardrail 做工具调用约束
- 持久化步骤记录到内存存储
- 使用 MockModel 模拟智能决策

## 快速运行

```bash
go run ./cmd/agentd -goal "先扫描工作区，再输出一份 demo 报告"
```

运行后会生成文件：`output/demo-report.md`

## 目录结构

- `cmd/agentd/main.go`: 程序入口
- `internal/core`: 共享类型定义
- `internal/runtime`: Agent 主循环与状态管理
- `internal/model`: 模型接口与 MockModel
- `internal/tools`: 工具协议与工具实现
- `internal/store`: 步骤记录存储
- `internal/guard`: 执行策略与约束
- `docs/FRAMEWORK.md`: 框架设计说明
