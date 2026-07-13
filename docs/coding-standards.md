# Aliyun Mns Go Sdk — 编码规范与测试

<!-- 范围: 命名/日志/异常/测试规范 + Commit message 约定;sdd=true 仓的硬约束在 openspec/constitution.md(本段引用即可) -->

## 编码规范

### 命名规范
<!-- 形态: 列表 + 例子;不复述语言通用规范(Java naming convention),只列项目特化 -->

- <naming-rule-1>
- <naming-rule-2>
- <naming-rule-3>

### 日志使用

- 使用 <log-framework>,日志类:`<log-class-pattern>`
- <log-rule-1>
- <log-rule-2>
- **禁止**:<log-forbidden>

### 异常处理

- <exception-rule-1>
- <exception-rule-2>
- **禁止**:<exception-forbidden>

## 测试规范

### 测试框架

- <test-framework> v<test-framework-version>
- 测试目录:`<test-dir>`
- 覆盖率工具:<coverage-tool>(报告路径:`<coverage-report-path>`)

### 测试哲学

<test-philosophy>

### 测试模式

- **<test-pattern-1-name>**:<test-pattern-1-detail>
- **<test-pattern-2-name>**:<test-pattern-2-detail>

### 测试目录结构

```
<test-tree>
```

## Commit Message 约定
<!-- 形态: 列表或表格,列出 type 前缀 + 例子;不复述 conventional-commits 通用规范,只列项目特化 -->

| Type | 用途 | 例子 |
|------|------|------|
| `feat` | 新功能 | `feat: add CreateQueueV2 API` |
| `fix` | 缺陷修复 | `fix: race condition in pushQueue` |
| `refactor` | 重构(无功能变化) | `refactor: extract DiamondConfiguration bean` |
| `test` | 仅测试改动 | `test: add concurrent push test` |
| `docs` | 仅文档改动 | `docs: update AGENTS.md 常用命令` |
| `sdd(<phase>)` | SDD 流程产出 | `sdd(apply): A.2 templates` |

<commit-msg-additional-rules>
