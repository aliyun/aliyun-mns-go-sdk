# Aliyun Mns Go Sdk — 技术栈与配置文档

<!-- 范围: 技术栈版本表 + 依赖列表 + 配置体系 + 环境变量 + 外部服务清单;运行时业务逻辑见 architecture.md -->

## 技术栈
<!-- 形态: 表格,列出语言/框架/构建/测试/日志各 1-2 行;不复述 AGENTS.md「项目概览」 -->

| 类别 | 技术 | 版本 |
|------|------|------|
| 语言 | <language> | <language-version> |
| 框架 | <framework> | <framework-version> |
| 构建工具 | <build-tool> | <build-tool-version> |
| 测试框架 | <test-framework> | <test-framework-version> |
| 日志框架 | <log-framework> | <log-framework-version> |

## 核心依赖
<!-- 形态: 列表 + 一句话用途;按用途分类(数据访问 / 外部 SDK / 工具库 等) -->

### <dep-category-1>
- `<dep-1>` v<dep-1-version> — <dep-1-purpose>
- `<dep-2>` v<dep-2-version> — <dep-2-purpose>

### <dep-category-2>
- `<dep-3>` v<dep-3-version> — <dep-3-purpose>

## 配置体系

### 配置加载机制
<!-- 推断源: 入口类初始化代码 / @ConfigurationProperties / Diamond 集成代码 / .env 加载逻辑 -->

<config-loading-mechanism>

### 配置文件

| 文件 | 用途 |
|------|------|
| <config-file-1> | <config-file-1-purpose> |
| <config-file-2> | <config-file-2-purpose> |

### 环境变量
<!-- 推断源: docker/Dockerfile / supervisord.conf / 启动脚本 grep $XXX -->

| 变量 | 用途 | 来源 |
|------|------|------|
| <env-var-1> | <env-var-1-purpose> | <env-var-1-source> |
| <env-var-2> | <env-var-2-purpose> | <env-var-2-source> |

## 安全与认证
<!-- 形态: 1-3 段,描述 AK/SK 来源、签名机制、权限校验链路;敏感信息硬约束在 sdd=true 的 constitution.md C-S1/C-S2 -->

<security-auth>

## 外部服务依赖
<!-- 形态: 表格,列出本仓运行时调用的外部服务;workspace 跨仓视角见 docs/architecture/01-components-dependencies.md -->

| 服务 | 用途 | 端点 / SDK |
|------|------|------|
| <ext-svc-1> | <ext-svc-1-purpose> | <ext-svc-1-endpoint> |
| <ext-svc-2> | <ext-svc-2-purpose> | <ext-svc-2-endpoint> |
