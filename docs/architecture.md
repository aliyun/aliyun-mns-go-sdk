# Aliyun Mns Go Sdk — 核心架构文档

<!-- 范围: 模块结构 + 包路径 + 边界规则 + 数据库 schema;运行时部署形态见 ../AGENTS.md「全局架构地图」+ workspace docs/architecture/01-components-dependencies.md -->

## 核心架构

### 请求处理流程

<arch-request-flow>
<!-- 推断源: 入口类源码(如 @SpringBootApplication / main 函数)+ controller / service 层调用图 -->

### 核心组件

| 组件 | 职责 |
|------|------|
| <component-1-name> | <component-1-role> |
| <component-2-name> | <component-2-role> |
| <component-3-name> | <component-3-role> |

## 核心包结构
<!-- 形态: 表格,列出每个 module / 包 + 一句话职责;不放具体类列表(那是 IDE 的事) -->

<package-structure-tables>

## 模块边界规则
<!-- 形态: 表格 + 关键原则列表;sdd=true 仓的硬约束在 openspec/constitution.md(本段引用即可) -->

各模块有明确的职责边界,新增代码必须放在正确的模块中:

| 代码类型 | 应放置的模块 | 说明 |
|---------|------------|------|
| <module-rule-1> | <module-target-1> | <module-rule-1-detail> |
| <module-rule-2> | <module-target-2> | <module-rule-2-detail> |
| <module-rule-3> | <module-target-3> | <module-rule-3-detail> |

**关键原则**:
- <key-principle-1>
- <key-principle-2>
- <key-principle-3>

## 与其他 MNS 仓的关系
<!-- 形态: 1-3 段,描述 HTTP/DB/Maven/Helm 等运行时耦合;workspace 跨仓依赖图见 ../../docs/architecture/01-components-dependencies.md -->

<inter-repo-relationships>

## 外部依赖服务

| 服务 | 用途 | 集成模块 |
|------|------|----------|
| <ext-svc-1> | <ext-svc-1-purpose> | <ext-svc-1-module> |
| <ext-svc-2> | <ext-svc-2-purpose> | <ext-svc-2-module> |

## 数据库
<!-- 仅适用 backend/ops 类有 DB 的仓;test/infra/sdk 可删除整段 -->

- **引擎**:<db-engine>
- **ORM**:<db-orm>
- **核心表**:<db-core-tables>
